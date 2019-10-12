package main

import (
	"fmt"
	"github.com/r0nni3/backend-test/exporter"
	"github.com/r0nni3/backend-test/utils"
	"log"
	"os"
)

func wasError(err error) {
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}

func main() {
	target, err := utils.ParseCLIArgs()
	if err != nil {
		// print CLI Usage information
		fmt.Printf("\t%s %s\n", "import", "<path-to-file-or-dir>")
	}
	wasError(err)

	// Executes tasks
	err = exporter.Run(target)
	wasError(err)

	os.Exit(0)
}
