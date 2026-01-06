package dal

import (
	"github.com/MoScenix/douyin-mall-backend/app/product/biz/dal/mysql"
	"github.com/MoScenix/douyin-mall-backend/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
