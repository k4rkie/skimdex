package parser

import (
	"bufio"
	"regexp"
	"strings"
	"unicode"

	"github.com/k4rkie/skimdex-core/internal/models"
)

func ParseMarkdownDocument(scanner *bufio.Scanner) (parsedDoc models.ParsedDocument) {
	var docTitle string

	// create a set for words in the document using map
	uniqueWordSet := make(map[string]struct{})

	//creating a regex engine for parsing h1 headings
	re := regexp.MustCompile(`^(#)\s+(.*)`)
	for scanner.Scan() {
		currentLine := scanner.Text()

		// if there exists a h1 heading then set it as document title
		if re.MatchString(currentLine) {
			docTitle = currentLine[2:]
		}

		currentLineSanitized := strings.ToLower(sanitizeLine(currentLine))
		words := strings.Split(currentLineSanitized, " ")

		for _, word := range words {
			if word == "" {
				continue
			}
			uniqueWordSet[word] = struct{}{}
		}
	}

	// creating keywords slice out of the set
	keywords := make([]string, 0, len(uniqueWordSet))

	for key, _ := range uniqueWordSet {
		keywords = append(keywords, key)
	}

	return models.ParsedDocument{
		Title:    docTitle,
		Keywords: keywords,
	}
}

// takes a string(line) as input and cleans out the punctuations and special characters
func sanitizeLine(input string) string {
	return strings.Map(func(r rune) rune {
		// unicode.IsPunct checks for things like . , ! ? ; : -,
		// unicode.IsSymbol checks for things like + = $ | ~
		if unicode.IsPunct(r) || unicode.IsSymbol(r) {
			return -1 // -1 doesnt return that char(rune)
		}
		return r
	}, input)
}
