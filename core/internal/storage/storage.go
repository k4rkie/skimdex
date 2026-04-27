package storage

import (
	"encoding/gob"
	"fmt"
	"os"

	"github.com/k4rkie/skimdex-core/internal/indexer"
)

func StoreToDisk(indexer *indexer.Indexer) (err error) {
	documentStoreFile, err := os.Create("documentStore.gob")
	if err != nil {
		return fmt.Errorf("Couldn't create docunemt store file: %v", err)
	}
	invertedIndexFile, err := os.Create("invertedIndex.gob")
	if err != nil {
		return fmt.Errorf("Couldn't create inverted index file: %v", err)
	}
	defer documentStoreFile.Close()
	defer invertedIndexFile.Close()

	documentStoreEncoder := gob.NewEncoder(documentStoreFile)
	invertedIndexEncoder := gob.NewEncoder(invertedIndexFile)

	documentStoreEncoder.Encode(indexer.DocumentStore)
	invertedIndexEncoder.Encode(indexer.InvertedIndex)

	fmt.Println("Data has been successfully written to the file✅")
	return nil
}

func LoadFromDisk() (documentStore map[uint64]indexer.DocMetaData, invertedIndex map[string][]uint64, err error) {
	documentStoreFile, err := os.Open("documentStore.gob")
	if err != nil {
		return documentStore, invertedIndex, fmt.Errorf("Couldn't open docunemt store file: %v", err)
	}

	invertedIndexFile, err := os.Open("invertedIndex.gob")
	if err != nil {
		return documentStore, invertedIndex, fmt.Errorf("Couldn't open inverted index file: %v", err)
	}
	defer documentStoreFile.Close()
	defer invertedIndexFile.Close()

	documentStoreDecoder := gob.NewDecoder(documentStoreFile)
	invertedIndexDecoder := gob.NewDecoder(invertedIndexFile)

	documentStoreDecoder.Decode(&documentStore)
	invertedIndexDecoder.Decode(&invertedIndex)

	return documentStore, invertedIndex, nil
}
