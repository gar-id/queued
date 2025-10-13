package client

import (
	"encoding/json"
	"fmt"

	"github.com/gar-id/queued/internal/server/api/types"
	"github.com/gar-id/queued/tools"
)

func QueuedAction(groupName, programName, processName []string, action string) {
	// Create payload
	var apiPayload = types.ActionPayload{
		GroupName:   groupName,
		ProgramName: programName,
		ProcessName: processName,
	}
	jsonPayload, err := json.Marshal(apiPayload)
	if err != nil {
		tools.ZapLogger("console").Fatal(fmt.Sprintf("Error when creating payload to server. Error message: %v", err.Error()))
		return
	}

	// Hit endpoint
	endpoint := fmt.Sprintf("program/%v", action)
	apiResponse := PostAPIServer(jsonPayload, endpoint)

	// Parse into object struct
	var statusResult types.GeneralObject
	json.Unmarshal(apiResponse, &statusResult)

	// Check http response
	if statusResult.HTTP_Code != 200 {
		for group, messages := range statusResult.Data.Message {
			for _, message := range messages {
				tools.ZapLogger("console").Fatal(fmt.Sprintf("Error when retrieve process(s). Error message: %v: %v", group, message))
			}
		}
		return
	}

	// Output to formatting
	for _, program := range statusResult.Data.Message {
		for _, message := range program {
			tools.ZapLogger("console").Info(message)
		}
	}
}
