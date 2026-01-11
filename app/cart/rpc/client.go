package rpc

import (
	"sync"

	"github.com/MoScenix/douyin-mall-backend/app/cart/conf"
	"github.com/MoScenix/douyin-mall-backend/common/clientsuit"
	"github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	err           error
	registryAddr  string
	serviceName   string
)

func InitClient() {
	once.Do(func() {
		registryAddr = conf.GetConf().Registry.RegistryAddress[0]
		serviceName = conf.GetConf().Kitex.Service
		initProductClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(&clientsuit.CommonClient{
			ServiceName: serviceName,
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	//cartutils.MustHandleError(err)
}
