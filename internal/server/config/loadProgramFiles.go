package config

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/gar-id/queued/internal/server/config/caches"
	"github.com/gar-id/queued/internal/server/config/types"
	"github.com/gar-id/queued/tools"
	"gopkg.in/yaml.v2"
)

func BulkLoadYaml(programLocation string) (err error) {
	// Check directory, and return error if empty
	programLocation = tools.DefaultString(programLocation, path.Join("/", "etc", "queued", "conf.d"))
	programDir, err := os.ReadDir(programLocation)
	if len(programDir) < 1 {
		return err
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
			for programName, program := range tempProgram {
				// Save to cache
				LoadProgramConfig(programName, program)
			}
		}
	}

	// Print error if exist
	if err != nil {
		tools.ZapLogger("console").Fatal(err.Error())
	}

	return err
}

func LoadProgramConfig(programName string, program types.ProgramConfig) {
	// Create child process if more than 1
	if program.NumProcs == 1 {
		var processConfig types.ProcessConfig
		// If group is empty, then simplify process name
		if program.Group == "" {
			processConfig.ProcessName = fmt.Sprintf("%v:%v", programName, 0)
		} else {
			processConfig.ProcessName = fmt.Sprintf("%v:%v:%v", program.Group, programName, 0)

			// Add process to $group group
			caches.Data.Do(func() {
				caches.Data.GroupIndex[program.Group] = append(caches.Data.GroupIndex[program.Group], processConfig.ProcessName)
			})
		}

		processConfig.ProcessIndex = 0
		processConfig.ProgramName = programName
		program.Process = append(program.Process, processConfig)

		// Add process to default group
		caches.Data.Do(func() {
			caches.Data.GroupIndex["default"] = append(caches.Data.GroupIndex["default"], processConfig.ProcessName)
		})
	} else if program.NumProcs > 1 {
		for order := 0; order < program.NumProcs; order++ {
			// If group is empty, then simplify process name
			var processConfig types.ProcessConfig
			if program.Group == "" {
				processConfig.ProcessName = fmt.Sprintf("%v:%v", programName, order)
			} else {
				processConfig.ProcessName = fmt.Sprintf("%v:%v:%v", program.Group, programName, order)

				// Add process to $group group
				caches.Data.Do(func() {
					caches.Data.GroupIndex[program.Group] = append(caches.Data.GroupIndex[program.Group], processConfig.ProcessName)
				})
			}
			processConfig.ProcessIndex = order
			processConfig.ProgramName = programName
			program.Process = append(program.Process, processConfig)

			// Add process to default group
			caches.Data.Do(func() {
				caches.Data.GroupIndex["default"] = append(caches.Data.GroupIndex["default"], processConfig.ProcessName)
			})
		}
	}

	// // Debug config
	// yamlConfig, _ := yaml.Marshal(program)
	// fmt.Println(string(yamlConfig))

	caches.Data.Do(func() {
		caches.Data.ProgramConfig[programName] = program
	})
}
