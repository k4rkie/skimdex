package indexer

import (
	"github.com/k4rkie/skimdex-core/internal/models"
	"hash/fnv"
)

type DocMetaData struct {
	Id    uint64
	Path  string
	Title string
}

type Indexer struct {
	InvertedIndex map[string][]uint64
	DocumentStore map[uint64]DocMetaData
}

// constuctor func for the Indexer
func NewIndexer() *Indexer {
	return &Indexer{
		InvertedIndex: make(map[string][]uint64),
		DocumentStore: make(map[uint64]DocMetaData),
	}
}

func (i *Indexer) IndexDocument(parsedDoc models.ParsedDocument) {
	docId := generateID(parsedDoc.Path)

	i.DocumentStore[docId] = DocMetaData{
		Id:    docId,
		Path:  parsedDoc.Path,
		Title: parsedDoc.Title,
	}

	for _, word := range parsedDoc.Keywords {
		i.InvertedIndex[word] = append(i.InvertedIndex[word], docId)
	}
}

// generate a unique 64bit document id
func generateID(input string) uint64 {
	//Create the hasher
	h := fnv.New64a()
	//Feed the data (The Input)
	h.Write([]byte(input))
	//Get the result
	return h.Sum64()
}
