package main

import (
	"context"
	"fmt"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	resource := schema.GroupVersionResource{Version: "v1", Resource: "pods"}
	unstructured, err := dynamicClient.
		Resource(resource).
		Namespace("kube-flannel").
		Get(context.TODO(), "kube-flannel-ds-v98vn", v1.GetOptions{})
	if err != nil {
		panic(err)
	}

	pod := &v12.Pod{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructured.UnstructuredContent(), pod)
	if err != nil {
		panic(err)
	}
	fmt.Println(pod.Status.PodIP)
}
