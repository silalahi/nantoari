package files

import (
	"github.com/google/uuid"
)

type usecase struct {
	r Repository
}

func NewUsecase(r Repository) Usecase {
	return &usecase{
		r: r,
	}
}

func (u *usecase) Get(uuid uuid.UUID) (File, error) {
	return File{}, nil
}

func (u *usecase) Set(file File) error {
	return nil
}
