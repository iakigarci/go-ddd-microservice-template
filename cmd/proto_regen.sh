#install -> go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
#install -> go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

export PATH="$PATH:$(go env GOPATH)/bin"

# add file name to generate the new proto

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./../api/proto/.proto
