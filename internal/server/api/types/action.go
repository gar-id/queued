package types

type ActionPayload struct {
	ProgramName []string `json:"programName" yaml:"programName" form:"programName"`
	ProcessName []string `json:"processName" yaml:"processName" form:"processName"`
	GroupName   []string `json:"groupName" yaml:"groupName" form:"groupName"`
}
