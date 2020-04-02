package main

import (
	"github.com/LimeHD/epg_api/entries"
	"github.com/savsgio/atreugo/v11"
)

func main() {
	config := atreugo.Config{
		Addr: "127.0.0.1:8000",
	}

	server := atreugo.New(config)
	server.Path("GET", "/channels", ChannelsAction)
	server.Path("GET", "/channels/{id}/programm", ProgrammAction)

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func ChannelsAction(ctx *atreugo.RequestCtx) error {
	exampleResponse := []entries.Channel{}

	exampleResponse = append(exampleResponse, entries.Channel{
		OurId:  105,
		NameRu: "Первый канал",
		NameEn: "Perviy Kanal",
	})

	exampleResponse = append(exampleResponse, entries.Channel{
		OurId:  115,
		NameRu: "Россия 1",
		NameEn: "Rossiya 1",
	})

	return ctx.JSONResponse(exampleResponse)
}

func ProgrammAction(ctx *atreugo.RequestCtx) error {
	//id := ctx.UserValue("id").(string)

	exampleResponse := []entries.ProgrammResponse{}
	exampleResponse = append(exampleResponse, entries.ProgrammResponse{
		Title: "27.03.2020",
		Name:  "27.03 ПТ",
		Data: []entries.Programm{
			{
				Title: "Проверено на себе",
				Desc:  "Что делать, если зима выдалась тёплая, а счёт за...",
			},
			{
				Title: "Время покажет",
				Desc:  "В студии программы обсуждают то, что волнует каждого из нас...",
			},
			{
				Title: "Наедине со всеми",
				Desc:  "Очень часто журналисты задают стандартные вопросы...",
			},
		},
	})

	exampleResponse = append(exampleResponse, entries.ProgrammResponse{
		Title: "28.03.2020",
		Name:  "28.03 СБ",
		Data: []entries.Programm{
			{
				Title: "Майлз Дэвис: Рождение нового джаза",
				Desc:  "Биография музыканта, собранная из прежде неизданных архивных материалов и...",
			},
			{
				Title: "Мужское/Женское",
				Desc:  "Громкие истории, в которых осталось немало домыслов, вопросов и...",
			},
			{
				Title: "Про любовь",
				Desc:  "Несмотря на фундаментальные разногласия в вопросах любви...",
			},
		},
	})

	return ctx.JSONResponse(exampleResponse)
}
