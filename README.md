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
https://www.uniprot.org/uniprot/?query=*&fil=reviewed%3Ayes+AND+organism%3A%22Homo+sapiens+%28Human%29+%5B9606%5D%22&sort=score
```
   
2. List of signal peptides.

The list of signal peptides can be generated at UniProt with the following search:
```
https://www.uniprot.org/uniprot/?query=annotation%3A%28type%3Asignal%29&sort=score
```

and then filtering by other criteria such as species, reviewed, etc. For reviewed human sequences, the url would be
```
https://www.uniprot.org/uniprot/?query=annotation:(type:signal)&fil=reviewed%3Ayes+AND+organism%3A%22Homo+sapiens+%28Human%29+%5B9606%5D%22&sort=score
```

And finally add a column to the results table for signal peptides. Click the `Columns` button above the table and look under PTM/Processing. Then download all results as a tab-separated file.

## Run

1. Specify input files in `main.go`

2. Build
```
go build
```

3. Run
```
remove-signal-peptides
```

4. Output: `./files/mature-sequences.txt`
