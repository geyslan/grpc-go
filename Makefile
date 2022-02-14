PROTOS = $(wildcard ./proto/*.proto)
PROTOS_PB_GO = $(patsubst %.proto, %.pb.go, $(PROTOS))

.PHONY: protoc
protoc: $(PROTOS_PB_GO)

$(PROTOS_PB_GO): $(PROTOS)
	./generate.sh

.PHONY: server
server:
	go run server/server.go
	
.PHONY: client
client:
	go run client/client.go

.PHONY: clean
clean:
	rm proto/*.pb.go