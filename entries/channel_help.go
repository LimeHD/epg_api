package entries

import "fmt"

func (c *Channel) GetImageURL(imageUrl string) string {
	return imageUrl + c.Image
}

func (c *Channel) GetUrl(isCDN bool) string {
	if isCDN == true {
		return c.NearestModel.Cdnvideo
	}

	return c.NearestModel.Origin
}

func (c *Channel) GetUrlArchive() string {
	if c.WithArchive == 0 {
		return ""
	}

	return c.PlaylistUrl.UrlArchive.String
}

func (c *Channel) GetUrlSound() string {
	if c.GetUrlArchive() == "" {
		return ""
	}

	return c.PlaylistUrl.UrlSound.String
}

func (c *Channel) IsWithArchive() bool {
	if c.GetUrlArchive() != "" {
		return c.WithArchive == 1
	}

	return false
}

func (c *Channel) GetCDNUrl() string {
	return c.NearestModel.Cdnvideo
}

func (c *Channel) ShowOnlyAfterPurchase() int32 {
	return c.PlaylistCountry.ShowAfterPurchase.Int32
}

func (c *Channel) GetCurrentEpg() Epg {
	return c.Epg
}

func (c *Channel) HasEpg() bool {
	return c.GetCurrentEpg().EpgId > 0
}

func (c *Channel) GetQualities() []int {
	return c.NearestModel.Quality
}

func (plUrl *ChannelUrl) UrlCDN(host string, md5 string, eType string, e int64) string {
	return fmt.Sprintf("%s/streaming/%s/324/%s?md5=%s&%s=%d", host, plUrl.UrlProtocol, plUrl.UrlStuff, md5, eType, e)
}

func (plUrl *ChannelUrl) UrlOrigin(host string, hash string, time int64) string {
	return fmt.Sprintf("%s/p/%s,%d/streaming/%s/324/%s", host, hash, time, plUrl.UrlProtocol, plUrl.UrlStuff)
}
