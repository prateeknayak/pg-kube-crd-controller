/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"time"

	"github.com/golang/glog"
	fruit_clientset "github.com/prateeknayak/pg-kube-crd-controller/pkg/client/clientset/versioned"
	fruit_informers "github.com/prateeknayak/pg-kube-crd-controller/pkg/client/informers/externalversions"
	"github.com/prateeknayak/pg-kube-crd-controller/pkg/signals"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeConfig string
)

func main() {

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	var config *rest.Config
	var err error

	// if flag has not been passed and env not set, presume running in cluster
	if kubeConfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfig)
	} else {
		config, err = rest.InClusterConfig()
	}

	if nil != err {
		glog.Fatalf("Error building kubeConfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		glog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	fruitClient, err := fruit_clientset.NewForConfig(config)
	if err != nil {
		glog.Fatalf("Error building example clientset: %s", err.Error())
	}

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	exampleInformerFactory := fruit_informers.NewSharedInformerFactory(fruitClient, time.Second*30)

	//controller := NewController(kubeClient, fruitClient, kubeInformerFactory, exampleInformerFactory)

	go kubeInformerFactory.Start(stopCh)
	go exampleInformerFactory.Start(stopCh)

	//if err = controller.Run(2, stopCh); err != nil {
	//	glog.Fatalf("Error running controller: %s", err.Error())
	//}
}

func init() {
	flag.StringVar(
		&kubeConfig,
		"kubeconfig",
		"",
		"Path to a kubeConfig. Only required if out-of-cluster.",
	)

	flag.Parse()
}
