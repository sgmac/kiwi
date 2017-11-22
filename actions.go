package main

import (
	"fmt"
	"io"
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

func stop(c *cli.Context) error {
	resp, err := client.Stop()
	if err != nil {
		logrus.Fatalf("stop-%s\n", err)
	}

	switch resp.Error {
	case 0:
		prettylines(os.Stdout, "OK")
	default:
		prettylines(os.Stdout, "NOK")

	}
	return nil
}

func start(c *cli.Context) error {
	resp, err := client.Start()
	if err != nil {
		logrus.Fatalf("start-%s\n", err)
	}

	switch resp.Error {
	case 0:
		prettylines(os.Stdout, "OK")
	default:
		prettylines(os.Stderr, "NOK")

	}
	return nil
}

func kill(c *cli.Context) error {
	resp, err := client.Kill()
	if err != nil {
		logrus.Fatalf("kill-%s\n", err)
	}

	switch resp.Error {
	case 0:
		prettylines(os.Stdout, "OK")
	default:
		prettylines(os.Stderr, "NOK")
	}
	return nil
}

func reboot(c *cli.Context) error {
	resp, err := client.Reboot()
	if err != nil {
		logrus.Fatalf("reboot-%s\n", err)
	}

	switch resp.Error {
	case 0:
		prettylines(os.Stdout, "OK")
	default:
		prettylines(os.Stderr, "NOK")
	}
	return nil
}

func hostname(c *cli.Context) error {
	var host string
	if c.NArg() < 1 {
		prettylines(os.Stderr, "Provide a hostname.")
		return nil
	}
	host = c.Args().Get(0)

	resp, err := client.Hostname(host)
	if err != nil {
		logrus.Fatalf("hostname-%s\n", err)
	}

	switch resp.Error {
	case 0:
		prettylines(os.Stdout, "OK")
	default:
		prettylines(os.Stderr, "NOK")

	}
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

func prettylines(w io.Writer, msg string) {
	fmt.Fprintf(w, "%s\n", msg)
}
