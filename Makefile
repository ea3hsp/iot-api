# Albert Espin 2020
# set variables
BINDIR = ./bin
VERSION = 1.0
BINARYFILE = iot-api
DOCKERFILE = iot-api
DOCKERREGISTRY = ea3hsp
DOCKERPUSH = $(DOCKERREGISTRY)/$(DOCKERFILE):$(VERSION)

.PHONY: clean

build:
	go build -o $(BINDIR)/$(BINARYFILE) -i cmd/main.go
docker:
	docker build -t $(DOCKERFILE) .
docker-push:
	docker tag $(DOCKERFILE) $(DOCKERPUSH)
	docker push $(DOCKERPUSH)