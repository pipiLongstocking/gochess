package defs

type u64 uint64

// Basic board/engine metadata
const (
	Name           = "gochess 1.0"
	BoardSquareNum = 120
)

// Enums for pieces on the board
const (
	Empty = iota
	WPawn
	WKnight
	WBishop
	WRook
	WQueen
	WKing
	BPawn
	BKnight
	BBishop
	BRook
	BQueen
	BKing
)

// Enums for files on the board
const (
	FileA = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
	FileNone
)

// Enums for ranks on the board
const (
	RankA = iota
	RankB
	RankC
	RankD
	RankE
	RankF
	RankG
	RankH
	RankNone
)

const (
	White = iota
	Black
	Both
)

// Enums for squares
const (
	A1 = iota + 21
	B1
	C1
	D1
	E1
	F1
	G1
	H1
)
const (
	A2 = iota + 31
	B2
	C2
	D2
	E2
	F2
	G2
	H2
)
const (
	A3 = iota + 41
	B3
	C3
	D3
	E3
	F3
	G3
	H3
)
const (
	A4 = iota + 51
	B4
	C4
	D4
	E4
	F4
	G4
	H4
)
const (
	A5 = iota + 61
	B5
	C5
	D5
	E5
	F5
	G5
	H5
)
const (
	A6 = iota + 71
	B6
	C6
	D6
	E6
	F6
	G6
	H6
)
const (
	A7 = iota + 81
	B7
	C7
	D7
	E7
	F7
	G7
	H7
)
const (
	A8 = iota + 91
	B8
	C8
	D8
	E8
	F8
	G8
	H8
	NoSquare
)
