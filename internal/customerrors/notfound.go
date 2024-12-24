package customerrors

// NotFound - error if source does not find
type NotFound struct {
	Err error
}

// Unwrap - implements unwrap method of error
func (err *NotFound) Unwrap() error {
	return err.Err
}

// Error - implements error method of error
func (err *NotFound) Error() string {
	return err.Err.Error()
}
