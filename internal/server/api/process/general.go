package process

import (
	"fmt"
	"time"

	"github.com/gar-id/queued/internal/general/config/caches"
	"github.com/gar-id/queued/internal/server/api/types"
	programTypes "github.com/gar-id/queued/internal/server/config/types"
)

func Welcome(httpCode int, statusText, clientIP, message string) types.Welcome {
	var result = types.Welcome{
		HTTP_Code: httpCode,
		Status:    statusText,
		ClientIP:  clientIP,
		Data: struct {
			Date    string "json:\"date\""
			Message string "json:\"message\""
			Version string "json:\"version\""
		}{
			Date:    fmt.Sprint(time.Now().String()),
			Message: message,
			Version: caches.Version,
		}}

	return result
}

func General(httpCode int, statusText, clientIP, message string) types.General {
	var result = types.General{
		HTTP_Code: httpCode,
		Status:    statusText,
		ClientIP:  clientIP,
		Data: struct {
			Date    string "json:\"date\""
			Message string "json:\"message\""
		}{
			Date:    fmt.Sprint(time.Now().String()),
			Message: message,
		}}

	return result
}

func GeneralObject(httpCode int, statusText, clientIP string, message map[string][]string) types.GeneralObject {
	var result = types.GeneralObject{
		HTTP_Code: httpCode,
		Status:    statusText,
		ClientIP:  clientIP,
		Data: struct {
			Date    string              "json:\"date\""
			Message map[string][]string "json:\"message\""
		}{
			Date:    fmt.Sprint(time.Now().String()),
			Message: message,
		}}

	return result
}

func Program(httpCode int, statusText, clientIP string, program map[string]programTypes.ProgramConfig) types.Program {
	var result = types.Program{
		HTTP_Code: httpCode,
		Status:    statusText,
		ClientIP:  clientIP,
		Data: struct {
			Date     string                                "json:\"date\""
			Programs map[string]programTypes.ProgramConfig "json:\"programs\""
		}{
			Date:     fmt.Sprint(time.Now().String()),
			Programs: program,
		}}

	return result
}
