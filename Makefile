GO111MODULE := on
DOCKER_TAG := $(or ${GIT_TAG_NAME}, latest)

all: rgw-exporter

.PHONY: rgw-exporter
rgw-exporter:
	go build -tags netgo -o bin/rgw-exporter *.go
	strip bin/rgw-exporter

.PHONY: dockerimages
dockerimages:
	docker build -t mwennrich/rgw-exporter:${DOCKER_TAG} .

.PHONY: dockerpush
dockerpush:
	docker push mwennrich/rgw-exporter:${DOCKER_TAG}

.PHONY: clean
clean:
	rm -f bin/*
