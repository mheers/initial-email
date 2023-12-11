all: build

build: docker

docker: ##  Builds the application for amd64 and arm64
	$(eval VERSION := $(shell cat .VERSION | grep VERSION | cut -d'=' -f2))
	docker buildx build --platform linux/amd64,linux/arm64 -t mheers/initial-email:$(VERSION) --push .
	docker buildx build --platform linux/amd64,linux/arm64 -t mheers/initial-email:latest --push .
