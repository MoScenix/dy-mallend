package dal

import (
	"github.com/MoScenix/douyin-mall-backend/app/user/biz/dal/mysql"
	"github.com/MoScenix/douyin-mall-backend/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
