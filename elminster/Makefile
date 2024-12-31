GO_FILES ?= $$(find . -name '*.go' | grep -v generated)

clean:
	rm -f ./dist/elminster

build:
	go build -o ./dist/elminster

rest: clean build
	./dist/elminster server

grpc: clean build
	./dist/elminster grpc

migrate: clean build
	./dist/elminster migrate

migrate-down: clean build
	./dist/elminster migrate --down

generate:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/server/connections/grpc/pb/character.proto \
	pkg/server/connections/grpc/pb/weapon.proto \
	pkg/server/connections/grpc/pb/ping.proto

format:
	echo $(GO_FILES)
	gofmt -s -l -w $(GO_FILES)
	goimports -w $(GO_FILES)
