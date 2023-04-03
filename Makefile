export APP_PORT := 8084

# load env vars
-include .env
export auth0-api-identifier  := $(value AUTH0_API_IDENTIFIER)
export auth0-domain  := $(value AUTH0_DOMAIN)
export auth0-client-id  := $(value AUTH_CLIENT_ID)
export auth0-client-secret  := $(value AUTH_CLIENT_SECRET)
export audience  := $(value AUTH_AUDIENCE)

# proto generates code from the most recent proto file(s)
.PHONY: proto
proto:
	cd proto && buf mod update
	buf lint
	# buf breaking --against './.git#branch=main,ref=HEAD~1'
	buf build
	buf generate
	cd proto && buf push

.PHONY: run
rungo:
	go run main.go

.PHONY: run
run:
	dapr run \
		--app-id auth \
		--app-port ${APP_PORT} \
		--app-protocol http \
		go run .

.PHONY: kill
kill:
	lsof -P -i TCP -s TCP:LISTEN | grep ${APP_PORT} | awk '{print $2}' | { read pid; kill -9 ${pid}; }
	lsof -P -i TCP -s TCP:LISTEN | grep 9090 | awk '{print $2}' | { read pid; kill -9 ${pid}; }

.PHONY: lint
lint:
	golangci-lint run ./...
	
.PHONY: test
test:
	go test -v ./handler/...
