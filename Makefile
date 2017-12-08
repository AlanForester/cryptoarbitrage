# ./Makefile
VERSION := $(shell date +'%Y%m%d%H').$(shell git rev-parse --short=8 HEAD)
NAME := $(shell echo cryptoarbitrage)

GOPWD := $(shell pwd)
GOBASEDIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/../..)
GOHOMEDIR := $(abspath $(HOME)/go)

GOPATH = $(GOPWD):$(GOBASEDIR):$(GOHOMEDIR)
GOBIN = $(GOHOMEDIR)/bin

$(info root makefile GOPATH=$(GOPATH))
$(info root makefile GOBIN=$(GOBIN))

all:
	@echo "Project:" $(NAME) $(VERSION)

bulds: $(NAME)
$(NAME): *.go
	go build -o ./deploy/build/$(NAME) -v

start:
	go run main.go -e=production -d start

stop:
	go run main.go -e=production -d stop

run:
	go run main.go -e=development

modelsgen:
	kallax gen --output ./models.go --input ./models/
	git add ./models/models.go

migrations:
	kallax migrate --input ./models --out ./data/migrations --name initial_schema
	git add ./data/migrations/*

migrate $(action) $(env):
	$(foreach var,$(shell . ./bin/parse_config.sh; parse_yaml ./config/$(env).yaml 'config_'),$(eval $(var)))
	$(eval DSN = '$(config_storage_postgres_user):$(config_storage_postgres_pass)@$(config_storage_postgres_host):$(config_storage_postgres_port)/$(config_storage_postgres_name)?sslmode=disable' | sed 's/\"//g')
ifeq ($(action),up)
	kallax migrate up --dir ./data/migrations --dsn '$(shell echo $(DSN))' --all
endif
ifeq ($(action),down)
	kallax migrate down --dir=./data/migrations --dsn '$(shell echo $(DSN))' -n 1
endif

release:
	mkdir -p deploy/releases/$(NAME)-"$(VERSION)"
	/src/$(NAME)

	rsync -avzr --delete \
		--filter='- $(NAME)-*' \
		--filter='- /$(NAME)' \
		--filter='+ /.git/' \
		--filter='+ /.gitignore/' \
		--filter='+ /releases/' \
		--filter='+ /glide.lock/' \
		--filter='+ /README.MD/' \
		--filter='- .*' \
		--filter='- *~' \
		--filter='- *.org' \
		. deploy/releases/$(NAME)-"$(VERSION)"/src/$(NAME)

	tar czf deploy/releases/$(NAME)-"$(VERSION)".tgz $(NAME)-"$(VERSION)"
