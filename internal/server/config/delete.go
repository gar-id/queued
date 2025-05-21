package config

import (
	"fmt"
	"os"

	generalCaches "github.com/gar-id/queued/internal/general/config/caches"
	"github.com/gar-id/queued/internal/server/config/caches"
	programTypes "github.com/gar-id/queued/internal/server/config/types"
	"github.com/gar-id/queued/internal/server/programs"
	"github.com/gar-id/queued/tools"
	"gopkg.in/yaml.v2"
)

func DeleteConfig(ProgramName []string) (returnMessage map[string][]string, httpStatus string, err error) {
	var statusMessage = make(map[string][]string)
	// Check process name and execute if exist
	for _, programName := range ProgramName {
		programConfig, ok := caches.Data.ProgramConfig[programName]
		if ok {
			caches.Data.Do(func() {
				programConfig.AutoRestart = false
				caches.Data.ProgramConfig[programName] = programConfig
			})
			// Check every process in program. If running, then stop
			for order := range caches.Data.ProgramConfig[programName].Process {
				if caches.Data.ProgramConfig[programName].Process[order].Status == "running" {
					programs.Validate(caches.Data.ProgramConfig[programName].Process[order].ProcessName, "process", "stop")
					statusMessage[programName] = append(statusMessage[programName], fmt.Sprintf("process %v is stopped and will be deleted", caches.Data.ProgramConfig[programName].Process[order].ProcessName))

					// Delete channel
					caches.ProcessChannel.Do(func() {
						<-*caches.ProcessChannel.Data[caches.Data.ProgramConfig[programName].Process[order].ProcessName].StopChannel
						delete(caches.ProcessChannel.Data, caches.Data.ProgramConfig[programName].Process[order].ProcessName)
					})
				} else {
					statusMessage[programName] = append(statusMessage[programName], fmt.Sprintf("process %v is deleted", caches.Data.ProgramConfig[programName].Process[order].ProcessName))

					caches.ProcessChannel.Do(func() {
						delete(caches.ProcessChannel.Data, caches.Data.ProgramConfig[programName].Process[order].ProcessName)
					})
				}

			}

			// Delete program from config file
			// Search file location
			filePath, yamlFile, err := SearchProgramFile(generalCaches.MainConfig.QueueD.Include, programName)
			if err != nil {
				httpStatus = "error"
				statusMessage[programName] = append(statusMessage[programName], err.Error())
			}

			// Parse program config file into struct
			var tempProgram map[string]programTypes.ProgramConfig
			err = yaml.Unmarshal(yamlFile, &tempProgram)
			if err != nil {
				tools.ZapLogger("console").Error(err.Error())
				statusMessage[programName] = append(statusMessage[programName], err.Error())
				httpStatus = "error"
				continue
			}

			// Truncate program
			delete(tempProgram, programName)
			yamlFile, _ = yaml.Marshal(tempProgram)

			// Delete program
			programFile, err := os.Create(filePath)
			if err != nil {
				tools.ZapLogger("file").Error(err.Error())
				statusMessage[programName] = append(statusMessage[programName], err.Error())
				httpStatus = "error"
				continue
			}
			programFile.Write(yamlFile)
			programFile.Close()

			// Delete data from caches
			caches.Data.Do(func() {
				delete(caches.Data.ProgramConfig, programName)
			})
			statusMessage[programName] = append(statusMessage[programName], fmt.Sprintf("program %v is deleted", programName))
		} else {
			httpStatus = "error"
			statusMessage[programName] = append(statusMessage[programName], fmt.Sprintf("program %v not found", programName))
		}
	}

	return returnMessage, httpStatus, err
}
