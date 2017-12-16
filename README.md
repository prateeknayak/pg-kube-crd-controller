# pg-kube-crd-controller

Playing around with crds and controllers in minikube environment.

## Code Gen (make me sad)

In order to get started with writing controllers there is a fair amount of code that needs to be generated. Unfortunately, this is the pattern or proposed way at the moment. You don't have to but it is good if you do this.

1. Run the following commands to get the `code-generator` and `api-machinery`

	```
	go get -d k8s.io/code-genereator
	```
	```
	go get -d k8s.io/apimachinery
	```

	**NOTE** Make sure you are using the right version of apimachinery depending 	on your needs. Above step will get you the HEAD of the repo.

1. Run the generator

	```
	touch /tmp/hate-this.txt;
	./generate-groups.sh all \
	github.com/prateeknayak/pg-kube-crd-controller/pkg/client \
	github.com/prateeknayak/pg-kube-crd-controller/pkg/apis \
	"mycontroller:v1alpha1" \
	--go-header-file /tmp/hate-this.txt
	```
	

## Tools of Power

```
brew tap ahmetb/kubectx https://github.com/ahmetb/kubectx.git
brew install kubectx --with-short-names
```