package jet

import "fmt"

type Error struct {
	Reason   string                 `json:"reason,omitempty"`
	Message  string                 `json:"message,omitempty"`
	Position Position               `json:"position,omitempty"`
	Details  map[string]interface{} `json:"details,omitempty"`
}

type Position struct {
	Line   int `json:"line,omitempty"`
	Column int `json:"column,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s (%d:%d): %s [%v]", e.Reason, e.Position.Line, e.Position.Column, e.Message, e.Details)
}

func NewError(reason, message string, position Position, details map[string]interface{}) *Error {
	return &Error{
		Reason:   reason,
		Message:  message,
		Position: position,
		Details:  details,
	}
}

const (
	TemplateError = "Jet Template Error"
	RuntimeError  = "Jet Runtime Error"
)
