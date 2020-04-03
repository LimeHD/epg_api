package actions

import (
	"github.com/LimeHD/epg_api/entries"
	"github.com/LimeHD/epg_api/helpers"
	"github.com/LimeHD/epg_api/utils"
	"github.com/savsgio/atreugo/v11"
	"sort"
	"strconv"
	"time"
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
	var response interface{}

	programm := entries.GetProgramm(id, helpers.GetTimezoneByValues(tz, msk), curdate)

	// return empty array
	if len(programm) == 0 {
		return ctx.JSONResponse(make([]string, 0))
	}

	// check only current day returned
	if curdate == 1 {
		nowDay := utils.YearMonthDay(time.Now().Unix())

		response = make([]string, 0)

		if v, ok := programm[nowDay]; ok {
			response = v
		}

		return ctx.JSONResponse(response)
	}

	v := []*entries.ProgrammResponse{}

	// make flat
	for _, value := range programm {
		v = append(v, value)
	}

	sort.Slice(v[:], func(i, j int) bool {
		return v[i].Title < v[j].Title
	})

	return ctx.JSONResponse(v)
}
