package main

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var signalPeptidesText = `Entry	Entry name	Status	Protein names	Gene names	Organism	Length	Signal peptide
Q96IY4	CBPB2_HUMAN	reviewed	Carboxypeptidase B2 (EC 3.4.17.20) (Carboxypeptidase U) (CPU) (Plasma carboxypeptidase B) (pCPB) (Thrombin-activable fibrinolysis inhibitor) (TAFI)	CPB2	Homo sapiens (Human)	423	SIGNAL 1 22 {ECO:0000255}.
P22362	CCL1_HUMAN	reviewed	C-C motif chemokine 1 (Small-inducible cytokine A1) (T lymphocyte-secreted protein I-309)	CCL1 SCYA1	Homo sapiens (Human)	96	SIGNAL 2 23 {ECO:0000269|PubMed:15340161, ECO:0000269|PubMed:1557400}.
Q04771	ACVR1_HUMAN	reviewed	Activin receptor type-1 (EC 2.7.11.30) (Activin receptor type I) (ACTR-I) (Activin receptor-like kinase 2) (ALK-2) (Serine/threonine-protein kinase receptor R1) (SKR1) (TGF-B superfamily receptor type I) (TSR-I)	ACVR1 ACVRLK2	Homo sapiens (Human)	509	SIGNAL 1 20 {ECO:0000250}.
`

func TestReadSignalPeptides(t *testing.T) {
	oldFs := FS
	defer func() { FS = oldFs }()
	FS = afero.NewMemMapFs()

	// Create test directory and files.
	FS.MkdirAll("test", 0755)
	afero.WriteFile(
		FS,
		"test/signal-peptides.txt",
		[]byte(signalPeptidesText),
		0444,
	)

	expected := map[string]int{
		"Q96IY4": 22,
		"Q04771": 20,
	}
	assert.Equal(t, expected, readSignalPeptides("test/signal-peptides.txt"), "should parse signal peptides")
}
