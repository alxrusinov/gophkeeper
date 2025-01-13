package customerrors

// UnverifiedUser - error if user does not verify
type UnverifiedUser struct {
	Err error
}

// Unwrap - implements unwrap method of error
func (err *UnverifiedUser) Unwrap() error {
	return err.Err
}

// Error - implements error method of error
func (err *UnverifiedUser) Error() string {
	return err.Err.Error()
}
