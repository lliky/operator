package main

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/api/meta"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	list, err := clientSet.CoreV1().Pods("default").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(len(list.Items))

	watch, err := clientSet.CoreV1().Pods("default").Watch(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for {
		event := <-watch.ResultChan()
		fmt.Println(event.Object.GetObjectKind().GroupVersionKind())
		accessor, err := meta.ListAccessor(event.Object)
		if err != nil {
			panic(err)
		}
		fmt.Println(accessor.GetResourceVersion())
	}
}
