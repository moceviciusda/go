package chessboard

import "fmt"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

func StringifyChessboard(cb Chessboard) string {
	var board string

	board += "  A B C D E F G H\n"

	for i := 0; i < 8; i++ {
		board += fmt.Sprint(i + 1)
		for _, file := range cb {
			if file[i] {
				board += " #"
			} else {
				board += " _"
			}
		}

		board += " " + fmt.Sprint(i+1) + "\n"
	}
	board += "   A B C D E F G H"

	return board
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var total int

	for _, occupied := range cb[file] {
		if occupied {
			total++
		}
	}

	return total

}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var total int

	if rank < 1 || rank > 8 {
		return total
	}

	for _, file := range cb {
		if file[rank-1] {
			total++
		}
	}

	return total
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var total int

	for _, file := range cb {
		total += len(file)
	}

	return total
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var total int

	for _, file := range cb {
		for _, sq := range file {
			if sq {
				total++
			}
		}
	}

	return total
}
