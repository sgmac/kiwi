package main

import (
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

const Usage = "Manage your Bandwagon VPS"

func main() {
	app := cli.NewApp()
	app.Name = "kiwi"
	app.Usage = Usage
	app.Run(os.Args)
}
