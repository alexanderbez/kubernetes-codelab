.PHONY: build push

GCR_URI=gcr.io/kubernetes-codelab-143900
IMAGE=hello-world
TAG?=0.0.3
GOARCH?=amd64
GOOS?=linux

build:
	env GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o pkg/$(IMAGE)
	docker build -t $(IMAGE):$(TAG) .

push:
	docker tag $(IMAGE):$(TAG) $(GCR_URI)/$(IMAGE):$(TAG)
	gcloud docker -- push $(GCR_URI)/$(IMAGE):$(TAG)

deploy:
	kubectl set image deployment/$(IMAGE) $(IMAGE)=$(GCR_URI)/$(IMAGE):$(TAG)
