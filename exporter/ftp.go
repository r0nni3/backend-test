package exporter

import "errors"

// FTP exporter this identifies a local dir or file to process
type FTP struct{}

// Read complies to reader interface
func (ex *FTP) Read(b []byte) (n int, err error) {
	return -1, errors.New("Not implemented")
}

// IsValid checks if target is local dir or file
func (ex *FTP) IsValid(target string) (bool, error) {
	return false, errors.New("Not implemented")
}

// Process processes input target
func (ex *FTP) Process(target string) error {
	return errors.New("Not implemented")
}
