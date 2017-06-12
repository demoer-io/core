package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/demoer-io/demoer/api"
	"github.com/demoer-io/demoer/repository"
)

func main() {
	app := cli.NewApp()

	app.Name = "demoer"
	app.Version = "0.0.0"
	app.Usage = "demoer core runner"

	app.Commands = []cli.Command{
		repository.Command,
		api.Command,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
