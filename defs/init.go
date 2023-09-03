package defs

import "fmt"

// FileRankTo120SquareNumber takes as input the file, and the rank enum
// and returns the square number on a 120-square board
func FileRankTo120SquareNumber(f, r int) int {
	return 21 + f + (10 * r)
}

// InitSq120to64 initialises the Square120to64 and Square64to120 arrays
func InitSq120to64() {
	sq64index := 0
	impValue := ChessBoardSquares + 1
	for i := 0; i < BoardSquareNum; i += 1 {
		Square120to64[i] = impValue
	}

	for rank := Rank1; rank <= Rank8; rank += 1 {
		for file := FileA; file <= FileH; file += 1 {
			sq120index := FileRankTo120SquareNumber(file, rank)
			Square64to120[sq64index] = sq120index
			Square120to64[sq120index] = sq64index
			sq64index += 1
		}
	}
}

func ChessMetaInit() {
	InitSq120to64()
}

func PrintSq120Board() {
	for i := 0; i < BoardSquareNum; i += 1 {
		fmt.Printf("%5d", Square120to64[i])
		if (i+1)%10 == 0 {
			fmt.Printf("\n")
		}
	}
}

func PrintSq64Board() {
	for i := 0; i < ChessBoardSquares; i += 1 {
		fmt.Printf("%5d", Square64to120[i])
		if (i+1)%8 == 0 {
			fmt.Printf("\n")
		}
	}
}
