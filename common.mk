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
	go generate ./...

migrate:
	kallax migrate --input ./models --out ./data/migrations --name initial_schema

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
