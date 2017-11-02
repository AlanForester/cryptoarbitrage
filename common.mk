all: $(NAME)
$(NAME): *.go
	go build -o ./deploy/build/$(NAME) -v

release:
	mkdir -p deploy/releases/$(NAME)-"$(VERSION)"/src/$(NAME)

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
