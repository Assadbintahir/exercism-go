package strand

// ToRNA function takes dna strand string and returns RNA complement string
func ToRNA(dna string) string {
	result := ""
	DNA := []byte(dna)

	for _, value := range DNA {
		result += RNACompToDNANucleotide(value)
	}

	return result
}

// RNACompToDNANucleotide takes a DNA nucleotide and returns RNA complement nucleotide
func RNACompToDNANucleotide(nucleotide byte) string {
	complement := ""

	switch nucleotide {
	case 'G':
		complement = "C"
		break
	case 'C':
		complement = "G"
		break
	case 'T':
		complement = "A"
		break
	case 'A':
		complement = "U"
		break
	}

	return complement
}
