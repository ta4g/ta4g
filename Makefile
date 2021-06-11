.PHONY: clean proto

clean:
	@rm -rf gen

proto:
	docker run -v $(shell pwd):/workspace --rm grpckit/omniproto
