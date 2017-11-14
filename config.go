package main

import (
	"fmt"
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

type tomlConfig struct {
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

func readConfig() error {
	config := filepath.Join(appPath, filename)
	if _, err := os.Stat(config); os.IsNotExist(err) {
		f, err := os.Create(config)
		if err != nil {
			return err
		}

		server := tomlConfig{}
		tml := toml.NewEncoder(f)
		err = tml.Encode(server)
		if err != nil {
			return err
		}
	}
	return nil
}
