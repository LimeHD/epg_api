package entries

type Category struct {
	Id         int    `json:"id"`
	Identifier string `json:"identifier"`
	NameRu     string `json:"name_ru"`
	Sort       int    `json:"sort"`
}
