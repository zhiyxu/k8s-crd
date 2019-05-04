//main函数的主要工作就是：
// 定义并初始化一个自定义控制器（Customer Controller）
// 然后启动它

package main

import (
	"flag"
	"time"

	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	clientset "github.com/zhiyxu/k8s-crd/pkg/client/clientset/versioned"
	informers "github.com/zhiyxu/k8s-crd/pkg/client/informers/externalversions"
	"github.com/zhiyxu/k8s-crd/pkg/signals"
)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	flag.Parse()

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	// 第一步：
	// 根据提供的Master配置，创建一个Kubernetes的client和Network对象的client
	// 如果没有提供Master的配置，main函数会直接使用一种名叫InClusterConfig的方式创建client
	// 这个方式，会假设自定义控制器是以Pod的方式运行在Kubernetes集群里的
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	// 创建kubeClient
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	// 创建networkClient
	networkClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %s", err.Error())
	}

	// 第二步：
	// 为Network对象创建一个叫做InformerFactory的工厂
	// 并使用它生成一个Network对象的Informer，传递给控制器
	networkInformerFactory := informers.NewSharedInformerFactory(networkClient, time.Second*30)

	controller := NewController(kubeClient, networkClient,
		networkInformerFactory.Samplecrd().V1().Networks())

	// 第三步：
	// 启动上述Informer，然后执行controller.Run启动自定义控制器
	go networkInformerFactory.Start(stopCh)

	if err = controller.Run(2, stopCh); err != nil {
		glog.Fatalf("Error running controller: %s", err.Error())
	}
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}
