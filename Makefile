.PHONY: clean proto

clean:
	@rm -rf gen

proto:
	docker run -v $(shell pwd):/workspace --user  $(shell id -u):$(shell id -g) --rm grpckit/omniproto
