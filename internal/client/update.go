package client

import (
	"encoding/json"
	"fmt"

	"github.com/gar-id/queued/internal/server/api/types"
	"github.com/gar-id/queued/tools"
)

func QueuedUpdate() {
	// Hit endpoint
	apiResponse := PostAPIServer(nil, "config/update")

	// Parse into general struct first
	var generalResult types.General
	json.Unmarshal(apiResponse, &generalResult)

	// Check http response
	if generalResult.HTTP_Code != 200 {
		tools.ZapLogger("console").Fatal(fmt.Sprintf("Error when retrieve process(s). Error message: %v", string(apiResponse)))
		return
	}

	// Parse into specific struct then
	var statusResult types.GeneralObject
	json.Unmarshal(apiResponse, &statusResult)

	// Output to formatting
	for _, messages := range statusResult.Data.Message {
		for _, message := range messages {
			tools.ZapLogger("console").Info(message)
		}
	}
}
