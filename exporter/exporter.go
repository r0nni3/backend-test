package exporter

import (
	"errors"
	"github.com/r0nni3/backend-test/utils"
	"io"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// Exporter interface definition for different Exporters
type Exporter interface {
	io.Reader
	// IsValid Check type: local dir, file, url, ftp, ssh, et...
	IsValid(string) (bool, error)
	// Actual file processing
	Process(string) error
}

// Video represents feed exports data, this assumes is standarized
// if its not a type by provider must be done
type Video struct {
	Tags  []string `json:"tags,ommitempty" yaml:"labels,flow"`
	URL   string   `json:"url,ommitempty" yaml:"url"`
	Title string   `json:"title,ommitempty" yaml:"name"`
}

// Videos represents feed exports data, this assumes is standarized
// if its not a type by provider must be done
type Videos struct {
	Videos []Video
}

// Run executes actual exports
func Run(target string) error {
	var wg sync.WaitGroup
	targets, local, _, err := utils.CheckTarget(target)
	if err != nil {
		return err
	}

	log.Println("Running exports...")
	wg.Add(len(targets))
	for _, t := range targets {
		exporter, err := getExporter(t, local)
		if err != nil {
			log.Println(err)
			continue
		}

		// Process files in parallel
		go func(t string) {
			defer wg.Done()
			_, _ = exporter.Read([]byte{})
			_ = exporter.Process(t)
		}(t)
	}
	wg.Wait()

	return nil
}

func processVideo(i int, v Video, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("\timporting[START]: \"%s\"; Url: %s; Tags: %s\n", v.Title, v.URL, strings.Join(v.Tags, ", "))

	// TODO: Do actual job, download, save on S3, etc...
	r := rand.Intn(5)
	time.Sleep(time.Duration(r) * time.Second)
	// ENDTODO

	log.Printf("\timporting[FINISH]: \"%s\"\n", v.Title)
}

func getExporter(target string, local bool) (Exporter, error) {
	if local {
		return getLocalExporter(target)
	}

	return getRemoteExporter(target)
}

func getLocalExporter(target string) (exporter Exporter, err error) {
	err = errors.New("File type not valid")
	fileType, err := utils.GetFileType(target)
	if err != nil {
		exporter = nil
		return
	}

	switch fileType {
	case "yaml":
		exporter = NewYAMLExporter(target)
		err = nil
	case "yml":
		exporter = NewYAMLExporter(target)
		err = nil
	case "json":
		exporter = NewJSONExporter(target)
		err = nil
	default:
		exporter = nil
		err = errors.New("File type not valid")
	}

	ok, err := exporter.IsValid(fileType)
	if ok {
		return
	}

	exporter = nil
	return

}

func getRemoteExporter(target string) (Exporter, error) {
	return nil, errors.New("Not implemented YET")
}
