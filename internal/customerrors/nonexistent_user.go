package customerrors

type NonexistentUser struct {
	Err error
}

func (err *NonexistentUser) Unwrap() error {
	return err.Err
}

func (err *NonexistentUser) Error() string {
	return err.Err.Error()
}
