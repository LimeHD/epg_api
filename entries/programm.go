package entries

import (
	"database/sql"
	"github.com/LimeHD/epg_api/service"
	"github.com/LimeHD/epg_api/utils"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"time"
)

type Epg struct {
	EpgId     int            `json:"epg_id"`
	Title     string         `json:"title"`
	Desc      string         `json:"desc"`
	Cdnvideo  int            `json:"cdnvideo"`
	Rating    sql.NullString `json:"rating"`
	Timestart int64          `json:"timestart"`
	Timestop  int64          `json:"timestop"`
}

type Programm struct {
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Time        string `json:"time"`
	Begin       int64  `json:"begin"`
	End         int64  `json:"end"`
	Current     bool   `json:"current"`
	Rating      string `json:"rating"`
	Url         string `json:"url"`
	AspectRatio string `json:"aspect_ratio"`
}

type ProgrammResponse struct {
	Title  string     `json:"title"`
	Name   string     `json:"name"`
	Data   []Programm `json:"data"`
	Active bool       `json:"active"`
}

func GetProgramm(id int, tz int, curdate int) map[string]*ProgrammResponse {
	nowStamp := time.Now().Unix()
	nowDay := utils.YearMonthDay(nowStamp)

	prgrammResponse := make(map[string]*ProgrammResponse)
	var epgModels []Epg

	addSecs := utils.ResolveByTimezone(tz)

	ok, channel := FindOnePlaylist(id)
	valid, epgId := channel.GetEpgId(tz)

	// if channel is not exist on database
	if ok == false {
		return prgrammResponse
	}

	if valid == false && epgId == 0 {
		return prgrammResponse
	}

	err := service.GetInstance().Database.
		Select("unix_timestamp(timestart) as timestart",
			"unix_timestamp(timestop) as timestop", "title", "desc", "cdnvideo", "rating",
		).
		From("epg").
		Where(dbx.HashExp{"epg_id": epgId}).
		OrderBy("date", "timestart").
		All(&epgModels)

	// o my god, records, where records
	if err != nil {
		// TODO off error or handle
		panic(err)
	}

	for _, model := range epgModels {
		day := utils.YearMonthDay(model.Timestart)

		if _, ok = prgrammResponse[day]; !ok {
			prgrammResponse[day] = &ProgrammResponse{}
		}

		extendedStart := model.Timestart + addSecs
		extendedStop := model.Timestop + addSecs
		current := extendedStart <= nowStamp && extendedStop >= nowStamp

		prgrammResponse[day].Title = day
		prgrammResponse[day].Name = utils.MonthDayWeekName(extendedStart)
		prgrammResponse[day].Data = append(prgrammResponse[day].Data, Programm{
			Title:       model.Title,
			Desc:        model.Desc,
			Time:        utils.AsTime(extendedStart),
			Begin:       model.Timestart,
			End:         model.Timestop,
			Current:     current,
			Rating:      model.Rating.String,
			Url:         "",
			AspectRatio: channel.AspectRatio,
		})
	}

	// if current day epg -- to mark as active sheet on devices (clients)
	if _, ok := prgrammResponse[nowDay]; ok {
		prgrammResponse[nowDay].Active = true
		prgrammResponse[nowDay].Name = "Today"
	}

	return prgrammResponse
}
