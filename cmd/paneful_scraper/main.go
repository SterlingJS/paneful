package main

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func getPodLogs() {
	var kubeconfig *string
	podLogOptions := corev1.PodLogOptions{}
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	string[] namespaces := clientset.CoreV1().Namespaces()

	for _, namespace := range namespaces {
		pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metaV1.ListOptions{})
		for _, pod := range pods {
			fmt.Printf("Namespace: %s\t\tPod Name: %s")
		}
	}
}

func main() {
	fmt.Println("Hello, World!")
}
