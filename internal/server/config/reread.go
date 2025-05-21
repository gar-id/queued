package config

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/gar-id/queued/internal/server/config/caches"
	"github.com/gar-id/queued/internal/server/config/types"
	"github.com/gar-id/queued/internal/server/programs"
	"github.com/gar-id/queued/tools"
	"gopkg.in/yaml.v2"
)

func RereadConfig(programLocation string) (returnMessage map[string][]string, err error) {
	// Check directory, and return error if empty
	programLocation = tools.DefaultString(programLocation, path.Join("/", "etc", "queued", "conf.d"))
	programDir, err := os.ReadDir(programLocation)
	if len(programDir) < 1 {
		return nil, err
	}

	// Get every files in config dir
	var tempProgramConfig = make(map[string]types.ProgramConfig)
	var statusMessage = make(map[string][]string)
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
				tempProgramConfig[programName] = program
			}
		}
	}

	// Print error if exist
	if err != nil {
		tools.ZapLogger("console").Error(err.Error())
		return nil, err
	}

	// // Debug
	// jsonNew, _ := json.Marshal(tempProgramConfig)
	// jsonOld, _ := json.Marshal(caches.ProgramConfig)
	// fmt.Println(string(jsonNew))
	// fmt.Println(string(jsonOld))

	// Compare to current running program
	for programName, tempProgram := range tempProgramConfig {
		cacheConfig, exist := caches.Data.ProgramConfig[programName]
		if exist {
			if tempProgram.AutoRestart != cacheConfig.AutoRestart ||
				tempProgram.AutoStart != cacheConfig.AutoStart ||
				tempProgram.SlowStart != cacheConfig.SlowStart ||
				tempProgram.StartSecs != cacheConfig.StartSecs ||
				tempProgram.NumProcs != cacheConfig.NumProcs ||
				tempProgram.Command != cacheConfig.Command ||
				tempProgram.Stderr != cacheConfig.Stderr ||
				tempProgram.Stdout != cacheConfig.Stdout ||
				tempProgram.Group != cacheConfig.Group ||
				tempProgram.User != cacheConfig.User {
				// If process for program $programName is running, then stop. But if not running, then check autostart for program
				for order, process := range caches.Data.ProgramConfig[programName].Process {
					if caches.Data.ProgramConfig[programName].Process[order].Status == "running" {
						statusMessage[programName] = append(statusMessage[programName], programs.Validate(process.ProcessName, "process", "stop")...)
					}
				}

				// If process for program $programName is not stopped, then wait until stopped
				var deleteConfig bool = false
				for !deleteConfig {
					deleteConfig = true
					for order := range caches.Data.ProgramConfig[programName].Process {
						if caches.Data.ProgramConfig[programName].Process[order].Status != "stopped" {
							deleteConfig = false
						}
					}
				}

				// Delete old config
				if deleteConfig {
					caches.Data.Do(func() {
						delete(caches.Data.ProgramConfig, programName)
					})
					// Save to cache
					LoadProgramConfig(programName, tempProgram)
				}

				if caches.Data.ProgramConfig[programName].AutoStart {
					programs.Validate(programName, "program", "start")
					statusMessage[programName] = append(statusMessage[programName], fmt.Sprintf("%v is updated and started", programName))
				} else {
					statusMessage[programName] = append(statusMessage[programName], fmt.Sprintf("%v is updated", programName))
				}
			}
		} else {
			// If program is new, then add to cache and start if autostart is true
			LoadProgramConfig(programName, tempProgram)

			// Run command
			if caches.Data.ProgramConfig[programName].AutoStart {
				programs.Validate(programName, "program", "start")
				statusMessage[programName] = append(statusMessage[programName], fmt.Sprintf("%v is added and started", programName))
			} else {
				statusMessage[programName] = append(statusMessage[programName], fmt.Sprintf("%v is added", programName))
			}
		}
	}

	// Check and delete config if doesn't exist in new program config
	for programName, programConfig := range caches.Data.ProgramConfig {
		_, exist := tempProgramConfig[programName]
		if !exist {
			caches.Data.Do(func() {
				programConfig.AutoRestart = false
				caches.Data.ProgramConfig[programName] = programConfig
			})
			// If process for program $programName is running, then stop. But if not running, then check autostart for program
			for order, process := range caches.Data.ProgramConfig[programName].Process {
				if caches.Data.ProgramConfig[programName].Process[order].Status == "running" {
					statusMessage[programName] = append(statusMessage[programName], programs.Validate(process.ProcessName, "process", "stop")...)
				}
			}

			// If process for program $programName is not stopped, then wait until stopped
			var deleteConfig bool = false
			for !deleteConfig {
				deleteConfig = true
				for order := range caches.Data.ProgramConfig[programName].Process {
					if caches.Data.ProgramConfig[programName].Process[order].Status != "stopped" {
						deleteConfig = false
					}
				}
			}

			// Delete old config
			if deleteConfig {
				caches.Data.Do(func() {
					delete(caches.Data.ProgramConfig, programName)
				})
			}
			statusMessage[programName] = append(statusMessage[programName], fmt.Sprintf("program %v is deleted", programName))
		}
	}

	// // Debug
	// jsonNew, _ = json.Marshal(statusMessage)
	// fmt.Println(string(jsonNew))
	return statusMessage, err
}
