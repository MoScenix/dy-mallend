package rpc

import (
	"sync"

	"github.com/MoScenix/douyin-mall-backend/app/frontend/conf"
	"github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client
	CartClient    cartservice.Client
	once          sync.Once
	once2         sync.Once
	once3         sync.Once
)

func Init() {
	once.Do(initUserClient)
	once2.Do(initProductClient)
	once3.Do(initCartClient)
}

func initUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Consul.Address)
	if err != nil {
		hlog.Fatal(err)
	}

	UserClient, err = userservice.NewClient(
		"user",
		client.WithResolver(r),
	)
	if err != nil {
		hlog.Fatal(err)
	}
}
func initProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Consul.Address)
	if err != nil {
		hlog.Fatal(err)
	}
	ProductClient, err = productcatalogservice.NewClient(
		"product",
		client.WithResolver(r),
	)

	if err != nil {
		hlog.Fatal(err)
	}
}

func initCartClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Consul.Address)
	if err != nil {
		hlog.Fatal(err)
	}
	CartClient, err = cartservice.NewClient(
		"cart",
		client.WithResolver(r),
	)

	if err != nil {
		hlog.Fatal(err)
	}
}
