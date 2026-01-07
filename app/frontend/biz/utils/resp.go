package utils

import (
	"context"
	"os"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}
func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	content["user_id"] = ctx.Value(UserIdKey)
	if ctx.Value(UserIdKey) != nil {
		userIdStr := strconv.Itoa(int(ctx.Value(UserIdKey).(float64)))
		if _, statErr := os.Stat("./static/image/avatar/" + userIdStr + ".jpg"); statErr == nil {
			content["avatar"] = "./static/image/avatar/" + userIdStr + ".jpg"
		}
	}
	return content
}
