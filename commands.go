package main

import (
	cli "gopkg.in/urfave/cli.v1"
)

func commands() []cli.Command {
	return []cli.Command{
		{
			Name:    "images",
			Aliases: []string{"i"},
			Usage:   "List available images.",
			Action:  listimages,
		},
	}
}
