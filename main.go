// Package main removes signal peptides from UniProt sequences.
package main

func main() {
	fastaFile := "./files/human.fasta"
	peptideSequences := "./files/human-signal-peptides.tab"
	outFile := "./files/mature-sequences.txt"

	uniprotdatabase := readFasta(fastaFile)
	signalPeptides := readSignalPeptides(peptideSequences)

	trimPeptides(signalPeptides, uniprotdatabase)
	writeDatabase(uniprotdatabase, outFile)
}
