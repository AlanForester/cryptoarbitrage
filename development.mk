.PHONY: run
pwd := $$(pwd)
setup_env := { export GOPATH="$(pwd):$(pwd)/../.."; };

run:
	bash -c ' $(setup_env) \
	while true; do \
		go run main.go -l debug || true; sleep 1.5; \
	done'
