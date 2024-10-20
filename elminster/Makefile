clean:
	rm -f ./dist/elminster

build:
	go build -o ./dist/elminster

rest: clean build
	./dist/elminster server

grpc: clean build
	./dist/elminster grpc

generate:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/server/connections/grpc/pb/character.proto \
	pkg/server/connections/grpc/pb/weapon.proto \
	pkg/server/connections/grpc/pb/ping.proto

