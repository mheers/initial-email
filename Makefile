all: docker

docker:
	docker build -t mheers/initial-email .

push:
	docker push mheers/initial-email
