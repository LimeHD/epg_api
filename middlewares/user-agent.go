package middlewares

import (
	"fmt"
	"github.com/LimeHD/epg_api/helpers"
	"github.com/savsgio/atreugo/v11"
)

func UserAgent(ctx *atreugo.RequestCtx) error {
	platform := helpers.Platform{Id: helpers.SELF_ID, VersionCode: 1}
	userAgent := ctx.Request.Header.Peek("User-Agent")
	userAgentSmart := ctx.Request.Header.Peek("User-Agent-Smart")

	if userAgentSmart != nil {
		userAgent = userAgentSmart
	}

	helpers.SetPlatform(userAgent, &platform)
	ctx.SetUserValue("platform", platform)

	fmt.Println(platform.Marshal())

	return ctx.Next()
}
