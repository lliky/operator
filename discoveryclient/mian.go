package main

import (
	"fmt"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	// discovery client
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err)
	}
	// get data
	APIGroup, APIResources, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		panic(err)
	}
	fmt.Printf("APIGroup: \n\n %v \n\n\n", APIGroup)
	_ = APIResources
	//fmt.Println(APIResources)
}
