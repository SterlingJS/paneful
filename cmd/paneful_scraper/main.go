package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func getPodLogs() {
	// var kubeconfig *string
	// podLogOptions := corev1.PodLogOptions{}
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	for _, namespace := range namespaces.Items {
		pods, err := clientset.CoreV1().Pods(namespace.String()).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			fmt.Printf("Error fetching pods in the %s namespace!", namespace.String())
			fmt.Println(err.Error())
			continue
		}
		for _, pod := range pods.Items {
			fmt.Printf("Namespace: %s\t\tPod Name: %s", namespace.String(), pod.String())
		}
	}
}

func main() {
	fmt.Println("Hello, World!")
}
