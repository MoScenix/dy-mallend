package dal

import (
	"github.com/MoScenix/douyin-mall-backend/app/cart/biz/dal/mysql"
	"github.com/MoScenix/douyin-mall-backend/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
