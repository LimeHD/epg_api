package helpers

import (
	"github.com/LimeHD/epg_api/utils"
	"math"
	"sort"
)

// Example 0 is MSK time or GMT +3
// The rest follow the pattern +number by MSK, second element is "2" => +2 MSK => +5 GMT
// This description of a simple function serves to prevent confusion.
func GetAllTimezones() []int {
	return []int{0, 2, 3, 4, 6, 7, 8}
}

func GetNearestTimezones(timezone int) []int {
	// The user sends in GMT, the database is stored in MSK time, WTF!? I do not understand...
	timezone -= 3
	zones := map[int]int{}

	for _, tz := range GetAllTimezones() {
		zone := int(math.Abs(float64(timezone - tz)))

		if !utils.ContainsMap(zones, zone) {
			zones[zone] = tz
		}
	}

	var sortZones []int
	keys := utils.IntMapKeys(zones)
	sort.Ints(keys)

	for _, k := range keys {
		sortZones = append(sortZones, zones[k])
	}

	return sortZones
}

func GetTimezoneByValues(tz int, msk int) int {
	if tz <= 0 || msk == 1 {
		return 3
	}

	return tz
}
