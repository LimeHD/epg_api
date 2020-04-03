package helpers

// Потом посмотрю, стоит ли выделать в отдельное репо

import (
	"encoding/json"
	"github.com/LimeHD/epg_api/utils"
)

const (
	ANDROID = "android"
	IOS     = "ios"
	SMART   = "smart"
	WEB     = "web"
)

const (
	SELF_ID    = 10
	WEB_ID     = 10
	ANDROID_ID = 20
	IOS_ID     = 30
	SMART_ID   = 40
)

type Platform struct {
	Id          int    `json:"id"`
	Platform    string `json:"platform"`
	App         string `json:"app"`
	VersionCode int    `json:"version_code"`
	VersionName string `json:"version_name"`
	DeviceId    string `json:"device_id"`
	Guid        string `json:"guid"`
	Name        string `json:"name"`
	SDK         int    `json:"sdk"`

	IsInitial bool
}

var platform_enum = map[string]int{
	ANDROID: ANDROID_ID,
	IOS:     IOS_ID,
	SMART:   SMART_ID,
	WEB:     WEB_ID,
}

func define(platformName string) (int, bool) {
	id, exist := platform_enum[platformName]

	return id, exist
}

// for debug
func (platform *Platform) Marshal() string {
	jsonString, _ := json.Marshal(platform)

	return string(jsonString)
}

func SetPlatform(userAgent []byte, v *Platform) {
	// не прошел user-agent, что же теперь делать..
	if err := json.Unmarshal(userAgent, &v); err != nil {
		v.Id = SELF_ID
		v.VersionCode = 1

		return
	}

	if id, ok := define(v.Platform); ok {
		v.Id = id
	}

	if v.Guid != "" {
		// fow smart tv
	}

	if v.Name != "" {
		v.Name, _ = utils.URLDecode(v.Name)
	}

	if !utils.Empty(v.SDK) {
		// sdk check
	}

	v.IsInitial = true
}
