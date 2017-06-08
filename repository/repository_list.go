package repository

import (
	"fmt"

	"github.com/urfave/cli"
)

// provider flag of the repositoriesList command
var provider string

var repositoriesList = cli.Command{
	Name:   "ls",
	Usage:  "Repositories list",
	Action: list,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "provider, p",
			Usage:       "Repositories provider (github, gitlab, ...)",
			Destination: &provider,
		},
	},
}

func list(c *cli.Context) error {
	var repositories []Repository
	var err error

	switch provider {
	case "github":
		var github GitHub
		repositories, err = github.getAllRepositories()
	default:
		panic(fmt.Sprintf("Repository provider %v not supported !", provider))
	}

	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(repositories)

	return nil
}
