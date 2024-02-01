package main

import (
	"context"
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/client-go/tools/clientcmd"
	"os"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// 解析命令行参数
	var kubeconfigPath, ns, deploymentName string
	var cpuRequestStr, memoryRequestStr, cpuLimitStr, memoryLimitStr string
	flag.StringVar(&kubeconfigPath, "kubeconfig", "", "kubeconfig path")
	flag.StringVar(&ns, "namespace", "", "namespace")
	flag.StringVar(&deploymentName, "deployment", "", "Deployment Name")
	flag.StringVar(&cpuRequestStr, "cpu-request", "", "CPU request")
	flag.StringVar(&memoryRequestStr, "memory-request", "", "Mem request")
	flag.StringVar(&cpuLimitStr, "cpu-limit", "", "CPU limit")
	flag.StringVar(&memoryLimitStr, "memory-limit", "", "Mem limit")
	flag.Parse()

	ctx := context.Background()

	// 检查参数
	if deploymentName == "" {
		fmt.Println("Deployment name is empty：--deployment=<Deployment Name>")
		os.Exit(1)
	}

	if ns == "" {
		ns = "default"
	}

	// 创建 Kubernetes 客户端
	var config *rest.Config
	var err error
	if kubeconfigPath != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err != nil {
			fmt.Println("build Kubernetes client failed:", err)
			os.Exit(1)
		}
	} else {
		// Get in-cluster k8s config
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Println("build Kubernetes client failed:", err)
			os.Exit(1)
		}
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("can not create Kubernetes client:", err)
		os.Exit(1)
	}

	// 获取 Deployment 对象
	deployment, err := clientSet.AppsV1().Deployments(ns).Get(ctx, deploymentName, v1.GetOptions{})
	if err != nil {
		fmt.Println("can not get deployment:", err)
		os.Exit(1)
	}
	var needUpdate bool

	// 修改资源请求和限制
	if cpuRequestStr != "" {
		cpuRequest, err := resource.ParseQuantity(cpuRequestStr)
		if err != nil {
			fmt.Println("parse cpu request failed: ", err)
			os.Exit(1)
		}
		needUpdate = true
		deployment.Spec.Template.Spec.Containers[0].Resources.Requests["cpu"] = cpuRequest
	}
	if memoryRequestStr != "" {
		memoryRequest, err := resource.ParseQuantity(memoryRequestStr)
		if err != nil {
			fmt.Println("parse memory request failed: ", err)
			os.Exit(1)
		}
		needUpdate = true
		deployment.Spec.Template.Spec.Containers[0].Resources.Requests["memory"] = memoryRequest
	}
	if cpuLimitStr != "" {
		cpuLimit, err := resource.ParseQuantity(cpuLimitStr)
		if err != nil {
			fmt.Println("parse cpu limit failed: ", err)
			os.Exit(1)
		}
		needUpdate = true
		deployment.Spec.Template.Spec.Containers[0].Resources.Limits["cpu"] = cpuLimit
	}
	if memoryLimitStr != "" {
		memoryLimit, err := resource.ParseQuantity(memoryLimitStr)
		if err != nil {
			fmt.Println("parse memory limit failed: ", err)
			os.Exit(1)
		}
		needUpdate = true
		deployment.Spec.Template.Spec.Containers[0].Resources.Limits["memory"] = memoryLimit
	}

	// 更新 Deployment 对象
	if needUpdate {
		_, err = clientSet.AppsV1().Deployments(ns).Update(ctx, deployment, v1.UpdateOptions{})
		if err != nil {
			fmt.Println("can not update deployment: ", err)
			os.Exit(1)
		}

		fmt.Printf("succeed to update deployment %s \n", deploymentName)
		os.Exit(0)
	}

	fmt.Printf("nothing to update")

}
