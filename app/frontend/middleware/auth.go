package middleware

import (
	"context"
	"strings"

	"github.com/MoScenix/douyin-mall-backend/app/frontend/biz/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sessions"
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId == nil {
			c.Next(ctx)
			return
		}
		ctx = context.WithValue(ctx, utils.UserIdKey, userId)
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		path := c.Path()
		if ctx.Value(utils.UserIdKey) == nil {
			if strings.HasPrefix(string(path), "/order") {
				c.Redirect(consts.StatusFound, []byte("/sign-in"))
				c.Abort()
				return
			}
		}
		c.Next(ctx)
	}
}
