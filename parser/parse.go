// run command: go run parse.go --file=<text file name>.txt

package main

import (
	"flag"
	"os"
	"log"
	"io"
	"encoding/json"
	"fmt"
	"bufio"
	"strings"
)

var (
	fileName = flag.String("file", "", "Enter full path of the file to parse")
)

func main() {
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	wCount, lCount, fq := Parse(f)

	log.Printf("Word Count = %d\n", wCount)
	log.Printf("Line Count = %d\n", lCount)
	log.Printf("letter Frequencies = %s\n", prettify(fq))
}

func Parse(r io.ReadCloser) (w int, l int, f map[string]int) {
	defer r.Close()
	lines := ParseLines(r)
	l = len(lines)

	words := ParseWords(lines)
	w = len(words)

	f = LetterFrequency(words)
	return
}

func ParseLines(r io.Reader) []string {
	var lines []string

	reader := bufio.NewReader(r)
	var line string
	var err error

	for err == nil {
		line, err = reader.ReadString('\n')
		if err == nil {
			lines = append(lines, line)
		}
	}

	if err == io.EOF && len(line) > 0 {
		lines = append(lines, line)
	}

	return lines
}

func ParseWords(lines []string) []string {
	var Words []string

	for _, l := range lines {
		words := strings.Fields(l)
		Words = append(Words, words...)
	}

	return Words
}

func LetterFrequency(words []string) map[string]int {
	store := map[string]int{}

	for _, w := range words {
		for _, l := range w {
			store[string(l)] += 1
		}
	}

	return store
}

func prettify(x interface{}) string {
	b, err := json.MarshalIndent(x, "", " ")
	if err != nil {
		return fmt.Sprintf("Error while prettifying data: %+v", err)
	}

	return string(b)
}
