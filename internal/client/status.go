package client

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/fatih/color"
	"github.com/gar-id/queued/internal/server/api/types"
	"github.com/gar-id/queued/tools"
	"github.com/olekukonko/tablewriter"
)

func QueuedStatus() {
	// Hit endpoint
	apiResponse := GetAPIServer("program")

	// Parse into general struct first
	var generalResult types.General
	json.Unmarshal(apiResponse, &generalResult)

	// Check http response
	if generalResult.HTTP_Code != 200 {
		tools.ZapLogger("console").Fatal(fmt.Sprintf("Error when retrieve process(s). Error message: %v", generalResult.Data.Message))
		return
	}

	// Parse into specific struct then
	var statusResult types.Program
	json.Unmarshal(apiResponse, &statusResult)

	// Output to formatting
	var valuesArray [][]string
	for programName, program := range statusResult.Data.Programs {
		for _, process := range program.Process {
			// Set uptime color
			if process.Status == "running" {
				process.Status = fmt.Sprintf(color.GreenString("%s"), process.Status)
			} else if process.Status == "stopped" || process.Status == "fatal" {
				process.Status = fmt.Sprintf(color.RedString("%s"), process.Status)
			} else {
				process.Status = fmt.Sprintf(color.YellowString("%s"), process.Status)
			}

			// Add to array
			if program.Group == "" {
				program.Group = "default"
			}
			valueConvert := []string{
				program.Group,
				programName,
				process.ProcessName,
				strconv.Itoa(*process.PID),
				process.Status,
				process.LastStart.String(),
				program.User,
				strconv.FormatBool(program.AutoStart),
				strconv.FormatBool(program.AutoRestart)}
			valuesArray = append(valuesArray, valueConvert)
		}
	}

	// Sort result
	sort.SliceStable(valuesArray, func(i, j int) bool {
		return valuesArray[i][2] < valuesArray[j][2]
	})

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Group name", "Program Name", "Process Name", "PID", "Status", "Last Start", "User", "Auto Start", "Auto Restart"})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)
	table.AppendBulk(valuesArray) // Add Bulk Data
	table.Render()
}
