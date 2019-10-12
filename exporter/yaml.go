package exporter

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

// YAML exporter this identifies a local dir or file to process
type YAML struct {
	//r io.Reader
	//e Exporter
	f    *os.File
	Data *YAMLVideos
}

// YAMLVideos spefic implementation of Videos array for yaml files
type YAMLVideos struct {
	Videos []YAMLVideo
}

// YAMLVideo spefic implementation of Video  for yaml files
type YAMLVideo struct {
	Tags  string `yaml:"labels"`
	URL   string `yaml:"url"`
	Title string `yaml:"name"`
}

// NewYAMLExporter creates new JSON exporte
func NewYAMLExporter(target string) *YAML {
	f, _ := os.Open(target)
	return &YAML{
		f:    f,
		Data: &YAMLVideos{Videos: make([]YAMLVideo, 3)},
	}
}

// Read complies to reader interface
func (ex *YAML) Read(b []byte) (n int, err error) {
	defer func(f *os.File) { _ = f.Close() }(ex.f)

	byteValue, _ := ioutil.ReadAll(ex.f)
	_ = yaml.Unmarshal(byteValue, &ex.Data.Videos)

	return len(byteValue), err
}

// IsValid checks if target is local dir or file
func (ex *YAML) IsValid(target string) (bool, error) {
	if target == "yaml" || target == "yml" {
		return true, nil
	}

	return false, errors.New("No valid YAML file")
}

// Process processes input target
func (ex *YAML) Process(target string) error {
	log.Printf("Processing [START]: %s\n", target)

	var wg sync.WaitGroup
	videos := ex.Data.Videos
	wg.Add(len(videos))
	for i, video := range videos {
		go processVideo(i, Video{
			Tags:  strings.Split(video.Tags, ","),
			URL:   video.URL,
			Title: video.Title,
		}, &wg)
	}
	wg.Wait()

	log.Printf("Processing [FINISH]: %s\n", target)
	return nil
}
