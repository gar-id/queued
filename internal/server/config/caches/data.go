package caches

import (
	"sync"

	"github.com/gar-id/queued/internal/server/config/types"
)

// Create types for caches
type ProgramTypes struct {
	Mutex         sync.RWMutex
	ProgramConfig map[string]types.ProgramConfig
	// ProcessChannel map[string]types.ProcessChannel
	GroupIndex map[string][]string
	WaitGroup  sync.WaitGroup
}
type ProcessTypes struct {
	Mutex sync.RWMutex
	Data  map[string]types.ProcessChannel
}

// Queue program data is stored in memory with this variable
var Data = ProgramTypes{
	ProgramConfig: make(map[string]types.ProgramConfig),
	GroupIndex:    make(map[string][]string),
}
var ProcessChannel = ProcessTypes{
	Data: make(map[string]types.ProcessChannel),
}

// Global func to execute func with mutex
func (c *ProgramTypes) Do(run func()) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	run()
}
func (c *ProgramTypes) ROLock(run func()) {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()
	run()
}
func (c *ProcessTypes) Do(run func()) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	run()
}
func (c *ProcessTypes) ExportProcessChannel(run func() types.ProcessChannel) types.ProcessChannel {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	export := run()
	return export
}
