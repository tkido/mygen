package ui

import (
	"log"
)

// UiEvent is type of all UI event
type UiEvent int

// UiEvents
const (
	GotFocus UiEvent = iota
	LostFocus
)

// String() for fmt.Stringer interface
func (e UiEvent) String() string {
	switch e {
	case GotFocus:
		return "GotFocus"
	case LostFocus:
		return "LostFocus"
	default:
		log.Panicf("unknown UiEvent %d", e)
		return ""
	}
}

// uiCallbacks is callbacks for ui event
type uiCallbacks map[UiEvent]Callback

// handleUiEvent handle ui event
func (b *Box) handleUiEvent(e UiEvent) bool {
	if c, ok := b.uiCallbacks[e]; ok {
		c(b.Self)
		return true
	}
	return false
}

// SetUiCallback set callback function for key. set nil means delete.
func (u uiCallbacks) SetUiCallback(e UiEvent, cb Callback) {
	if cb == nil {
		delete(u, e)
		return
	}
	u[e] = cb
}
