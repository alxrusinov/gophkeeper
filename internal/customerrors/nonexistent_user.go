package customerrors

// NonexistentUser - error if user does not exist
type NonexistentUser struct {
	Err error
}

// Unwrap - implements unwrap method of error
func (err *NonexistentUser) Unwrap() error {
	return err.Err
}

// Error - implements error method of error
func (err *NonexistentUser) Error() string {
	return err.Err.Error()
}
