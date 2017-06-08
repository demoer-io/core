package repository

import (
	"github.com/urfave/cli"
)

// Repository model
type Repository struct {
	ID   int
	Name string
}

// Provider interface for (GitHub, GitLab, ...)
type Provider interface {
	getRepositoriesList() ([]Repository, error)
}

// Command to manage repositories
var Command = cli.Command{
	Name:    "repo",
	Aliases: []string{"rp"},
	Usage:   "manage repositories",
	Subcommands: []cli.Command{
		repositoriesList,
	},
}
