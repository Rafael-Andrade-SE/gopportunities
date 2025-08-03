package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return errors.New("fake error for testing")
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	return NewLogger(p)
}
