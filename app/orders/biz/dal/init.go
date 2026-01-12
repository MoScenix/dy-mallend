package dal

import (
	"github.com/MoScenix/douyin-mall-backend/app/orders/biz/dal/mysql"
	"github.com/MoScenix/douyin-mall-backend/app/orders/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
