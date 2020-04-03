package main

import (
	"fmt"
	"github.com/LimeHD/epg_api/actions"
	"github.com/LimeHD/epg_api/middlewares"
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

		// temporary debug flags
		fmt.Println(fmt.Sprintf("Full flags: host: %s, port: %d \nmysql://%s:%s@%s:%d/%s",
			host, port, dbuser, dbpass, dbhost, dbport, dbname))

		// todo run through unix socket
		config := atreugo.Config{
			Addr: fmt.Sprintf("%s:%d", host, port),
		}

		server := atreugo.New(config)
		server.UseBefore(middlewares.UserAgent)
		server.Path("GET", "/channels", actions.ChannelsAction)
		server.Path("GET", "/channels/{id}/programm", actions.ProgrammAction)

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
