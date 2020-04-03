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
	exampleResponse := []entries.Channel{}

	exampleResponse = append(exampleResponse, entries.Channel{
		OurId:  105,
		NameRu: "Первый канал",
		NameEn: "Perviy Kanal",
	})

	exampleResponse = append(exampleResponse, entries.Channel{
		OurId:  115,
		NameRu: "Россия 1",
		NameEn: "Rossiya 1",
	})

	return ctx.JSONResponse(exampleResponse)
}
