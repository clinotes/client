VERSION=0.0.2
PATH_BUILD=build/
S3_BUCKET_NAME=dl.clinot.es

clean:
	@rm -rf ./$(PATH_BUILD)

compile:
	@go get -v github.com/laher/goxc
	@$(GOPATH)/bin/goxc -pv=latest -build-ldflags "-X main.VERSION=$(VERSION)"

build: clean
	docker run -t --rm -v "$(CURDIR):/go/src/github.com/clinotes/client" -w /go/src/github.com/clinotes/client -e GOPATH=/go golang:1.7 make compile

deploy:
	aws s3 sync $(PATH_BUILD)latest s3://$(S3_BUCKET_NAME)/latest

checksum:
	@cat ./build/latest/cn_latest_darwin_amd64.zip | shasum -a256 | awk '{print $$1}'

version:
	@echo $(VERSION)

publish: clean build deploy checksum

.PHONY: build
