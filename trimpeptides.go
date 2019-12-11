package main

func trimPeptides(signalPeptides map[string]int, uniprotDatabase *database) {
	for id, end := range signalPeptides {
		if _, ok := (*uniprotDatabase)[id]; ok {
			(*uniprotDatabase)[id].Sequence = trimPeptide((*uniprotDatabase)[id].Sequence, end)
		}
	}
}

func trimPeptide(sequence string, end int) string {
	position := 0
	for i := range sequence {
		if position >= end {
			return sequence[i:]
		}
		position++
	}
	return sequence[:0]
}
