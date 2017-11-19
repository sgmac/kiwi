package main

import (
	"fmt"
	"log"
	"os"

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

func info(c *cli.Context) error {
	info, err := client.Info()
	if err != nil {
		logrus.Fatalf("info-%s\n", err)
	}
	prettyOutput(info)
	return nil
}

func prettyOutput(v interface{}) {
	var header string

	switch t := v.(type) {

	case *bandwagon.Images:
		imgs := v.(*bandwagon.Images)
		header = "IMAGE"
		fmt.Fprintf(os.Stdout, "%s\n", header)
		for _, img := range imgs.Templates {
			fmt.Fprintf(os.Stdout, "%s\n", img)
		}

	case *bandwagon.InfoVPS:
		info := v.(*bandwagon.InfoVPS)
		header = "VIRTUALIZATION\tHOSTNAME\tOS\t\t\t\tIP"
		fmt.Fprintf(os.Stdout, "%s\n", header)
		fmt.Fprintf(os.Stdout, "%s\t\t%s\t\t%s\t\t%s\n", info.VMType, info.Hostname, info.OS, info.IPAddresses)
	default:
		_ = t
	}
}
