package entries

import (
	"database/sql"
	"fmt"
	"github.com/LimeHD/epg_api/helpers"
	"github.com/LimeHD/epg_api/service"
	"github.com/LimeHD/epg_api/utils"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type Channel struct {
	OurId         int    `json:"our_id"`
	NameRu        string `json:"name_ru"`
	NameEn        string `json:"name_en"`
	IsForeign     int    `json:"is_foreign"`
	ForeignUrl    string `json:"foreign_url"`
	ForeignEpgId  int    `json:"foreign_epg_id"`
	Public        int    `json:"public"`
	DayArchive    int    `json:"day_archive"`
	WithArchive   int    `json:"with_archive"`
	Tvprogram     int    `json:"tvprogram"`
	Cdnvideo      int    `json:"cdnvideo"`
	DescriptionEn string `json:"description_en"`
	DescriptionRu string `json:"description_ru"`
	Image         string `json:"image"`
	AspectRatio   string `json:"aspect_ratio"`

	PlaylistUrl ChannelUrl
}

type ChannelUrl struct {
	Id            int
	PlaylistOurId int
	UrlProtocol   string
	UrlArchive    sql.NullString
	UrlSound      sql.NullString
	UrlStuff      string
	Tz            int
	EpgId         sql.NullInt32
}

func GetChannelList() []Channel {
	var channels []Channel

	err := service.GetInstance().Database.Select("p.our_id", "p.name_ru", "p.name_en").
		From("playlist p").
		OrderBy("sort ASC").
		All(&channels)

	if err != nil {
		// todo bugsnag
		panic(err)
	}

	return channels
}

func (playlist *Channel) GetEpgId(tz int) (bool, int) {
	if playlist.OurId == 0 {
		return false, 0
	}

	if playlist.IsForeign > 0 {
		return true, playlist.ForeignEpgId
	}

	playlist.FindPlaylistByTZ(tz)

	return playlist.PlaylistUrl.EpgId.Valid, int(playlist.PlaylistUrl.EpgId.Int32)
}

func (playlist *Channel) FindPlaylistByTZ(tz int) {
	tzKeys := helpers.GetNearestTimezones(tz)

	err := service.GetInstance().Database.
		Select("playlist_url.*").
		From("playlist_url").
		Where(dbx.HashExp{"playlist_url.playlist_our_id": playlist.OurId}).
		AndWhere(dbx.NewExp("playlist_url.url_protocol != ''")).
		OrderBy(fmt.Sprintf("field(tz, %s)", utils.ArrayToString(tzKeys, ","))).
		Limit(1).
		One(&playlist.PlaylistUrl)

	if err != nil {
		service.GetInstance().BugsnagNotifier.Notify(err)
	}
}

func FindOnePlaylist(id int) (bool, Channel) {
	var pl Channel

	err := service.GetInstance().Database.Select("our_id", "name_ru", "public").
		From("playlist").
		Where(dbx.HashExp{"our_id": id}).
		Limit(1).
		One(&pl)

	if err != nil {
		service.GetInstance().BugsnagNotifier.Notify(err)
		return false, pl
	}

	return true, pl
}
