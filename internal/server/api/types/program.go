package types

import "github.com/gar-id/queued/internal/server/config/types"

// General struct
type Program struct {
	HTTP_Code int    `json:"httpCode"`
	Status    string `json:"status"`
	ClientIP  string `json:"clientIP"`
	Data      struct {
		Date     string                         `json:"date"`
		Programs map[string]types.ProgramConfig `json:"programs"`
	} `json:"data"`
}
