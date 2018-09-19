.PHONY: compile
install:
	go install ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	gometalinter
