package types

// General struct
type General struct {
	HTTP_Code int    `json:"httpCode"`
	Status    string `json:"status"`
	ClientIP  string `json:"clientIP"`
	Data      struct {
		Date    string `json:"date"`
		Message string `json:"message"`
	} `json:"data"`
}
type GeneralObject struct {
	HTTP_Code int    `json:"httpCode"`
	Status    string `json:"status"`
	ClientIP  string `json:"clientIP"`
	Data      struct {
		Date    string              `json:"date"`
		Message map[string][]string `json:"message"`
	} `json:"data"`
}

type Welcome struct {
	HTTP_Code int    `json:"httpCode"`
	Status    string `json:"status"`
	ClientIP  string `json:"clientIP"`
	Data      struct {
		Date    string `json:"date"`
		Message string `json:"message"`
		Version string `json:"version"`
	} `json:"data"`
}

type ErrorHandler struct {
	HTTP_Code int    `json:"httpCode"`
	Status    string `json:"status"`
	ClientIP  string `json:"clientIP"`
}
