package main

import (
	cli "gopkg.in/urfave/cli.v1"
)

func commands() []cli.Command {
	return []cli.Command{
		{
			Name:    "images",
			Aliases: []string{"i"},
			Usage:   "List OS images.",
			Action:  listimages,
		},
		{
			Name:    "status",
			Aliases: []string{"f"},
			Usage:   "Status VPS.",
			Action:  info,
		},
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "Start VPS.",
			Action:  start,
		},
		{
			Name:    "stop",
			Aliases: []string{"t"},
			Usage:   "Stop VPS.",
			Action:  stop,
		},
		{
			Name:    "kill",
			Aliases: []string{"k"},
			Usage:   "Kill VPS.",
			Action:  kill,
		},
		{
			Name:    "reboot",
			Aliases: []string{"r"},
			Usage:   "Reboot VPS.",
			Action:  reboot,
		},
		{
			Name:    "hostname",
			Aliases: []string{"h"},
			Usage:   "Set hostname.",
			Action:  hostname,
		},
		{
			Name:    "install",
			Aliases: []string{"l"},
			Usage:   "Install OS.",
			Action:  install,
		},
	}
}
