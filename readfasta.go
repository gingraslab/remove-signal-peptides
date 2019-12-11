package main

import (
	"bufio"
	"log"
	"regexp"
	"strings"
)

type database map[string]*databaseEntry

type databaseEntry struct {
	Header   string
	Sequence string
}

type fasta struct {
	Header        string
	ID            string
	StringBuilder strings.Builder
}

func readFasta(filename string) *database {
	file, err := FS.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reHeader := regexp.MustCompile(`>\w+\|([^|]+)\|`)

	entries := &database{}
	entry := initializeEntry()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			addEntry(entries, entry)
			entry = initializeEntry()
			entry.Header = line
			entry.ID = parseIDFromHeader(reHeader, line)
		} else {
			entry.StringBuilder.WriteString(line)
		}
	}
	addEntry(entries, entry)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return entries
}

func initializeEntry() fasta {
	return fasta{}
}

func addEntry(entries *database, entry fasta) {
	if entry.ID != "" {
		(*entries)[entry.ID] = &databaseEntry{
			Header:   entry.Header,
			Sequence: entry.StringBuilder.String(),
		}
	}
}

func parseIDFromHeader(re *regexp.Regexp, header string) string {
	matches := re.FindStringSubmatch(header)
	if len(matches) > 0 {
		return matches[1]
	}
	return ""
}
