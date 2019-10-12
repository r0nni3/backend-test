package exporter

import "errors"

// URL exporter this identifies a local dir or file to process
type URL struct{}

// Read complies to reader interface
func (ex *URL) Read([]byte) (int, error) {
	return -1, errors.New("Not implemented")
}

// IsValid checks if target is local dir or file
func (ex *URL) IsValid(target string) (bool, error) {
	return false, errors.New("Not implemented")
}

// Process processes input target
func (ex *URL) Process(target string) error {
	return errors.New("Not implemented")
}
