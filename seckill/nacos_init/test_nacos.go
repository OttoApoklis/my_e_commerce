package nacos_init

import (
	"fmt"
	"log"
	"time"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func NacosInit() {
	time.Sleep(10 * time.Second)
	log.Printf("开始初始化nacos")
	time.Sleep(2 * time.Second)
	// 创建 Nacos 客户端配置
	sc := []constant.ServerConfig{
		{
			IpAddr: "localhost",
			Port:   8848,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:          "public",
		TimeoutMs:            5000,
		NotLoadCacheAtStart:  true,
		UpdateCacheWhenEmpty: true,
		LogDir:               "/tmp/nacos_init/log",
		CacheDir:             "/tmp/nacos_init/cache",
		LogLevel:             "debug",
	}

	// 创建 Nacos 客户端
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		log.Fatalf("Error creating Nacos client: %v", err)
	}

	// 服务注册
	instance := vo.RegisterInstanceParam{
		Ip:          "localhost",
		Port:        8081,
		ServiceName: "seckill-service",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"version": "1.0"},
		ClusterName: "cluster-a",
		GroupName:   "group-a",
	}

	_, err = client.RegisterInstance(instance)
	if err != nil {
		log.Fatalf("Error registering instance: %v", err)
	}
	log.Println("Service registered successfully.")

	// 服务发现
	serviceName := "example-service"
	groupName := "group-a"
	clusters := []string{"cluster-a"}
	log.Println("开始服务发现\n")
	// 等待服务发现完成
	time.Sleep(5 * time.Second)

	instances, err := client.SelectInstances(vo.SelectInstancesParam{
		ServiceName: serviceName,
		GroupName:   groupName,
		Clusters:    clusters,
	})
	if err != nil {
		log.Fatalf("Error selecting instances: %v", err)
	}

	for _, instance := range instances {
		fmt.Printf("Discovered instance: %+v\n", instance)
	}

	// 服务注销
	//_, err = client.DeregisterInstance(vo.DeregisterInstanceParam{
	//	Ip:          "127.0.0.1",
	//	Port:        8080,
	//	ServiceName: "example-service",
	//	Ephemeral:   true,
	//	ClusterName: "cluster-a",
	//	GroupName:   "group-a",
	//})
	//if err != nil {
	//	log.Fatalf("Error deregistering instance: %v", err)
	//}
	//log.Println("Service deregistered successfully.")
}
