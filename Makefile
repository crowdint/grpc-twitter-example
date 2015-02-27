proto:
	protoc -I proto/ proto/twitter.proto --go_out=plugins=grpc:proto/

.PHONY: proto
