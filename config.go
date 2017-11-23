package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const (
	dirPerms = 0755
	filename = "config"
)

var (
	appPath = filepath.Join(os.Getenv("HOME"), ".kiwi")
)

type configServ struct {
	VPS serverInfo
}

type serverInfo struct {
	VeID   string
	APIKey string
}

func configPath() error {
	if _, err := os.Stat(appPath); os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, "Initializing ~/.kiwi")
		err := os.Mkdir(appPath, dirPerms)
		if err != nil {
			return err
		}
	}
	return nil
}

func readConfig() (*configServ, error) {
	config := filepath.Join(appPath, filename)
	var server *configServ
	if _, err := os.Stat(config); os.IsNotExist(err) {
		err := configPath()
		if err != nil {
			return nil, err
		}

		f, err := os.Create(config)
		if err != nil {
			return nil, err
		}

		server = &configServ{}
		tml := toml.NewEncoder(f)
		err = tml.Encode(server)
		if err != nil {
			return nil, err
		}
	}

	data, err := ioutil.ReadFile(config)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(data, &server)
	if err != nil {
		return nil, err
	}

	return server, nil
}
