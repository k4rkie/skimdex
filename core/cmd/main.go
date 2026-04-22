package main

import (
	"log"
	"os"

	"github.com/k4rkie/skimdex-core/internal/crawler"
	"github.com/k4rkie/skimdex-core/internal/index"
)

func main() {

	// NOTE: Just a place holder path
	rootDir := "path/to/docs"

	filePaths, err := crawler.FindFiles(rootDir)
	if err != nil {
		log.Printf("Error finding files in the given directory: %v", err)
	}

	// Create new instance of the indexer
	indexer := index.NewIndexer()

	for _, filePath := range filePaths {
		file, _ := os.Open(filePath)
		defer file.Close()
		indexer.RegisterDocStruct(filePath, file)
	}

}
