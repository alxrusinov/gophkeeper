package usecase

// Repository - interface of repository
type Repository interface{}

type Usecase struct {
	repository Repository
}

// NewUsecase - create new instance of Usecase
func NewUsecase(repository Repository) *Usecase {
	uc := &Usecase{repository: repository}

	return uc
}
