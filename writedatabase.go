package main

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/spf13/afero"
)

func writeDatabase(updatedDatabase *database, outFile string) {
	var buffer bytes.Buffer

	outputOrder := getOutputOrder(updatedDatabase)

	for _, id := range outputOrder {
		entry := (*updatedDatabase)[id]
		buffer.WriteString(fmt.Sprintf("%s\n", entry.Header))
		writePeptides(&buffer, entry.Sequence)
	}

	afero.WriteFile(FS, outFile, buffer.Bytes(), 0644)
}

func getOutputOrder(updatedDatabase *database) []string {
	keys := make([]string, len(*updatedDatabase))

	i := 0
	for id := range *updatedDatabase {
		keys[i] = id
		i++
	}

	sort.Strings(keys)
	return keys
}

func writePeptides(buffer *bytes.Buffer, sequence string) {
	sequenceLength := len(sequence)
	start := 0
	end := getMinimumInt(sequenceLength, 60)
	for start < sequenceLength {
		buffer.WriteString(fmt.Sprintf("%s\n", sequence[start:end]))
		start = end
		end = getMinimumInt(sequenceLength, end+60)
	}
}

func getMinimumInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
