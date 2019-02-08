package files

import (
	"encoding/csv"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

var (
	// ErrNotFound is an error when the requested resource is not found
	ErrNotFound = errors.New("the requested resource is not found")
	// ErrInvalidUUID is an error when the given UUID is not valid / cannot parsed
	ErrInvalidUUID = errors.New("invalid UUID")
	// ErrInvalidFileFormat is an error when the given file is not match with expected
	ErrInvalidFileFormat = errors.New("invalid file format")
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

// IsCSV performs file fetching from given URL and parse into CSV.
// If an error raised, it will returns false. Otherwise, it returns true.
func (f File) IsCSV() bool {
	res, err := http.Get(f.URL())
	if err != nil {
		return false
	}
	defer res.Body.Close()

	reader := csv.NewReader(res.Body)
	_, err = reader.ReadAll()
	if err != nil {
		return false
	}

	return true
}

// Read reads CSV file from given URL and read all of its rows
func (f File) Read() ([][]string, error) {
	var records [][]string
	res, err := http.Get(f.URL())
	if err != nil {
		return records, err
	}
	defer res.Body.Close()

	reader := csv.NewReader(res.Body)
	records, err = reader.ReadAll()
	if err != nil {
		return records, err
	}

	return records, nil
}

// Parse parses records read from Read()
func (f File) Parse() (interface{}, error) {
	records, err := f.Read()
	if err != nil {
		return nil, err
	}

	headers := records[0]
	results := make([]interface{}, 0)

	for _, row := range records[1:] {
		result := make(map[string]interface{}, 0)
		for i, column := range row {
			result[headers[i]] = column
		}

		results = append(results, result)
	}

	return results, nil
}
