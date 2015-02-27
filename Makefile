proto:
	protoc --go_out=plugins=gprc:. proto/twitter.proto

.PHONY: proto
