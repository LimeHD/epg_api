package actions

import (
	"github.com/LimeHD/epg_api/entries"
	"github.com/LimeHD/epg_api/helpers"
	"github.com/LimeHD/epg_api/utils"
	"github.com/savsgio/atreugo/v11"
	"strconv"
)

// ProgrammAction godoc
// @Summary Show TV programm list
// @Description get string by ID
// @ID get-string-by-int
// @Accept json
// @Produce json
// @Param id path int true "Channel ID"
// @Router /channels/{id}/programm [get]
// @Success 200 {array} entries.ProgrammResponse "ok"
func ProgrammAction(ctx *atreugo.RequestCtx) error {
	id, _ := strconv.Atoi(ctx.UserValue("id").(string))
	tz := utils.ByteToInt(ctx.QueryArgs().Peek("tz"))
	msk := utils.ByteToInt(ctx.QueryArgs().Peek("msk"))
	curdate := utils.ByteToInt(ctx.QueryArgs().Peek("curdate"))

	programm, ok := entries.GetProgramm(id, helpers.GetTimezoneByValues(tz, msk), curdate)

	if !ok {
		return ctx.JSONResponse(make([]string, 0))
	}

	if curdate == 1 {
		if today := programm.GetToday(); today != nil {
			return ctx.JSONResponse(today)
		}

		return ctx.JSONResponse(make([]string, 0))
	}

	return ctx.JSONResponse(programm.MakeFlat())
}
