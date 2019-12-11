// Package main removes signal peptides from UniProt sequences.
package main

func main() {
	fastaFile := "./files/human_29-9-2019.fasta"
	peptideSequences := "./files/human-signal-peptides.txt"
	outFile := "./files/mature-sequences.txt"

	uniprotdatabase := readFasta(fastaFile)
	signalPeptides := readSignalPeptides(peptideSequences)

	trimPeptides(signalPeptides, uniprotdatabase)
	writeDatabase(uniprotdatabase, outFile)
}
