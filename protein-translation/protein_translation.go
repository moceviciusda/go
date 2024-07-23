package protein

import (
	"errors"
)

var (
	ErrStop        = errors.New("stop")
	ErrInvalidBase = errors.New("invalid codon")
)

func FromRNA(rna string) ([]string, error) {
	var proteins []string
	for i := 0; i < len(rna)-2; i += 3 {
		protein, err := FromCodon(rna[i : i+3])
		if err == ErrInvalidBase {
			return proteins, err
		}
		if err == ErrStop {
			break
		}

		proteins = append(proteins, protein)
	}

	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	var protein string
	switch codon {
	case "AUG":
		protein = "Methionine"
	case "UUU", "UUC":
		protein = "Phenylalanine"
	case "UUA", "UUG":
		protein = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		protein = "Serine"
	case "UAU", "UAC":
		protein = "Tyrosine"
	case "UGU", "UGC":
		protein = "Cysteine"
	case "UGG":
		protein = "Tryptophan"
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}

	return protein, nil
}
