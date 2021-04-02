package main

import (
	"fmt"
	"time"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/util"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	sc := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   8848,
		},
	}
	//or a more graceful way to create ServerConfig
	_ = []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848),
	}

	cc := constant.ClientConfig{
		NamespaceId:         "dev", //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}
	//or a more graceful way to create ClientConfig
	_ = *constant.NewClientConfig(
		constant.WithNamespaceId("dev"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithRotateTime("1h"),
		constant.WithMaxAge(3),
		constant.WithLogLevel("debug"),
	)

	// a more graceful way to create naming client
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	//Register with default cluster and group
	//ClusterName=DEFAULT,GroupName=DEFAULT_GROUP
	ServiceClient_RegisterServiceInstance(client, vo.RegisterInstanceParam{
		Ip:          "10.0.0.10",
		Port:        8848,
		ServiceName: "demo.go",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
	})

	//Register with cluster name
	//GroupName=DEFAULT_GROUP
	ServiceClient_RegisterServiceInstance(client, vo.RegisterInstanceParam{
		Ip:          "10.0.0.11",
		Port:        8848,
		ServiceName: "demo.go",
		Weight:      10,
		ClusterName: "cluster-a",
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	})

	//Register different cluster
	//GroupName=DEFAULT_GROUP
	ServiceClient_RegisterServiceInstance(client, vo.RegisterInstanceParam{
		Ip:          "10.0.0.12",
		Port:        8848,
		ServiceName: "demo.go",
		Weight:      10,
		ClusterName: "cluster-b",
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	})

	//Register different group
	ServiceClient_RegisterServiceInstance(client, vo.RegisterInstanceParam{
		Ip:          "10.0.0.13",
		Port:        8848,
		ServiceName: "demo.go",
		Weight:      10,
		ClusterName: "cluster-b",
		GroupName:   "group-a",
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	})
	ServiceClient_RegisterServiceInstance(client, vo.RegisterInstanceParam{
		Ip:          "10.0.0.14",
		Port:        8848,
		ServiceName: "demo.go",
		Weight:      10,
		ClusterName: "cluster-b",
		GroupName:   "group-b",
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	})

	//DeRegister with ip,port,serviceName
	//ClusterName=DEFAULT, GroupName=DEFAULT_GROUP
	//Note:ip=10.0.0.10,port=8848 should belong to the cluster of DEFAULT and the group of DEFAULT_GROUP.
	ServiceClient_DeRegisterServiceInstance(client, vo.DeregisterInstanceParam{
		Ip:          "10.0.0.10",
		Port:        8848,
		ServiceName: "demo.go",
		Ephemeral:   true, //it must be true
	})

	//DeRegister with ip,port,serviceName,cluster
	//GroupName=DEFAULT_GROUP
	//Note:ip=10.0.0.10,port=8848,cluster=cluster-a should belong to the group of DEFAULT_GROUP.
	ServiceClient_DeRegisterServiceInstance(client, vo.DeregisterInstanceParam{
		Ip:          "10.0.0.11",
		Port:        8848,
		ServiceName: "demo.go",
		Cluster:     "cluster-a",
		Ephemeral:   true, //it must be true
	})

	//DeRegister with ip,port,serviceName,cluster,group
	ServiceClient_DeRegisterServiceInstance(client, vo.DeregisterInstanceParam{
		Ip:          "10.0.0.14",
		Port:        8848,
		ServiceName: "demo.go",
		Cluster:     "cluster-b",
		GroupName:   "group-b",
		Ephemeral:   true, //it must be true
	})

	//Get service with serviceName
	//ClusterName=DEFAULT, GroupName=DEFAULT_GROUP
	ServiceClient_GetService(client, vo.GetServiceParam{
		ServiceName: "demo.go",
	})
	//Get service with serviceName and cluster
	//GroupName=DEFAULT_GROUP
	ServiceClient_GetService(client, vo.GetServiceParam{
		ServiceName: "demo.go",
		Clusters:    []string{"cluster-a", "cluster-b"},
	})
	//Get service with serviceName ,group
	//ClusterName=DEFAULT
	ServiceClient_GetService(client, vo.GetServiceParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",
	})

	//SelectAllInstance return all instances,include healthy=false,enable=false,weight<=0
	//ClusterName=DEFAULT, GroupName=DEFAULT_GROUP
	ServiceClient_SelectAllInstances(client, vo.SelectAllInstancesParam{
		ServiceName: "demo.go",
	})

	//SelectAllInstance
	//GroupName=DEFAULT_GROUP
	ServiceClient_SelectAllInstances(client, vo.SelectAllInstancesParam{
		ServiceName: "demo.go",
		Clusters:    []string{"cluster-a", "cluster-b"},
	})

	//SelectAllInstance
	//ClusterName=DEFAULT
	ServiceClient_SelectAllInstances(client, vo.SelectAllInstancesParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",
	})

	//SelectInstances only return the instances of healthy=${HealthyOnly},enable=true and weight>0
	//ClusterName=DEFAULT,GroupName=DEFAULT_GROUP
	ServiceClient_SelectInstances(client, vo.SelectInstancesParam{
		ServiceName: "demo.go",
	})

	//SelectOneHealthyInstance return one instance by WRR strategy for load balance
	//And the instance should be health=true,enable=true and weight>0
	//ClusterName=DEFAULT,GroupName=DEFAULT_GROUP
	ServiceClient_SelectOneHealthyInstance(client, vo.SelectOneHealthInstanceParam{
		ServiceName: "demo.go",
	})

	//Subscribe key=serviceName+groupName+cluster
	//Note:We call add multiple SubscribeCallback with the same key.
	param := &vo.SubscribeParam{
		ServiceName: "demo.go",
		Clusters:    []string{"cluster-b"},
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			fmt.Printf("callback111 return services:%s \n\n", util.ToJsonString(services))
		},
	}
	ServiceClient_Subscribe(client, param)
	param2 := &vo.SubscribeParam{
		ServiceName: "demo.go",
		Clusters:    []string{"cluster-b"},
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			fmt.Printf("callback222 return services:%s \n\n", util.ToJsonString(services))
		},
	}
	ServiceClient_Subscribe(client, param2)
	ServiceClient_RegisterServiceInstance(client, vo.RegisterInstanceParam{
		Ip:          "10.0.0.112",
		Port:        8848,
		ServiceName: "demo.go",
		Weight:      10,
		ClusterName: "cluster-b",
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	})
	//wait for client pull change from server
	time.Sleep(10 * time.Second)

	//Now we just unsubscribe callback1, and callback2 will still receive change event
	ServiceClient_UnSubscribe(client, param)
	ServiceClient_DeRegisterServiceInstance(client, vo.DeregisterInstanceParam{
		Ip:          "10.0.0.112",
		Ephemeral:   true,
		Port:        8848,
		ServiceName: "demo.go",
		Cluster:     "cluster-b",
	})
	//wait for client pull change from server
	time.Sleep(10 * time.Second)

	//GeAllService will get the list of service name
	//NameSpace default value is public.If the client set the namespaceId, NameSpace will use it.
	//GroupName default value is DEFAULT_GROUP
	ExampleServiceClient_GetAllService(client, vo.GetAllServiceInfoParam{
		PageNo:   1,
		PageSize: 10,
	})

	ExampleServiceClient_GetAllService(client, vo.GetAllServiceInfoParam{
		NameSpace: "0e83cc81-9d8c-4bb8-a28a-ff703187543f",
		PageNo:    1,
		PageSize:  10,
	})
}
