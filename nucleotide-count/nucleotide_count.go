package dna

import (
	"fmt"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts method generates a histogram of valid nucleotides in the given DNA.
func (d DNA) Counts() (Histogram, error) {
	strand := []byte(d)
	var histogram = Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	for _, value := range strand {
		if !isValidNucleotide(value) {
			return nil, fmt.Errorf("%v is not a valid nucleotide", value)
		}

		histogram[rune(value)] = histogram[rune(value)] + 1
	}

	return histogram, nil
}

// isValidNucleotide function checks takes a DNA nucleotide and tells if its valid or not
func isValidNucleotide(nucl byte) bool {
	return nucl == 'A' || nucl == 'C' || nucl == 'G' || nucl == 'T'
}
