package index

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"slices"
	"strings"
	"unicode"
)

type Document struct {
	id      uint64
	path    string
	title   string
	snippet string
}

type Indexer struct {
	invertedIndex map[string][]uint64
	documentStore map[uint64]Document
}

// constuctor func for the Indexer
func NewIndexer() *Indexer {
	return &Indexer{
		invertedIndex: make(map[string][]uint64),
		documentStore: make(map[uint64]Document),
	}
}

// Creates the Document struct identifier for looking up the file where the keyword occured
// TODO: add title and snippet to the document struct
func (i *Indexer) RegisterDocStruct(filePath string, file *os.File) {
	docId := generateID(filePath)
	i.documentStore[docId] = Document{
		id:   docId,
		path: filePath,
	}
}

// Creates an inverted index for each work in the document
func (i *Indexer) IndexDocument(file *os.File, documentId uint64) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// get the current line
		currentLine := scanner.Text()

		sanitizedLine := sanitizeLine(currentLine)

		// split the line into words
		keywords := strings.Split(sanitizedLine, " ")
		fmt.Println(keywords)

		for _, word := range keywords {
			if word == "" {
				continue
			}
			// if document id already exits then donot append again
			if slices.Contains(i.invertedIndex[word], documentId) {
				continue
			}
			i.invertedIndex[word] = append(i.invertedIndex[word], documentId)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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

// takes a string(line) as input and cleans out the punctuations and special characters
func sanitizeLine(input string) string {
	return strings.Map(func(r rune) rune {
		// unicode.IsPunct checks for things like . , ! ? ; : -, unicode.IsSymbol checks for things like + = $ | ~
		if unicode.IsPunct(r) || unicode.IsSymbol(r) {
			return -1 // -1 doesnt return that char(rune)
		}
		return r
	}, input)
}
