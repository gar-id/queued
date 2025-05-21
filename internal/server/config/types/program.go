package types

import (
	"time"
)

type ProgramConfig struct {
	Group       string          `json:"group" yaml:"group"`
	Process     []ProcessConfig `json:"processConfig" yaml:"processConfig"`
	Command     string          `json:"command" yaml:"command"`
	AutoStart   bool            `json:"autoStart" yaml:"autoStart"`
	AutoRestart bool            `json:"autoRestart" yaml:"autoRestart"`
	StartSecs   int             `json:"startSecs" yaml:"startSecs"`
	SlowStart   int             `json:"slowStart" yaml:"slowStart"`
	NumProcs    int             `json:"numProcs" yaml:"numProcs"`
	User        string          `json:"user" yaml:"user"`
	Stdout      string          `json:"stdout" yaml:"stdout"`
	Stderr      string          `json:"stderr" yaml:"stderr"`
	Env         []string        `json:"env" yaml:"env"`
}

type ProcessConfig struct {
	LastStart    time.Time `json:"lastStart" yaml:"lastStart"`
	ProcessName  string    `json:"processName" yaml:"processName"`
	ProgramName  string    `json:"programName" yaml:"programName"`
	ProcessIndex int       `json:"processIndex" yaml:"processIndex"`
	PID          *int      `json:"pid" yaml:"pid"`
	Status       string    `json:"status" yaml:"status"`
}

type ProcessChannel struct {
	StopChannel *chan bool
	// Exit        bool
	Name string
}
