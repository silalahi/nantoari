package config

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

var (
	// ValidFileExtensions is a list of valid config file format
	ValidFileExtensions = map[string]bool{
		"yaml": true,
		"yml":  true,
	}

	// ErrFileFormatNotSupported error when config file format not supported
	ErrFileFormatNotSupported = errors.New("config file format not supported")
)

// Server is a struct to configure server
type Server struct {
	Port int `yaml:"port"`
}

// Config is a struct to load application configuration
type Config struct {
	Debug  bool    `yaml:"debug"`
	Server *Server `yaml:"server"`
}

// Default returns a default config instance
func Default() *Config {
	return &Config{
		Server: &Server{
			Port: DefaultServerPort,
		},
	}
}

// Load creates a Config struct from a config file path
func Load(path string) (*Config, error) {
	if !IsValidFile(path) {
		return nil, ErrFileFormatNotSupported
	}

	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	var cfg Config

	decoder := yaml.NewDecoder(fp)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}

	return &cfg, nil
}

// IsValidFile returns whether file name f is one of the supported
// config formats.
func IsValidFile(f string) bool {
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(f), "."))
	_, valid := ValidFileExtensions[ext]
	return valid
}
