package main

import (
	"bufio"
	"log"
	"os"

	"github.com/k4rkie/skimdex-core/internal/crawler"
	"github.com/k4rkie/skimdex-core/internal/indexer"
	"github.com/k4rkie/skimdex-core/internal/parser"
)

func main() {

	// NOTE: Just a place holder path
	rootDir := "path/to/docs"

	filePaths, err := crawler.FindFiles(rootDir)
	if err != nil {
		log.Printf("Error finding files in the given directory: %v", err)
	}

	// Create new instance of the indexer
	indexer := indexer.NewIndexer()

	for _, filePath := range filePaths {
		// open the file to parse and index
		file, _ := os.Open(filePath)

		// create a scanner for the opened file
		scanner := bufio.NewScanner(file)
		parsedDoc := parser.ParseMarkdownDocument(scanner)

		file.Close()
		parsedDoc.Path = filePath
		indexer.IndexDocument(parsedDoc)
	}

}
