package main

import (
	"encoding/csv"
	"io"
	"log"
	"regexp"
	"strconv"

	"github.com/spf13/afero"
)

func readSignalPeptides(filename string) map[string]int {
	file, err := FS.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reSignal := regexp.MustCompile(`^SIGNAL 1 (\d+)`)

	reader := initializeReader(file)
	signalPeptides := make(map[string]int, 0)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		end := parseEnd(reSignal, line[7])
		if end != 0 {
			signalPeptides[line[0]] = end
		}
	}

	return signalPeptides
}

func initializeReader(file afero.File) *csv.Reader {
	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	skipHeader(reader)

	return reader
}

func skipHeader(reader *csv.Reader) {
	_, err := reader.Read()
	if err != nil {
		log.Fatalln(err)
	}
}

func parseEnd(re *regexp.Regexp, str string) int {
	matches := re.FindStringSubmatch(str)
	if len(matches) > 0 {
		end, _ := strconv.Atoi(matches[1])
		return end
	}
	return 0
}
