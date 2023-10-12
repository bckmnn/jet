package jet

import (
	"fmt"
)

type Error struct {
	Template string                 `json:"template,omitempty"`
	Reason   string                 `json:"reason,omitempty"`
	Message  string                 `json:"message,omitempty"`
	Position *Position              `json:"position,omitempty"`
	Details  map[string]interface{} `json:"details,omitempty"`
}

type Position struct {
	Line   int `json:"line,omitempty"`
	Column int `json:"column,omitempty"`
}

func errBuilder() *Error {
	return &Error{}
}

func (e *Error) Error() string {
	place := ""

	if e.Template != "" {
		place = fmt.Sprintf(" %s", e.Template)

		if e.Position != nil {
			place = fmt.Sprintf("%s:%d:%d", place, e.Position.Line, e.Position.Column)
		}
	}

	return fmt.Sprintf(
		"%s%s %s",
		e.Reason, place, e.Message,
	)
}

func (e *Error) WithReason(reason string) *Error {
	e.Reason = reason
	return e
}

func (e *Error) CompleteReason(reason string) *Error {
	e.Reason = fmt.Sprintf("%s.%s", e.Reason, reason)
	return e
}

func (e *Error) WithMessage(message string) *Error {
	e.Message = message
	return e
}

func (e *Error) WithPosition(line, col int) *Error {
	e.Position = &Position{Line: line, Column: col}
	return e
}

func (e *Error) WithDetail(key, value string) *Error {
	if e.Details == nil {
		e.Details = make(map[string]interface{})
	}
	e.Details[key] = value
	return e
}

func (e *Error) WithDetails(details map[string]interface{}) *Error {
	e.Details = details
	return e
}

func NewError(reason, template, message string, position *Position) *Error {
	return &Error{
		Reason:   reason,
		Template: template,
		Message:  message,
		Position: position,
	}
}

const (
	TemplateErrorReason = "jet.template.error"
	RuntimeErrorReason  = "jet.runtime.error"

	InvalidValueReason       = "invalid.value"
	InvalidNumberOfArguments = "invalid.number_of_arguments"

	Unexpected               = "unexpected"
	UnexpectedKeywordReason  = "unexpected.keyword"
	UnexpectedTokenReason    = "unexpected.token"
	UnexpectedNodeReason     = "unexpected.node"
	UnexpectedNodeTypeReason = "unexpected.node.type"
	UnexpectedExpressionType = "unexpected.expression.type"
	UnexpectedCommand        = "unexpected.command"
	UnexpectedClause         = "unexpected.clause"

	NotFoundFieldOrMethod = "not_found.field_or_method"
)
