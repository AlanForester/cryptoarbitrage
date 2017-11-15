all: $(NAME)
$(NAME): *.go
	go build -o ./deploy/build/$(NAME) -v

start:
	go run main.go -d start

stop:
	go run main.go -d stop

env:
	pwd := $$(pwd)
	export GOBIN=$(pwd)/bin
	export GOPATH=$(pwd):$$(pwd)/../..:~/go
	export PATH=$(PATH):$$(pwd)/bin

gen:
	kallax gen --output ./models.go --input ./models

migrations:
	kallax migrate --input ./models --out ./data/migrations --name initial_schema

migrate $(action) $(env):
	$(foreach var,$(shell . ./bin/parse_config.sh; parse_yaml ./config/$(env).yaml 'config_'),$(eval $(var)))
	$(eval DSN = '$(config_postgres__user):$(config_postgres__pass)@$(config_postgres__host):$(config_postgres__port)/$(config_postgres__name)?sslmode=disable' | sed 's/\"//g')
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
