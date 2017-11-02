.PHONY: run
pwd := $$(pwd)
setup_env := { export GOPATH="$(pwd):$(pwd)/../.."; };

run:
	bash -c ' $(setup_env) \
	go run main.go \
	done'
