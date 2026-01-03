package dal

import (
	"github.com/MoScenix/douyin-mall-backend/app/frontend/biz/dal/mysql"
	"github.com/MoScenix/douyin-mall-backend/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
