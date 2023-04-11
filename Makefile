export APP_PORT := 8080

# proto generates code from the most recent proto file(s)
.PHONY: proto
proto:
	cd proto && buf mod update
	buf lint
	buf build
	# buf breaking --against './.git#branch=main,ref=HEAD~1'
	buf generate
	cd proto && buf push

.PHONY: run
run:
	dapr run \
		--app-id auth \
		--app-port ${APP_PORT} \
		--app-protocol http \
		--config ./.dapr/config.yaml \
		--components-path ./.dapr/components \
		go run .

.PHONY: lint
lint:
	golangci-lint run ./...
	
.PHONY: test
test:
	go test -v ./...
