package middlewares

import (
	"github.com/LimeHD/epg_api/helpers"
	"github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger"
)

func Country(ctx *atreugo.RequestCtx) error {
	country := helpers.Country{Id: helpers.DEF_ID}

	helpers.SetCountry(UserIp(ctx), &country)
	ctx.SetUserValue("country", country)
	logger.Info(country.Marshal())

	return ctx.Next()
}

func UserIp(ctx *atreugo.RequestCtx) string {
	IPAddress := string(ctx.Request.Header.Peek("X-Real-Ip"))

	if IPAddress == "" {
		IPAddress = string(ctx.Request.Header.Peek("X-Forwarded-For"))
	}

	if IPAddress == "" {
		IPAddress = ctx.RemoteAddr().String()
	}

	return IPAddress
}
