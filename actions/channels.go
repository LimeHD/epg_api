package actions

import (
	"github.com/LimeHD/epg_api/entries"
	"github.com/savsgio/atreugo/v11"
)

// Channels godoc
// @Summary Show list of all channels
// @Description get string by ID
// @ID get-string-by-int
// @Accept json
// @Produce json
// @Router /channels [get]
// @Success 200 {array} entries.Channel "ok"
func ChannelsAction(ctx *atreugo.RequestCtx) error {
	channelResponse := entries.GetChannelList()
	return ctx.JSONResponse(channelResponse)
}
