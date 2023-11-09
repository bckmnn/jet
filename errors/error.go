package errors

type Error interface {
	error

	Reason() Reason
	WithReason(Reason) Error
	CompleteReason(Reason) Error

	Message() Message
	WithMessage(Message) Error

	Position() Position
	WithPosition(Line, Column) Error
	WithLine(Line) Error
	WithColumn(Column) Error

	Details() Details
	WithDetail(string, string) Error
	WithDetails(Details) Error
}
