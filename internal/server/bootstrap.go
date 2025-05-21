package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gar-id/queued/internal/general/config/caches"
	generalCaches "github.com/gar-id/queued/internal/general/config/caches"
	"github.com/gar-id/queued/internal/server/api"
	"github.com/gar-id/queued/internal/server/config"
	serverCaches "github.com/gar-id/queued/internal/server/config/caches"
	"github.com/gar-id/queued/internal/server/programs"
	"github.com/gar-id/queued/tools"
)

func Bootstrap() {
	// Setup for graceful termination
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	// Show banner
	fmt.Print(tools.Banner())

	// Start HTTP API
	go api.Start()
	tools.ZapLogger("both").Info(fmt.Sprintf("Starting QueueD RestAPI Server with listen to %v", caches.MainConfig.QueueD.API.HTTPListen))

	// Load programs config
	config.BulkLoadYaml(generalCaches.MainConfig.QueueD.Include)

	// Start program
	for program := range serverCaches.Data.ProgramConfig {
		// Start program if autostart is true
		if serverCaches.Data.ProgramConfig[program].AutoStart {
			programs.Validate(program, "program", "start")
		}
	}

	// listen for the interrupt signal
sigTerm:
	for range done {
		tools.ZapLogger("both").Warn("SIGTERM called: QueueD process will be terminated")
		if len(serverCaches.Data.ProgramConfig) == 0 {
			break sigTerm
		}

		// stop all programs
		for programName := range serverCaches.Data.ProgramConfig {
			serverCaches.Data.Do(func() {
				programConfig := serverCaches.Data.ProgramConfig[programName]
				programConfig.AutoRestart = false
				serverCaches.Data.ProgramConfig[programName] = programConfig
			})
			programs.Validate(programName, "program", "stop")
		}

		// If process for program $programName is not stopped, then wait until stopped
		var stopQueued bool = false
		for programName := range serverCaches.Data.ProgramConfig {
			for !stopQueued {
				stopQueued = true
				for order := range serverCaches.Data.ProgramConfig[programName].Process {
					if serverCaches.Data.ProgramConfig[programName].Process[order].Status != "stopped" {
						stopQueued = false
					}
				}
			}
		}
		if stopQueued {
			break sigTerm
		}
	}

	// Wait then exit
	serverCaches.Data.WaitGroup.Wait()
	os.Exit(0)

}
