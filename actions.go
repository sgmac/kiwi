package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/Sirupsen/logrus"
	"github.com/sgmac/bandwagon"
	cli "gopkg.in/urfave/cli.v1"
)

var client *bandwagon.Client

func init() {
	config, err := readConfig()
	if err != nil {
		log.Fatalf("readConfig-%s\n", err)
	}

	creds := bandwagon.Credentials{
		VeID:   config.VPS.VeID,
		APIKey: config.VPS.APIKey,
	}

	client = bandwagon.NewClient(creds)
}

func listimages(c *cli.Context) error {
	images, err := client.ListImages()
	if err != nil {
		logrus.Fatalf("listimages-%s\n", err)
	}
	prettyOutput(images)
	return nil
}

func prettyOutput(v interface{}) {
	w := tabwriter.NewWriter(os.Stdout, 40, 8, 0, ' ', 0)
	var header string

	switch v {
	case v.(*bandwagon.Images):
		imgs := v.(*bandwagon.Images)
		header = "IMAGE"
		fmt.Fprintf(w, "%s\n", header)
		for _, img := range imgs.Templates {
			fmt.Fprintf(w, "%s\n", img)
		}

	default:
	case v.(*bandwagon.InfoVPS):
		fmt.Fprintf(w, "images %v", v)
	}

	w.Flush()
}
