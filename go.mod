module github.com/endo-checker/auth

go 1.20

require (
	buf.build/gen/go/envoyproxy/protoc-gen-validate/protocolbuffers/go v1.30.0-20221025150516-6607b10f00ed.1
	github.com/bufbuild/connect-go v1.6.0
	github.com/endo-checker/protostore v1.1.0
	github.com/gookit/cache v0.4.0
	github.com/joho/godotenv v1.5.1
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
)

replace github.com/bufbuild/connect-go => github.com/bufbuild/connect-go v1.5.1

require (
	github.com/go-chi/chi v1.5.4 // indirect
	github.com/gookit/gsr v0.0.8 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
)
