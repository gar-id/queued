package tools

import (
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

func BulkLoadINI(filelocation string) (*ini.File, error) {
	configdir, err := os.ReadDir(filelocation)
	var configlists []string

	// init empty ini loader
	uptimeConfig := ini.Empty()

	if len(configdir) < 1 {
		return uptimeConfig, err
	}

	// read every files in config dir
	for _, dir := range configdir {
		fileconfig := strings.ReplaceAll(dir.Name(), "- ", "")
		configTemp := filelocation + "/" + fileconfig
		configlists = append(configlists, configTemp)
		uptimeConfig.Append(configTemp)
	}

	// load every file config
	if err != nil {
		return uptimeConfig, err
	}

	return uptimeConfig, err
}

func SingleLoadINI(filelocation string) (*ini.File, error) {
	// init empty ini loader
	uptimeConfig, err := ini.LoadSources(ini.LoadOptions{
		SkipUnrecognizableLines: true,
	}, filelocation)

	// load every file config
	if err != nil {
		return uptimeConfig, err
	}

	return uptimeConfig, err
}
