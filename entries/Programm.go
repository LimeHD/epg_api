package entries

type Programm struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type ProgrammResponse struct {
	Title string     `json:"title"`
	Name  string     `json:"name"`
	Data  []Programm `json:"data"`
}
