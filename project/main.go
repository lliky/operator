package main

import (
	"github.com/liky/operator/project/pkg"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {
	// 1.config
	// 2.client
	// 3.informer
	// 4.add event handler
	// 5.informer.Start

	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		clusterConfig, err := rest.InClusterConfig()
		if err != nil {
			log.Fatalln("can not get config")
		}
		config = clusterConfig
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("can not get client set")
	}

	factory := informers.NewSharedInformerFactory(clientSet, 0)
	serviceInformer := factory.Core().V1().Services()
	ingressInformer := factory.Networking().V1().Ingresses()

	controller := pkg.NewController(clientSet, serviceInformer, ingressInformer)
	stopCh := make(chan struct{})
	factory.Start(stopCh)
	factory.WaitForCacheSync(stopCh)
	controller.Run(stopCh)
}
