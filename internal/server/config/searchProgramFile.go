package config

import (
	"os"
	"path"
	"strings"

	"github.com/gar-id/queued/internal/server/config/types"
	"github.com/gar-id/queued/tools"
	"gopkg.in/yaml.v2"
)

func SearchProgramFile(programLocation, programName string) (filePath string, yamlFile []byte, err error) {
	// Check directory, and return error if empty
	programLocation = tools.DefaultString(programLocation, path.Join("/", "etc", "queued", "conf.d"))
	programDir, err := os.ReadDir(programLocation)
	if len(programDir) < 1 {
		return "", nil, err
	}

	// Get every files in config dir
	for _, dir := range programDir {
		getName, found := strings.CutSuffix(dir.Name(), "yaml")
		if !found {
			getName, _ = strings.CutSuffix(dir.Name(), "yml")
		}

		// Make sure to parse only .yaml or .yml files
		if strings.HasSuffix(getName, ".") {
			// Open program config file
			var tempProgram map[string]types.ProgramConfig
			filePath := path.Join(programLocation, dir.Name())
			yamlFile, err := os.ReadFile(filePath)
			if err != nil {
				tools.ZapLogger("console").Fatal(err.Error())
			}

			// Parse program config file into struct
			err = yaml.Unmarshal(yamlFile, &tempProgram)
			if err != nil {
				tools.ZapLogger("console").Fatal(err.Error())
			}

			// Append to program config and program channel
			for programNameSearch := range tempProgram {
				if programName == programNameSearch {
					return filePath, yamlFile, nil
				}
			}
		}
	}

	// Print error if exist
	if err != nil {
		tools.ZapLogger("console").Fatal(err.Error())
	}

	return "", nil, err
}
