package entries

import "github.com/LimeHD/epg_api/service"

type Quality struct {
	Id     int
	NameRu string
	Sort   int
}

func GetQualities() []Quality {
	var qualities []Quality

	err = service.GetInstance().Database.
		Select("id").
		From("quality").
		All(&qualities)

	if err != nil {
		panic(err)
	}

	return qualities
}
