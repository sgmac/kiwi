package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var tmpdir string

func setup() {
	tmpdir, _ = ioutil.TempDir("/tmp", "k-")
	appPath = filepath.Join(tmpdir, ".kiwi")

	err := os.Mkdir(appPath, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func cleanup() {
	os.RemoveAll(tmpdir)
}

func writeConfig(configFile string, data []byte) {

	filePath := filepath.Join(appPath, "config")
	_ = ioutil.WriteFile(filePath, data, 0644)
}
