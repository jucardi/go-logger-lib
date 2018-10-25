package log

import (
	"time"
)

// Entry represents a log entry.
type Entry struct {
	// Contains all the fields set by the user. TODO
	Data map[string]interface{}

	// Time at which the log entry was created
	Timestamp time.Time

	// Level the log entry was logged at: Debug, Info, Warn, Error, Fatal or Panic
	Level Level

	// Message passed to Debug, Info, Warn, Error, Fatal or Panic
	Message string

	metadata map[string]interface{}
}

func (e *Entry) AddMetadata(key string, val interface{}) {
	if e.metadata == nil {
		e.metadata = map[string]interface{}{}
	}
	e.metadata[key] = val
}
