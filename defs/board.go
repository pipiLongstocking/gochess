package defs

// SUndo contains the data needed to 'undo' a chess move and go back in history
type SUndo struct {
	// move denotes the move made
	move int
	// positionKey is a unique key generated for each position
	positionKey uint64
	// enPassant stores the square where an enPassant move can be done (if applicable)
	enPassant int
	// fiftyMoveCounter is one where if for fiftyMoves(one move by each side), if no piece is captured - the game is drawn
	fiftyMoveCounter int
	// castlePermissions stores whether the white and black kings can castle on either sides.
	// White(king, queen)side -> bit 1, bit 2
	// Black(king, queen)side -> bit 3, bit 4
	castlePermissions int
}

// SBoard denotes the board structure to be used by the engine
type SBoard struct {
	// pieces denotes the 120 squares (64 squares + 66 illegal squares should a piece make an out-of-bounds move)
	pieces [BoardSquareNum]int

	// pawns - each element contains the placements of the pawns on the board.
	// Each rank, starting from 1 -> 8, constitutes 8 bits, or one byte.
	// 64 bytes are what is needed to denote the pawn positions for a given colour piece
	pawns [3]uint64

	// pieceNum denotes the number of active pieces on the board, indexed by piece type constants.
	pieceNum [13]int

	// The below three arrays are needed to simplify evaluation cases where pawns are absent.
	// bigPieceNum stores the number of active pieces on the board, that are not pawns - indexed by colour type constants.
	bigPieceNum [3]int
	// majorPieceNum stores the number of active pieces on the board, that are queens and rooks - indexed by colour type constants.
	majorPieceNum [3]int
	// minorPieceNum stores the number of active pieces on the board, that are knights and bishops - indexed by colour type constants.
	minorPieceNum [3]int

	// kingSquare stores the square number in which the king is present - indexed by colour type constants.
	kingSquare [2]int

	// side denotes the current side that has to make the next move.
	side int
	// enPassant stores the square where an enPassant move can be done (if applicable)
	enPassant int
	// fiftyMoveCounter is one where if for fiftyMoves(one move by each side), if no piece is captured - the game is drawn
	fiftyMoveCounter int
	// castlePermissions stores whether the white and black kings can castle on either sides.
	// White(king, queen)side -> bit 1, bit 2
	// Black(king, queen)side -> bit 3, bit 4
	castlePermissions int

	// play denotes how many half-moves are into the current search.
	play int
	// hisPlay denotes how many half-moves have gone on in the current game so far.
	hisPlay int
	// history contains the history of all the moves that were played.
	history [MaxGameMoves]SUndo

	// positionKey is a unique key generated for each position
	positionKey uint64
}

var (
	// Square120to64 is an auxiliary array used to interchange indices from the 120-square board with outlier squares to a standard 64-square board (0-indexed)
	Square120to64 [BoardSquareNum]int
	// Square64to120 is an auxiliary array used to interchange indices from a 0-indexed 64-square board to a 120-square board with outlier squares
	Square64to120 [ChessBoardSquares]int
)
