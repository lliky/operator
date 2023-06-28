package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	// client set
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	// get data
	coreV1 := clientset.CoreV1()
	pod, err := coreV1.Pods("kube-flannel").Get(context.TODO(), "kube-flannel-ds-v98vn", v1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pod.Status.PodIP)
	}
}
