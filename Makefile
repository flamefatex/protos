.PHONY: gen-go package

BIN_NAME:=$(notdir $(shell pwd))
IMAGE_NAME := flamefatex/${BIN_NAME}
REMOTE_DOCKER_URI := flamefatex/${BIN_NAME}

# git信息
BRANCH := $(shell git branch | grep \* | cut -d ' ' -f2)
GIT_COMMIT := $(shell git rev-parse HEAD)
GIT_DIRTY := $(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
GIT_COMMIT := ${GIT_COMMIT}${GIT_DIRTY}
#VERSION := $(shell git describe --tags)
VERSION := v1.1.3
ifeq "${VERSION}" ""
	VERSION := "notag"
endif

SWAGGER_FILES=swagger/title.json $(wildcard swagger/*/*/*.json) $(wildcard swagger/*/*/*/*.json)

projects=$(shell ls -l ./src| grep ^d |awk '{print $$9}' )

default: gen-proto swagger-mixin gen-doc

package:
	docker build --build-arg VERSION=${VERSION} --build-arg GIT_COMMIT=${GIT_COMMIT} -t ${IMAGE_NAME}:${VERSION} .

push: package
	docker push ${IMAGE_NAME}:${VERSION}

gen-proto:
	@rm -rf goout
	docker run -it -v `pwd`:/opt/protos ${IMAGE_NAME}:${VERSION} bash gen.sh proto

gen-doc:
	@rm -rf doc
	docker run -it -v `pwd`:/opt/protos ${IMAGE_NAME}:${VERSION} bash gen.sh doc

gen-swagger:
	@find swagger -mindepth 1 -type d -exec rm -rf {} +
	docker run -it -v `pwd`:/opt/protos ${IMAGE_NAME}:${VERSION} bash gen.sh swagger

gen-mock:
	@rm -rf gooutmock
	#docker run -it -v `pwd`:/opt/protos ${IMAGE_NAME}:${VERSION} bash gen.sh mock
	@docker run -it -v `pwd`:/go/src/github.com/flamefatex/protos ${IMAGE_NAME}:${VERSION} bash /go/src/github.com/flamefatex/protos/gen.sh mock
	@rm -rf gomock_reflect*

swagger-mixin: gen-swagger
	bash gen.sh swagger_mixin

swagger-ui:
	@echo "http://localhost:8080/"
	@docker run -it --rm -p 8080:8080 -v `pwd`/swagger/swagger.json:/app/swagger.json swaggerapi/swagger-ui:latest
	#docker run --rm -it -v `pwd`:/tmp/protos -p 8080:8080 -w /tmp/protos 192.168.0.103:8080/flamefatex/goswagger:latest serve swagger/swagger.json -Fswagger --port 8080 --no-open
clean:
	@rm -rf goout javaout swagger
