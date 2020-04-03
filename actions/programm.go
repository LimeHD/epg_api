package actions

import (
	"github.com/LimeHD/epg_api/entries"
	"github.com/savsgio/atreugo/v11"
)

// ProgrammAction godoc
// @Summary Show TV programm list
// @Description get string by ID
// @ID get-string-by-int
// @Accept json
// @Produce json
// @Param id path int true "Channel ID"
// @Router /channels/{id}/programm [get]
// @Success 200 {array} entries.ProgrammResponse "ok"
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
