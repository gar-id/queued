package config

import (
	"os"
	"path"

	"github.com/gar-id/queued/internal/general/config/caches"
	"github.com/gar-id/queued/internal/general/config/types"
	"github.com/gar-id/queued/tools"
	"gopkg.in/yaml.v2"
)

func LoadMainConfig(file_location string) {
	// Check server config location
	filePath := tools.DefaultString(file_location, path.Join("/", "etc", "queued", "config.yaml"))

	// Load server config file
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		tools.ZapLogger("console").Error(err.Error())
		var configEnv = types.MainConfig{
			QueueD: struct {
				API          types.ConfigAPI                     "json:\"api\" yaml:\"api\""
				Notification map[string]types.ConfigNotification "json:\"notification\" yaml:\"notification\""
				Log          types.ConfigLog                     "json:\"log\" yaml:\"log\""
				Include      string                              "json:\"include\" yaml:\"include\""
			}{
				API: types.ConfigAPI{
					HTTPListen:   os.Getenv("QUEUED_API_HTTPLISTEN"),
					Cors:         os.Getenv("QUEUED_API_CORS"),
					AuthEnabled:  tools.StrBool(os.Getenv("QUEUED_API_AUTH_ENABLED")),
					AuthAdmin:    os.Getenv("QUEUED_API_AUTH_ADMIN"),
					AuthReadOnly: os.Getenv("QUEUED_API_AUTH_READONLY"),
				},
				Log: types.ConfigLog{
					Level:    os.Getenv("QUEUED_LOG_LEVEL"),
					Location: os.Getenv("QUEUED_LOG_LOCATION"),
				},
				Include: os.Getenv("QUEUED_INCLUDE"),
			},
		}

		// Save config to cache
		caches.MainConfig = configEnv
		return
	}

	// Parse server config file into struct
	err = yaml.Unmarshal(yamlFile, &caches.MainConfig)
	if err != nil {
		tools.ZapLogger("console").Fatal(err.Error())
	}
}
