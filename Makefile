.PHONY: test

test:
	go test -tags all -race ./...

lint:
	go install github.com/mgechev/revive@latest
	revive -config revive.toml ./...