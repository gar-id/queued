package types

// Create enums for TokenType
type ProcessStatusType string
type setupProcessStatus string

// Enum values for TokenType
const (
	ProcessStatusStopped  ProcessStatusType = "stopped"
	ProcessStatusStopping ProcessStatusType = "stopping"
	ProcessStatusRunning  ProcessStatusType = "running"
	ProcessStatusStarting ProcessStatusType = "starting"
	ProcessStatusFatal    ProcessStatusType = "fatal"
)

// Values returns all known values for ProcessStatusType. Note that this can
// be expanded in the future.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (setupProcessStatus) Values() []ProcessStatusType {
	return []ProcessStatusType{
		"stopped",
		"stopping",
		"running",
		"starting",
		"fatal",
	}
}
