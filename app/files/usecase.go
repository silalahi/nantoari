package files

import (
	"github.com/google/uuid"
)

// usecase is implementation of Usecase
type usecase struct {
	r Repository
}

// NewUsecase returns an instance of usecase using
// using given r as repository
func NewUsecase(r Repository) Usecase {
	return &usecase{
		r: r,
	}
}

// Get retrieves File from repository by given UUID.
func (u *usecase) Get(uuid uuid.UUID) (File, error) {
	return u.r.Get(uuid)
}

// Set stores File to repository.
func (u *usecase) Set(file File) error {
	return u.r.Set(file)
}
