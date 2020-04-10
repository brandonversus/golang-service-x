PHONY:  vendor build bump image push deploy

vendor:
	go mod vendor
	go mod tidy

build: vendor
	CGO_ENABLED=0 go build -o main .

bump: build
	semver bump patch $(shell cat version) > version

image: bump
	docker build . -t brandonversus/service-x:latest
	docker tag brandonversus/service-x:latest brandonversus/service-x:$(shell cat version)

push:
	docker push brandonversus/service-x:$(shell cat version)

deploy:
	cd manifests/service-a/manifests/base && kustomize edit set image brandonversus/service-x=brandonversus/service-x:$(shell cat version) && kubectl apply -k .
	cd manifests/service-b/manifests/base && kustomize edit set image brandonversus/service-x=brandonversus/service-x:$(shell cat version) && kubectl apply -k .
	cd manifests/service-c/manifests/base && kustomize edit set image brandonversus/service-x=brandonversus/service-x:$(shell cat version) && kubectl apply -k .
	watch kubectl get pods
