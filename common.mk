all: $(NAME)
$(NAME): *.go
	go build -o ./deploy/build/$(NAME) -v

start:
	go run main.go -d start

stop:
	go run main.go -d stop

env:
	bash -c ' \
		pwd := $$(pwd) \
		export GOPATH="$(pwd):$(pwd)/../.." \
		export PATH="$(PATH):$(pwd)/bin" \
	done'

gen:
	go $$(pwd)/vendor/gopkg.in/src-d/go-kallax.v1/generator/cli/kallax/cmd.go gen
	# go generate ./...

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
