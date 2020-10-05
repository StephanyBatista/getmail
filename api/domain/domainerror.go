package domain

//Error is custom error of domain
type Error struct {
	message string
}

//NewError create a new Error
func NewError(message string) *Error {
	return &Error{message: message}
}

//Error return the error string
func (d Error) Error() string {
	return d.message
}
