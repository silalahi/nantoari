package file

import (
	"github.com/google/uuid"
	"github.com/silalahi/nantoari/store"
)

var kv store.KV

// SetStore sets store implementation
func SetStore(repo store.KV)  {
	kv = repo
}

// Get retrieves File from repository by given UUID.
func Get(id uuid.UUID) (File, error) {
	url, err := kv.Get(id.String())
	if err != nil {
		return File{}, ErrNotFound
	}

	return File{
		uuid: id,
		url:  url.(string),
	}, nil
}

// Set stores File to repository.
func Set(file File) error {
	return kv.Set(file.UUID().String(), file.URL())
}
