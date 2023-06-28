package main

import (
	"context"
	"fmt"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"
	//client
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}
	//get data
	pod := v1.Pod{}
	err = restClient.Get().
		Namespace("kube-flannel").
		Resource("pods").
		Name("kube-flannel-ds-v98vn").
		Do(context.Background()).Into(&pod)
	if err != nil {
		fmt.Println("err")
	} else {
		fmt.Println(pod.Status.PodIP)
		fmt.Println(pod.GetObjectMeta())
	}
}
