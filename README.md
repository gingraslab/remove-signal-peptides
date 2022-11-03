# Remove signal peptides

Remove signal peptides from a UniProt fasta database.

## Installation

1. Ensure [GO](https://golang.org/doc/install) is installed.

2. Clone repo
```
git clone https://github.com/gingraslab/remove-signal-peptides.git
cd remove-signal-peptides
```

3. Build executable
```
go build
```

## Requirements

1. UniProt database in fasta format. These can be downloaded by searching, for example, for reviewed human entries and downloading in fasta format

```
https://www.uniprot.org/uniprotkb?facets=model_organism%3A9606%2Creviewed%3Atrue&query=%2A
```
   
2. List of signal peptides.

The list of signal peptides can be generated at UniProt with the following search:
```
https://www.uniprot.org/uniprotkb?query=(ft_signal:*)
```

and then filtering by other criteria such as species, reviewed, etc. For reviewed human sequences, the url would be
```
https://www.uniprot.org/uniprotkb?facets=model_organism%3A9606%2Creviewed%3Atrue&query=%28ft_signal%3A%2A%29
```

And finally add a column to the results table for signal peptides. Click the `Customize columns` button above the table and look under PTM/Processing. Then download all results as a tab-separated (TSV) file. Ensure the `Signal peptide` column is the seventh column in the download file.

## Run

1. Specify input files in `main.go`

2. Build
```
go build
```

3. Run
```
./remove-signal-peptides
```

4. Output: `./files/mature-sequences.txt`
