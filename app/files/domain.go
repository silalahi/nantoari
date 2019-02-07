package files

import (
	"errors"

	"github.com/google/uuid"
)

var (
	// ErrNotFound is an error when the requested resource is not found
	ErrNotFound = errors.New("the requested resource is not found")
)

// Repository represent the file's repository contract
type Repository interface {
	Get(uuid uuid.UUID) (File, error)
	Set(file File) error
}

// Usecase represent the file's usecases
type Usecase interface {
	Get(uuid uuid.UUID) (File, error)
	Set(file File) error
}

// File represent UUID and URL of file paired
type File struct {
	uuid uuid.UUID
	url  string
}

// UUID returns file's UUID
func (f File) UUID() uuid.UUID {
	return f.uuid
}

// URL return file's location
func (f File) URL() string {
	return f.url
}

func (f File) IsValid() bool {
	// Perform url validation
	return false
}

func (f File) IsCSVFormat() bool {
	// Perform file validation
	return false
}
