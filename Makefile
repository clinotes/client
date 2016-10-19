VERSION=0.0.1
PATH_BUILD=build/
S3_BUCKET_NAME=dl.clinot.es

clean:
	@rm -rf ./$(PATH_BUILD)

goxc:
	go get -v github.com/laher/goxc

compile: goxc
	@$(GOPATH)/bin/goxc \
	  -pv=latest \
	  -build-ldflags "-X main.VERSION=$(VERSION)"

build: clean
	docker run -t --rm \
		-v "$(CURDIR):/go/src/github.com/clinotes/client" \
		-w /go/src/github.com/clinotes/client \
		-e GOPATH=/go \
		golang:1.7 \
		make compile

version:
	@echo $(VERSION)

deploy:
	aws s3 sync $(PATH_BUILD)latest s3://$(S3_BUCKET_NAME)/latest

.PHONY: build
