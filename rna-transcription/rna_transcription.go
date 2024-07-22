package strand

func ToRNA(dna string) (rna string) {
	for _, nuc := range dna {
		switch nuc {
		case 'G':
			rna += "C"
		case 'C':
			rna += "G"
		case 'T':
			rna += "A"
		case 'A':
			rna += "U"
		}
	}

	return rna
}
