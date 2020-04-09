package main

import (
	"fmt"
	"github.com/LimeHD/epg_api/actions"
	"github.com/LimeHD/epg_api/helpers"
	"github.com/LimeHD/epg_api/middlewares"
	"github.com/LimeHD/epg_api/service"
	"github.com/savsgio/atreugo/v11"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host epg.iptv2021.com
// @BasePath /v1
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
				Value: "@",
				Usage: "Database host",
			},
			&cli.StringFlag{
				Name:  "dbuser",
				Value: "",
				Usage: "Database user",
			},
			&cli.StringFlag{
				Name:  "dbpass",
				Value: "",
				Usage: "Database password",
			},
			&cli.StringFlag{
				Name:  "dbname",
				Value: "",
				Usage: "Database name",
			},
		},
	}

	// go run . --dbuser root2 --dbpass prizrak211 --dbname limehd --dbhost @
	app.Action = func(c *cli.Context) error {
		host := c.String("host")
		port := c.Int("port")
		dbhost := c.String("dbhost")
		dbuser := c.String("dbuser")
		dbpass := c.String("dbpass")
		dbname := c.String("dbname")

		service.GetInstance().ConnectDatabase(helpers.GetDbConnectionString(dbuser, dbpass, dbhost, dbname))
		defer service.GetInstance().Database.Close()

		// todo run through unix socket
		config := atreugo.Config{
			Addr: fmt.Sprintf("%s:%d", host, port),
		}

		server := atreugo.New(config)
		server.UseBefore(middlewares.UserAgent)
		server.Path("GET", "/channels", actions.ChannelsAction)
		server.Path("GET", "/channels/{id}/programm", actions.ProgrammAction)

		// todo нужно перенести ресурсы js, css и иконки на cdn и останется только swagger.yaml  индексной страничкой
		static := server.NewGroupPath("/docs")
		static.ServeFile("/swagger", "docs/index.html")
		static.ServeFile("/swagger-ui.css", "docs/assets/swagger-ui.css")
		static.ServeFile("/uifavicon-32x32.png", "docs/assets/uifavicon-32x32.png")
		static.ServeFile("/uifavicon-16x16.png", "docs/assets/uifavicon-16x16.png")
		static.ServeFile("/swagger-ui-bundle.js", "docs/assets/swagger-ui-bundle.js")
		static.ServeFile("/swagger-ui-standalone-preset.js", "docs/assets/swagger-ui-standalone-preset.js")
		static.ServeFile("/swagger.yaml", "docs/swagger.yaml")

		err := server.ListenAndServe()

		if err != nil {
			log.Fatal(err)
		}
		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
