package entries

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/LimeHD/epg_api/constants"
	"github.com/LimeHD/epg_api/entries/channel"
	"github.com/LimeHD/epg_api/entries/user"
	"github.com/LimeHD/epg_api/helpers"
	"github.com/LimeHD/epg_api/service"
	"github.com/LimeHD/epg_api/utils"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"google.golang.org/appengine/memcache"
	"time"
)

type (
	ChannelResponse struct {
		Key        string
		Channels   map[int]Channel  `json:"channels"`
		Categories map[int]Category `json:"categories"`
		Genres     []Genre          `json:"genres"`
		PaidPacks  []int            `json:"paid_packs"`

		UserPacks []user.UserPacks `json:"user_packs"`
	}

	Channel struct {
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

		PlaylistUrl      ChannelUrl
		PlaylistCategory struct{}
		PlaylistCountry  channel.PlaylistCountry
		PlaylistPack     channel.PlaylistPack
		Epg              Epg

		NearestModel Nearest
	}

	ChannelUrl struct {
		Id            int
		PlaylistOurId int
		UrlProtocol   string
		UrlArchive    sql.NullString
		UrlSound      sql.NullString
		UrlStuff      string
		Tz            int
		EpgId         sql.NullInt32
	}

	Nearest struct {
		Origin               string
		Cdnvideo             string
		Quality              []int
		EpgId                int
		PlaylistUrlQualities []int
	}
)

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

// common methods

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

// main

func GetPlaylist(tz int, platform *helpers.Platform, country *helpers.Country, user *user.Identity) *ChannelResponse {
	playlist := &ChannelResponse{}
	playlist.attachCategories().attachGenres()

	if user.GetIsGuest() == false {
		playlist.attachUserPacks(user.GetID())
	}

	var models []Channel

	err = service.GetInstance().Database.Select(
		"p.our_id", "p.name_ru", "p.is_foreign", "p.public", "p.sort", "p.with_archive",
		"playlist_url.url_protocol", "playlist_url.tz", "playlist_url.epg_id",
		"playlist_url.url_archive", "playlist_url.url_sound",

		"playlist_country.show_after_purchase",

		"playlist_packs.pack_id",
	).
		Distinct(true).
		From("playlist p").
		LeftJoin("playlist_country", dbx.NewExp("p.our_id=playlist_country.playlist_our_id")).
		LeftJoin("playlist_category", dbx.NewExp("p.our_id=playlist_category.playlist_our_id")).
		LeftJoin("playlist_url", dbx.NewExp("p.our_id=playlist_url.playlist_our_id")).
		LeftJoin("playlist_packs", dbx.NewExp("p.our_id=playlist_packs.playlist_our_id")).
		Where(dbx.HashExp{
			"p.status":                      constants.PLAYLIST_STATUS_ACTIVE,
			"p.public":                      constants.PLAYLIST_PUBLIC,
			"playlist_country.platform_id":  platform.Id,
			"playlist_country.country_id":   country.Id,
			"playlist_category.category_id": 4, // four is lite
		}).
		OrderBy("sort ASC").
		All(&models)

	if err != nil {
		panic(err)
	}

	imageUrlProtocol := channel.GetCDNStorageUrl(platform.CheckAndroidSDK() && platform.LessThanAvailable())
	playlist.Channels = make(map[int]Channel)

	for _, model := range models {
		model.ConstructedUrl(tz)

		if model.HasEpg() {
			// another logic
		}

		playlist.Channels[model.OurId] = Channel{
			Id:                model.OurId,
			EpgId:             model.PlaylistUrl.EpgId.Int32,
			NameRu:            model.NameRu,
			Url:               model.GetUrl(model.Cdnvideo == 1),
			Cdn:               model.GetCDNUrl(),
			UrlArchive:        model.GetUrlArchive(),
			UrlSound:          model.GetUrlSound(),
			WithArchive:       model.IsWithArchive(),
			ShowAfterPurchase: model.ShowOnlyAfterPurchase(),
			Image:             model.GetImageURL(imageUrlProtocol),
			Current:           model.GetCurrentEpg(),
			Quality:           model.GetQualities(),
		}

		if platform.Id == helpers.WEB_ID {
			//playlist.Channels[model.OurId].DescriptionEn = model.DescriptionEn
		}
	}

	return playlist
}

func (c *Channel) ConstructedUrl(timezone int) {
	now := utils.EqualLimeTimeFromPHP()
	nowStamp := time.Now().Unix()
	nowTime := nowStamp + constants.EXPIRE_TTL

	if c.IsForeign == 1 {
		c.NearestModel.Origin = c.ForeignUrl
		c.NearestModel.Cdnvideo = c.ForeignUrl

		// хз зачем это
		c.NearestModel.Quality = []int{1}

		if c.Tvprogram == 1 && c.ForeignEpgId > 0 {
			err = service.GetInstance().Database.
				Select("unix_timestamp(timestart) as timestart",
					"unix_timestamp(timestop) as timestop", "title", "desc", "cdnvideo", "rating",
				).
				Where(dbx.NewExp(fmt.Sprintf("timestart >= '%s'", now))).
				AndWhere(dbx.NewExp(fmt.Sprintf("timestop <= '%s'", now))).
				AndWhere(dbx.HashExp{"epg_id": c.ForeignEpgId}).
				One(&c.Epg)

			if err != nil {
				panic(err)
			}

			c.NearestModel.EpgId = c.ForeignEpgId
		}

		return
	}

	c.FindPlaylistByTZ(timezone)
	c.AttachAvailableQualities()

	if c.PlaylistUrl.UrlProtocol != "" {
		c.NearestModel.Cdnvideo = c.PlaylistUrl.UrlCDN(channel.CDNUrl(c.PlaylistUrl.PlaylistOurId), "md5here", "expires", nowTime)
		c.NearestModel.Origin = c.PlaylistUrl.UrlOrigin(channel.GetProtectedHost(false), "hashExample", nowTime)
	}

	if c.Tvprogram == 1 && c.Epg.EpgId > 0 {
		// что тут дальше уууу...
	}
}

// own methods

var err error

func (c *Channel) AttachAvailableQualities() {
	c.NearestModel.Quality = nil
	var qualities []channel.PlaylistQualities

	quO1 := map[int]int{}

	err = service.GetInstance().Database.
		Select("quality_id").
		From("playlist_url_quality").
		Where(dbx.HashExp{"playlist_url_id": c.PlaylistUrl.Id}).
		All(&qualities)

	if err != nil {
		panic(err)
	}

	for _, q := range qualities {
		quO1[q.QualityId] = q.QualityId
	}

	for _, q := range GetQualities() {
		if _, ok := quO1[q.Id]; ok {
			c.NearestModel.Quality = append(c.NearestModel.Quality, q.Id)
		}
	}
}

func (p *ChannelResponse) attachGenres() *ChannelResponse {
	err = service.GetInstance().Database.Select("id", "name_ru").
		From("genre").
		Where(dbx.HashExp{"type": constants.TYPE_PLAYLIST}).
		OrderBy("name_ru ASC").
		All(&p.Genres)

	if err != nil {
		panic(err)
	}

	return p
}

func (p *ChannelResponse) attachCategories() *ChannelResponse {
	var categories []Category
	p.Categories = make(map[int]Category)

	err = service.GetInstance().Database.Select("id", "identifier", "name_ru", "sort").
		From("category").
		Where(dbx.NewExp("id != 4")).AndWhere(dbx.NewExp("identifier != 'lite'")).
		OrderBy("sort ASC").
		All(&categories)

	for _, category := range categories {
		p.Categories[category.Id] = category
	}

	if err != nil {
		panic(err)
	}

	p.Categories[9999] = Category{
		Id:         9999,
		Identifier: "favorite",
		NameRu:     "Избранные",
		Sort:       1000,
	}

	return p
}

func (p *ChannelResponse) attachUserPacks(userId int) *ChannelResponse {
	err = service.GetInstance().Database.Select("user_packs.*").
		Distinct(true).
		From("playlist_packs").
		//LeftJoin("packs", dbx.NewExp("packs.id=playlist_packs.pack_id")).
		LeftJoin("user_packs", dbx.NewExp("playlist_packs.pack_id=user_packs.pack_id")).
		Where(dbx.HashExp{
			"playlist_packs.status": constants.PLAYLIST_STATUS_ACTIVE,
			"user_packs.status":     constants.PACK_STATUS_PAID,
			"user_packs.user_id":    userId,
		}).
		AndWhere(dbx.NewExp(fmt.Sprintf("user_packs.end > %d", time.Now().Unix()))).
		All(&p.UserPacks)

	if err != nil {
		panic(err)
	}

	fmt.Println("USER PACKS SELECT")
	return p
}
