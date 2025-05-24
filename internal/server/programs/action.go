package programs

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gar-id/queued/internal/server/config/caches"
	"github.com/gar-id/queued/internal/server/config/types"
	"github.com/gar-id/queued/tools"
)

func Validate(name, category, action string) (returnText []string) {
	// Check category
	switch category {
	case "group":
		processGroup, exist := caches.Data.GroupIndex[name]
		if exist {
			for _, processName := range processGroup {
				result := processAction(processName, action)
				returnText = append(returnText, result)
			}
		} else {
			result := fmt.Sprintf("group %v not found", name)
			returnText = append(returnText, result)
		}
	case "program":
		program, exist := caches.Data.ProgramConfig[name]
		if exist {
			for _, process := range program.Process {
				result := processAction(process.ProcessName, action)
				returnText = append(returnText, result)
			}
		} else {
			result := fmt.Sprintf("program %v not found", name)
			returnText = append(returnText, result)
		}
	case "process":
		_, exist := caches.ProcessChannel.Data[name]
		if exist {
			result := processAction(name, action)
			returnText = append(returnText, result)
		} else {
			result := fmt.Sprintf("process %v not found", name)
			returnText = append(returnText, result)
		}
	}
	return returnText
}

func processAction(processName, action string) (returnText string) {
	// Parse processname
	name := strings.Split(processName, ":")
	if len(name) == 3 {
		var newName []string
		newName = append(newName, name[1:]...)
		name = newName
	}
	processIndex, _ := strconv.Atoi(name[1])
	programName := &name[0]

	// Do action
	switch action {
	case "stop":
		if caches.Data.ProgramConfig[*programName].Process[processIndex].Status == types.ProcessStatusRunning {
			// Update status
			caches.Data.Do(func() {
				caches.Data.ProgramConfig[*programName].Process[processIndex].Status = types.ProcessStatusStopping
			})
		} else if caches.Data.ProgramConfig[*programName].Process[processIndex].Status == types.ProcessStatusStopping || caches.Data.ProgramConfig[*programName].Process[processIndex].Status == types.ProcessStatusStopped {
			returnText = fmt.Sprintf("%v is already %v", processName, caches.Data.ProgramConfig[*programName].Process[processIndex].Status)
			tools.ZapLogger("both").Info(returnText)
			return returnText
		} else if caches.Data.ProgramConfig[*programName].Process[processIndex].Status == types.ProcessStatusStarting {
			for caches.Data.ProgramConfig[*programName].Process[processIndex].Status != types.ProcessStatusRunning {
				time.Sleep(time.Millisecond)
			}
			// Update status
			caches.Data.Do(func() {
				caches.Data.ProgramConfig[*programName].Process[processIndex].Status = types.ProcessStatusStopping
			})

		}

		returnText = fmt.Sprintf("%v will be %vped", processName, action)
		tools.ZapLogger("both").Info(returnText)

		// Send channel to terminate program
		go func() {
			*caches.ProcessChannel.Data[processName].StopChannel <- true
		}()
		for caches.Data.ProgramConfig[*programName].Process[processIndex].Status != types.ProcessStatusStopped {
			time.Sleep(time.Millisecond)
		}

		returnText = fmt.Sprintf("%v is %vped", processName, action)
	case "restart":
		// Check current status
		if caches.Data.ProgramConfig[*programName].Process[processIndex].Status == types.ProcessStatusRunning {
			returnText = fmt.Sprintf("%v will be %ved", processName, action)
			tools.ZapLogger("both").Info(returnText)

			go func() {
				*caches.ProcessChannel.Data[processName].StopChannel <- true
			}()
			for caches.Data.ProgramConfig[*programName].Process[processIndex].Status != types.ProcessStatusStopped {
				time.Sleep(time.Millisecond)
			}

			returnText = fmt.Sprintf("%v is %ved", processName, action)
		} else if caches.Data.ProgramConfig[*programName].Process[processIndex].Status == types.ProcessStatusStopping {
			for caches.Data.ProgramConfig[*programName].Process[processIndex].Status != types.ProcessStatusStopped {
				time.Sleep(time.Millisecond)
			}
		} else if caches.Data.ProgramConfig[*programName].Process[processIndex].Status == types.ProcessStatusStarting {
			for caches.Data.ProgramConfig[*programName].Process[processIndex].Status != types.ProcessStatusRunning {
				time.Sleep(time.Millisecond)
			}
			// Update status
			caches.Data.Do(func() {
				caches.Data.ProgramConfig[*programName].Process[processIndex].Status = types.ProcessStatusStopping
			})
			returnText = fmt.Sprintf("%v will be %ved", processName, action)
			tools.ZapLogger("both").Info(returnText)
			returnText = fmt.Sprintf("%v is %ved", processName, action)
		} else if caches.Data.ProgramConfig[*programName].Process[processIndex].Status == types.ProcessStatusStopped {
			returnText = fmt.Sprintf("%v is not running, so we changes your action to start", processName)
			tools.ZapLogger("both").Warn(returnText)
			returnText = fmt.Sprintf("%v is %ved", processName, action)
		}

		// Run process
		caches.Data.Do(func() {
			caches.Data.ProgramConfig[*programName].Process[processIndex].Status = types.ProcessStatusStarting
		})
		go start(*programName, processName, processIndex)

		tools.ZapLogger("both").Info(returnText)
	case "start":
		// Check currenct status
		if caches.Data.ProgramConfig[*programName].Process[processIndex].Status == types.ProcessStatusRunning {
			returnText = fmt.Sprintf("%v is already %ved", processName, action)
			tools.ZapLogger("both").Info(returnText)
			return returnText
		} else if caches.Data.ProgramConfig[*programName].Process[processIndex].Status == types.ProcessStatusStopping {
			for caches.Data.ProgramConfig[*programName].Process[processIndex].Status != types.ProcessStatusStopped {
				time.Sleep(time.Millisecond)
			}
		} else if caches.Data.ProgramConfig[*programName].Process[processIndex].Status == types.ProcessStatusStarting {
			returnText = fmt.Sprintf("cannot start %v. %v is already %v",
				processName,
				processName,
				caches.Data.ProgramConfig[*programName].Process[processIndex].Status)
			tools.ZapLogger("both").Warn(returnText)
			return returnText
		}

		returnText = fmt.Sprintf("%v will be %ved", processName, action)
		tools.ZapLogger("both").Info(returnText)

		// Run program
		caches.Data.Do(func() {
			caches.Data.ProgramConfig[*programName].Process[processIndex].Status = types.ProcessStatusStarting
		})
		go start(*programName, processName, processIndex)

		returnText = fmt.Sprintf("%v is %ved", processName, action)
		tools.ZapLogger("both").Info(returnText)
	}
	return returnText
}
