#ReF: https://danishpraka.sh/2019/12/07/using-makefiles-for-go.html
BINARY_NAME=main
BINARY_LINUX=$(BINARY_NAME)-lin
APP_LABEL=api-wards
COMMIT_SHA=$(shell git rev-parse --short HEAD)
DOCKER_TAG=myorg/${APP_LABEL}

#.PHONY: bin
## bin: build , create binary for localhost
bin:
	#TODO investigate CGO_ENABLED
	#env CGO_ENABLED=0 go build   -ldflags "-s" -a -installsuffix cgo   -o $(BINARY_NAME) main.go
	go build -o $(BINARY_NAME) .

## run: run binary, localhost:8080
run: bin
	./$(BINARY_NAME) -configFile=config1.yml
## clean: clean go & docker binaries
clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_LINUX)
	docker system prune --all --force

## linuxbin: build binary for docker image
linuxbin:
	env GOOS=linux GOARCH=amd64 go build -o $(BINARY_LINUX)
#check-environment:
#ifndef ENV
#    $(error ENV not set, allowed values - `staging` or `production`)
#endif

## imgrundebug: DEBUG docker img, i.e you can ssh into it!
imgrun-debug: set-minikube-docker linuxbin
	$(eval C_NAME := $(APP_LABEL)-dev)
	$(eval IMG_TAG := $(DOCKER_TAG)-dev)

	docker stop $(C_NAME) || true && docker rm $(C_NAME) || true
	docker build --rm=true --tag=$(IMG_TAG):latest --tag=$(IMG_TAG):$(COMMIT_SHA) -f Dockerfile.debug .
	docker run -p 8081:8080 --name $(C_NAME) -d $(IMG_TAG):latest
	docker exec -it $(C_NAME) sh

imagerun-min: set-minikube-docker linuxbin
	eval $(minikube -p minikube docker-env)
	$(eval C_NAME := $(APP_LABEL))
	$(eval IMG_TAG := $(DOCKER_TAG))
	docker stop $(C_NAME) || true && docker rm $(C_NAME) || true
	docker build --rm=true --tag=$(IMG_TAG):latest --tag=$(IMG_TAG):$(COMMIT_SHA) -f Dockerfile.min .
	docker run -p 8080:8080 --name $(C_NAME) -d $(IMG_TAG):latest
	echo http://localhost:8080
	#docker run -p 8080:8080 --name $(C_NAME)  $(IMG_TAG):latest #debug

set-minikube-docker:
	eval $(minikube -p minikube docker-env)

#unexport DOCKER_TLS_VERIFY DOCKER_HOST DOCKER_CERT_PATH MINIKUBE_ACTIVE_DOCKERD
unset-minikube-docker:
	echo "removing some env variables set for minikube docker registry"
	eval $(minikube docker-env -u)

deploy: set-minikube-docker imagerun-min
	eval $(minikube -p minikube docker-env)            # unix shells
	#Deployment
	printf '=%.0s' {1..100} ; printf "\n"
	kubectl delete -f kube/3deploy.yml
	kubectl apply -f kube/3deploy.yml
	printf '=%.0s' {1..100} ; printf "\n"
	kubectl get deployments
	printf '=%.0s' {1..100} ; printf "\n"
	kubectl get pods -l app=$(APP_LABEL) -o wide
	printf '=%.0s' {1..100} ; printf "\n"
	kubectl delete -f kube/4service.yml
	kubectl apply -f kube/4service.yml
	printf '=%.0s' {1..100} ; printf "\n"
	kubectl get services
	kubectl describe svc $(APP_LABEL)
	kubectl get ep $(APP_LABEL)
	minikube service $(APP_LABEL) --url
#	kubectl apply -f kube/5route.yml
#	kubectl get routes
kdashboard:
	minikube dashboard --url

