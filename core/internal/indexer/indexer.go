package indexer

import (
	"hash/fnv"
	"slices"

	"github.com/k4rkie/skimdex-core/internal/models"
)

type DocMetaData struct {
	id    uint64
	path  string
	title string
}

type Indexer struct {
	invertedIndex map[string][]uint64
	documentStore map[uint64]DocMetaData
}

// constuctor func for the Indexer
func NewIndexer() *Indexer {
	return &Indexer{
		invertedIndex: make(map[string][]uint64),
		documentStore: make(map[uint64]DocMetaData),
	}
}

func (i *Indexer) IndexDocument(parsedDoc models.ParsedDocument) {
	docId := generateID(parsedDoc.Path)

	i.documentStore[docId] = DocMetaData{
		id:    docId,
		path:  parsedDoc.Path,
		title: parsedDoc.Title,
	}

	for _, word := range parsedDoc.Keywords {
		// if document id already exits then donot append again
		if slices.Contains(i.invertedIndex[word], docId) {
			continue
		}
		i.invertedIndex[word] = append(i.invertedIndex[word], docId)
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
