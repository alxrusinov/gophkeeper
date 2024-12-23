package customerrors

type UnverifiedUser struct {
	Err error
}

func (err *UnverifiedUser) Unwrap() error {
	return err.Err
}

func (err *UnverifiedUser) Error() string {
	return err.Err.Error()
}
