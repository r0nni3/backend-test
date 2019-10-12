package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestParseCLIArgs(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// missing target argument
	os.Args = []string{"import"}
	target, err := ParseCLIArgs()
	if target != "" || err == nil {
		t.Error("No target argument passed")
	}

	// provided target argument
	os.Args = []string{"import", "some-file"}
	target, err = ParseCLIArgs()
	if target == "" && err != nil {
		t.Error("Target argument passed")
	}
}

func TestCheckTarget(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Log(err)
		return
	}
	defer os.RemoveAll(tempDir)
	tempSubDir, _ := ioutil.TempDir(tempDir, "")
	tmpYAML, _ := ioutil.TempFile(tempSubDir, "*.yaml")
	_, _ = ioutil.TempFile(tempSubDir, "*.json")

	// Not a local file
	_, local, _, err := CheckTarget("http://some-domain/file.json")
	if local || err != nil {
		t.Error("Target is remote")

	}

	// Local file
	// - does not exists
	_, local, _, err = CheckTarget(filepath.Join(tempSubDir, "nonexistentthing"))
	if local || err != nil {
		t.Error("Non existent resources are treated as remote")

	}

	// - does exists is single file
	_, local, isDir, err := CheckTarget(tmpYAML.Name())
	if !local || isDir || err != nil {
		t.Error("Existent single file")

	}

	// - does exists is dir with files
	_, local, isDir, err = CheckTarget(tempSubDir)
	if !local || !isDir || err != nil {
		t.Error("Existent dir with files to process")

	}
}

func TestGetFileType(t *testing.T) {
	empty := ""
	capitalized := "./test.JSON"
	normal := "/path/to/file.yaml"

	target, err := GetFileType(empty)
	if target != "" || err == nil {
		t.Errorf("Empty file provided")
	}

	target, err = GetFileType(capitalized)
	if err != nil || target != "json" {
		t.Error()
	}

	target, err = GetFileType(normal)
	if err != nil || target != "yaml" {
		t.Error()
	}
}
