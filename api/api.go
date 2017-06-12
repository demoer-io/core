package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/urfave/cli"
)

// Command to manage API server
var Command = cli.Command{
	Name:  "api",
	Usage: "Manage API serve",
	Subcommands: []cli.Command{
		cli.Command{
			Name:   "start",
			Usage:  "Start API server",
			Action: start,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "port",
					Value: "8080",
					Usage: "API server port default 8080",
				},
			},
		},
	},
}

func start(c *cli.Context) error {
	// fmt.Sprintf(":%s", c.String("port"))

	r := Router()

	// Headers
	alowedHeaders := handlers.AllowedHeaders([]string{"Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%s", c.String("port")),
		handlers.LoggingHandler(
			os.Stdout,
			handlers.CORS(alowedHeaders, allowedOrigins)(r)),
	))

	return nil
}
