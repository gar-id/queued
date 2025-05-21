package process

import (
	"fmt"
	"os"
	"path"

	generalCaches "github.com/gar-id/queued/internal/general/config/caches"
	"github.com/gar-id/queued/internal/server/config"
	"github.com/gar-id/queued/internal/server/config/caches"
	"github.com/gar-id/queued/internal/server/config/types"
	"github.com/gar-id/queued/internal/server/programs"
	"github.com/gar-id/queued/tools"
	"gopkg.in/yaml.v2"

	"github.com/gofiber/fiber/v2"
)

func ConfigAdd(c *fiber.Ctx) (result error) {
	// Parse POST request
	var configPayloads = make(map[string]types.ProgramConfig)
	var statusMessage = make(map[string][]string)
	var httpStatus string
	err := c.BodyParser(&configPayloads)
	if err != nil {
		return c.JSON(General(fiber.StatusBadRequest, "error", c.IP(), err.Error()))
	}

	// Check POST payload
	for configName, configPayload := range configPayloads {
		// Check command value
		if configPayload.Command == "" {
			statusMessage[configName] = append(statusMessage[configName], "there's no value for key 'command'. this program will not be processed")
			httpStatus = "error"
			continue
		}

		// Check is config already exist
		_, exist := caches.Data.ProgramConfig[configName]
		if exist {
			for _, process := range caches.Data.ProgramConfig[configName].Process {
				statusMessage[configName] = append(statusMessage[process.ProcessName], fmt.Sprintf("program for program %v already exist and current status is %v", configName, process.Status))
			}
			httpStatus = "error"
			continue
		}

		// Convert struct into yaml format
		var configStruct = make(map[string]types.ProgramConfig)
		configStruct[configName] = configPayload
		yamlFile, err := yaml.Marshal(configStruct)
		if err != nil {
			statusMessage[configName] = append(statusMessage[configName], fmt.Sprintf("error when parsing config for program %v. error message: %v", configName, err.Error()))
			httpStatus = "error"
			continue
		}

		// Print config into file
		fileDestination := path.Join(generalCaches.MainConfig.QueueD.Include, fmt.Sprintf("%v.yml", configName))
		configFile, err := os.OpenFile(fileDestination, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			tools.ZapLogger("file").Fatal(err.Error())
			statusMessage[configName] = append(statusMessage[configName], fmt.Sprintf("error when create config file for program %v. error message: %v", configName, err.Error()))
			httpStatus = "error"
		}
		configFile.WriteString(string(yamlFile))
		configFile.Close()

		// Save to cache
		config.LoadProgramConfig(configName, configPayload)

		// Run program if autostart is true
		if caches.Data.ProgramConfig[configName].AutoStart {
			statusMessage[configName] = append(statusMessage[configName], programs.Validate(configName, "program", "start")...)
		}
	}

	// Return
	if httpStatus == "error" {
		return c.JSON(GeneralObject(fiber.StatusBadRequest, "error", c.IP(), statusMessage))
	}
	return c.JSON(GeneralObject(fiber.StatusBadRequest, "error", c.IP(), statusMessage))
}
