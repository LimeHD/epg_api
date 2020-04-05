package helpers

import (
	"encoding/json"
	"github.com/LimeHD/epg_api/service"
	"github.com/LimeHD/epg_api/utils"
	"github.com/yl2chen/cidranger"
	"net"
	"strings"
)

const (
	DEF_ID = 40

	LIME_ID   = 10
	RU_ID     = 20
	UA_ID     = 30
	SNG_ID    = 35
	WORLD_ID  = 40
	BALTIA_ID = 50
)

const (
	DEF = "WORLD"

	LIME   = "LIME"
	RU     = "RU"
	UA     = "UA"
	AZ     = "AZ"
	AM     = "AM"
	BY     = "BY"
	KZ     = "KZ"
	KG     = "KG"
	MD     = "MD"
	TJ     = "TJ"
	UZ     = "UZ"
	TM     = "TM"
	GE     = "GE"
	LV     = "LV"
	LT     = "LT"
	EE     = "EE"
	SNG    = "СНГ"
	BALTIA = "Страны Балтии"
)

type Country struct {
	Id          int
	Ip          string
	City        string `json:"city"`
	Subdivision interface{}
	Country     string `json:"country"`
	Timezone    string
	IsoCode     string
	Coords      []float64

	Net string `json:"net"`
}

func (country *Country) Get() int {
	return country.Id
}

func (country *Country) Set(id int, network string) {
	country.Id = id
	country.Net = network
}

func (country *Country) Marshal() string {
	jsonString, err := json.Marshal(country)

	if err != nil {
		panic(err)
	}

	return string(jsonString)
}

var (
	officeNetworks = []string{
		"85.234.0.0/19",
		"127.0.0.1/32",
		"5.158.232.0/21",
		"185.24.44.0/22",
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/24",
	}
)

func SetCountry(userIpOrCoord string, country *Country) {
	if userIpOrCoord == "" || strings.Index(userIpOrCoord, "127.0.0.1") != -1 {
		country.Set(LIME_ID, LIME)
		return
	}

	ip := net.ParseIP(userIpOrCoord)
	record, err := service.GetInstance().GeoReader.City(ip)
	ranger := cidranger.NewPCTrieRanger()

	if err != nil {
		for _, network := range officeNetworks {
			if ipCIDRCheck(ranger, network, userIpOrCoord) {
				country.Set(LIME_ID, LIME)
				return
			}
		}

		country.Set(RU_ID, RU)
		return
	}

	country.Country = record.Country.Names["ru"]
	country.City = record.City.Names["ru"]
	country.Timezone = record.Location.TimeZone
	country.IsoCode = record.Country.IsoCode
	country.Coords = []float64{record.Location.Latitude, record.Location.Longitude}
	country.Subdivision = record.Subdivisions

	if country.IsoCode == RU {
		for _, network := range officeNetworks {
			if ipCIDRCheck(ranger, network, userIpOrCoord) {
				country.Set(LIME_ID, LIME)
				return
			}
		}

		country.Set(RU_ID, RU)
		return
	}

	if country.IsoCode == UA {
		if record.Subdivisions[0].IsoCode == "40" || record.Subdivisions[0].IsoCode == "43" {
			country.Set(RU_ID, RU)
			return
		}

		country.Set(UA_ID, UA)
		return
	}

	if _, find := utils.In(getSNGCodes(), record.Country.IsoCode); find == true {
		country.Set(SNG_ID, SNG)
		return
	}

	if _, find := utils.In(getBaltiaCodes(), record.Country.IsoCode); find == true {
		country.Set(BALTIA_ID, BALTIA)
		return
	}

	country.Set(DEF_ID, DEF)
	return
}

func ipCIDRCheck(ranger cidranger.Ranger, network string, ip string) bool {
	_, network1, _ := net.ParseCIDR(network)

	if err := ranger.Insert(cidranger.NewBasicRangerEntry(*network1)); err != nil {
		panic(err)
	}

	contains, err := ranger.Contains(net.ParseIP(ip))

	if err != nil {
		panic(err)
	}

	return contains
}

func getSNGCodes() []string {
	return []string{AZ, AM, BY, KZ, KG, MD, TJ, UZ, TM, GE}
}

func getBaltiaCodes() []string {
	return []string{LT, EE, LV}
}
