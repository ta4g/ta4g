.PHONY: clean proto

clean:
	@rm -rf gen

proto:
	docker pull grpckit/grpckit:1.37_0
	docker run --rm -v $(shell pwd):/workspace -it grpckit/grpckit:1.37_0 /usr/local/bin/omniproto
	echo "Protos generated"
