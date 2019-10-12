package exporter

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// JSON exporter this identifies a local dir or file to process
type JSON struct {
	//r io.Reader
	//e Exporter
	f    *os.File
	data *Videos
}

// NewJSONExporter creates new JSON exporte
func NewJSONExporter(target string) *JSON {
	f, _ := os.Open(target)
	return &JSON{
		f:    f,
		data: &Videos{},
	}
}

// Read resource content
func (ex *JSON) Read(bytes []byte) (int, error) {
	defer func(f *os.File) { _ = f.Close() }(ex.f)

	byteValue, err := ioutil.ReadAll(ex.f)
	_ = json.Unmarshal(byteValue, ex.data)

	return len(byteValue), err
}

// IsValid checks if target is local dir or file
func (ex *JSON) IsValid(target string) (bool, error) {
	if target == "json" {
		return true, nil
	}

	return false, errors.New("No valid JSON file")
}

// Process processes input target
func (ex *JSON) Process(target string) error {
	log.Printf("Processing [START]: %s\n", target)

	var wg sync.WaitGroup
	wg.Add(len(ex.data.Videos))
	for i, video := range ex.data.Videos {
		go processVideo(i, video, &wg)
	}
	wg.Wait()

	log.Printf("Processing [FINISH]: %s\n", target)
	return nil
}
