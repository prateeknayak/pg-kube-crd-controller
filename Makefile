.PHONY: dep gen run minikube-up apply ns-apply crd-apply test-apply test-delete crd-delete ns-delete

dcr ?= docker-compose run --rm

dep:
	@$(dcr) dep ensure -v

gen:
	./bin/update-codegen.sh

run:
	go run *.go -kubeconfig=$${HOME}/.kube/config -logtostderr=true -v=2

minikube-up:
	minikube start

ns-apply:
	@$(dcr) kubectl apply -f ./kube-objects/namespaces.yml

ns-delete:
	@$(dcr) kubectl delete -f ./kube-objects/namespaces.yml

crd-apply:
	@$(dcr) kubectl apply -f ./kube-objects/crd.yml

crd-delete:
	@$(dcr) kubectl delete -f ./kube-objects/crd.yml

test-apply:
	@$(dcr) kubectl apply -f ./kube-objects/test.yml

test-delete:
	@$(dcr) kubectl delete -f ./kube-objects/test.yml

apply: ns-apply crd-apply test-apply

delete: test-delete crd-delete ns-delete
