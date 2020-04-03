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
	OurId         int
	NameRu        string
	NameEn        string
	IsForeign     int
	ForeignUrl    string
	ForeignEpgId  int
	Public        int
	DayArchive    int
	WithArchive   int
	Tvprogram     int
	Cdnvideo      int
	DescriptionEn string
	DescriptionRu string
	Image         string
	AspectRatio   string

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
		panic(err)
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
		return false, pl
	}

	return true, pl
}
