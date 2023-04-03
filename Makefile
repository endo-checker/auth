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
	buf build
	
	buf generate
	cd proto && buf push

.PHONY: run
run:
	go run -tags jwx_es256k .

.PHONY: lint
lint:
	golangci-lint run ./...
	
.PHONY: test
test:
	go test -v -cover -tags jwx_es256k ./...

