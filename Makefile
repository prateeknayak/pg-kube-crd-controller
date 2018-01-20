.PHONY: dep gen run minikube-up kube-apply

dep:
	docker-compose run --rm dep ensure -v

gen:
	./bin/update-codegen.sh

run:
	go run *.go -kubeconfig=$${HOME}/.kube/config -logtostderr=true -v=2

minikube-up:
	minikube start

kube-apply:
	kubectl apply -f ./kube-objects/crd.yml
	kubectl apply -f ./kube-objects/test.yml
