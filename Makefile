.PHONY: init dependencies code-gen build run

init:
	# Not working for some reason
	go get -u k8s.io/code-generator

dependencies:
	godep restore -v

code-gen:
	./hack/update-codegen.sh

build: dependencies code-gen
	go build

run: code-gen
	go run *.go -kubeconfig=$${HOME}/.kube/config -logtostderr=true -v=2

deploy-tribe-app:
	kubectl run -i --tty busybox --image=busybox --generator="run-pod/v1" -n "tribe-namespace"
