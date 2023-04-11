module github.com/endo-checker/auth

go 1.20

require (
	buf.build/gen/go/envoyproxy/protoc-gen-validate/protocolbuffers/go v1.30.0-20221025150516-6607b10f00ed.1
	github.com/bufbuild/connect-go v1.6.0
	github.com/endo-checker/protostore v1.0.34
	github.com/gookit/cache v0.4.0
	github.com/joho/godotenv v1.5.1
	google.golang.org/genproto v0.0.0-20230403163135-c38d8f061ccd
	google.golang.org/protobuf v1.30.0
)

replace github.com/bufbuild/connect-go => github.com/bufbuild/connect-go v1.5.1

require (
	github.com/gookit/gsr v0.0.8 // indirect
	github.com/rs/cors v1.8.3 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/text v0.9.0 // indirect
)
