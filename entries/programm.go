package entries

import (
	"database/sql"
	"github.com/LimeHD/epg_api/service"
	"github.com/LimeHD/epg_api/utils"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"sort"
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

type DaysResponse struct {
	Days map[string]*ProgrammResponse
}

func (dr *DaysResponse) DayExist(day string) bool {
	_, ok := dr.Days[day]

	return ok
}

func (dr *DaysResponse) AppendProgramm(day string, programm Programm) {
	dr.Days[day].Data = append(dr.Days[day].Data, programm)
}

func (dr *DaysResponse) CreateDay(day string, start int64) {
	dr.Days[day] = &ProgrammResponse{}
	dr.Days[day].Title = day
	dr.Days[day].Name = utils.MonthDayWeekName(start)
}

func (dr *DaysResponse) MarkToday(now int64) {
	nowDay := utils.YearMonthDay(now)

	// if current day epg -- to mark as active sheet on devices (clients)
	if _, ok := dr.Days[nowDay]; ok {
		dr.Days[nowDay].Active = true
		dr.Days[nowDay].Name = "Today"
	}
}

func (dr *DaysResponse) MakeFlat() []*ProgrammResponse {
	flatten := []*ProgrammResponse{}

	for _, value := range dr.Days {
		flatten = append(flatten, value)
	}

	sort.Slice(flatten[:], func(i, j int) bool {
		return flatten[i].Title < flatten[j].Title
	})

	return flatten
}

func (dr *DaysResponse) GetToday() *ProgrammResponse {
	nowDay := utils.YearMonthDay(time.Now().Unix())

	if v, ok := dr.Days[nowDay]; ok {
		return v
	}

	return nil
}

func GetProgramm(id int, tz int, curdate int) (*DaysResponse, bool) {
	nowStamp := time.Now().Unix()

	days := &DaysResponse{}
	days.Days = make(map[string]*ProgrammResponse)

	var epgModels []Epg

	addSecs := utils.ResolveByTimezone(tz)

	ok, channel := FindOnePlaylist(id)
	valid, epgId := channel.GetEpgId(tz)

	// if channel is not exist on database
	if ok == false {
		return days, false
	}

	if valid == false && epgId == 0 {
		return days, false
	}

	err := service.GetInstance().Database.
		Select("unix_timestamp(timestart) as timestart",
			"unix_timestamp(timestop) as timestop", "title", "desc", "cdnvideo", "rating",
		).
		From("epg").
		Where(dbx.HashExp{"epg_id": epgId}).
		OrderBy("date", "timestart").
		All(&epgModels)

	if err != nil {
		return days, false
	}

	for _, model := range epgModels {
		day := utils.YearMonthDay(model.Timestart)

		extendedStart := model.Timestart + addSecs
		extendedStop := model.Timestop + addSecs
		isActiveProgramm := extendedStart <= nowStamp && extendedStop >= nowStamp

		if exist := days.DayExist(day); !exist {
			days.CreateDay(day, extendedStart)
		}

		days.AppendProgramm(day, Programm{
			Title:       model.Title,
			Desc:        model.Desc,
			Time:        utils.AsTime(extendedStart),
			Begin:       model.Timestart,
			End:         model.Timestop,
			Current:     isActiveProgramm,
			Rating:      model.Rating.String,
			Url:         "",
			AspectRatio: channel.AspectRatio,
		})
	}

	days.MarkToday(nowStamp)

	return days, true
}
