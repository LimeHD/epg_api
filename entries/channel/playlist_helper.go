package channel

import (
	"fmt"
	"github.com/LimeHD/epg_api/constants"
	"github.com/LimeHD/epg_api/utils"
	"strings"
)

func GetProtectedHost(isHttpCondition bool) string {
	if isHttpCondition {
		return constants.HOST_PROTECTED_HTTP
	}

	return constants.HOST_PROTECTED
}

func getCDNHost(isHttpCondition bool) string {
	if isHttpCondition {
		return constants.HOST_CDN_VIDEO_HTTP
	}

	return constants.HOST_CDN_VIDEO
}

func GetArchiveHost(isHttpCondition bool) string {
	if isHttpCondition {
		return constants.HOST_ARCHIVE_HTTP
	}

	return constants.HOST_ARCHIVE
}

func GetCDNStorageUrl(isHttpCondition bool) string {
	str := "https://pica.iptv2021.com/"

	if isHttpCondition {
		return strings.Replace(str, "https", "http", -1)
	}

	return str
}

func CDNUrl(e int) string {
	if _, ok := utils.ContainsIntSlice([]int{666, 777}, e); ok {
		return constants.HOST_SDN_CORE_HTTP
	}

	return constants.HOST_SDN_CORE
}

func GetPlaylistCacheKey(id int, country int, tz int) string {
	return fmt.Sprintf("playlist_%d_%d_%d", id, country, tz)
}
