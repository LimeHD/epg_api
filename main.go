package main

import (
	"fmt"
	"github.com/LimeHD/epg_api/entries"
	"github.com/savsgio/atreugo/v11"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Value: "127.0.0.1",
				Usage: "run in host",
			},
			&cli.IntFlag{
				Name:  "port",
				Value: 8001,
				Usage: "run in port",
			},

			// database
			&cli.StringFlag{
				Name:  "dbhost",
				Value: "localhost",
				Usage: "Database host",
			},
			&cli.StringFlag{
				Name:  "dbuser",
				Value: "root",
				Usage: "Database user",
			},
			&cli.StringFlag{
				Name:  "dbpass",
				Value: "",
				Usage: "Database password",
			},
			&cli.StringFlag{
				Name:  "dbname",
				Value: "db",
				Usage: "Database name",
			},
			&cli.IntFlag{
				Name:  "dbport",
				Value: 9000,
				Usage: "Database port",
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		host := c.String("host")
		port := c.Int("port")
		dbhost := c.String("dbhost")
		dbuser := c.String("dbuser")
		dbpass := c.String("dbpass")
		dbname := c.String("dbname")
		dbport := c.Int("dbport")

		fmt.Println(fmt.Sprintf("Full flags: host: %s, port: %d \nmysql://%s:%s@%s:%d/%s",
			host, port, dbuser, dbpass, dbhost, dbport, dbname))

		// todo run through unix socket
		config := atreugo.Config{
			Addr: fmt.Sprintf("%v:%v", host, port),
		}

		server := atreugo.New(config)
		server.Path("GET", "/channels", ChannelsAction)
		server.Path("GET", "/channels/{id}/programm", ProgrammAction)

		err := server.ListenAndServe()

		if err != nil {
			panic(err)
		}
		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
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
