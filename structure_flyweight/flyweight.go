package flyweight

/*
	下棋
*/

var UintM map[string]*ChessPieceUint = map[string]*ChessPieceUint{
	"u1": {
		ID:    "u1",
		Name:  "u1",
		Color: "w",
	},
	"u2": {
		ID:    "u2",
		Name:  "u2",
		Color: "b",
	},
	// 其它 ... ...
}

func NewChessPieceUint(id string) *ChessPieceUint {
	return UintM[id]
}

func NewChessChessBoard() *ChessBoard {
	board := &ChessBoard{chessPieces: map[string]*ChessPiece{}}
	for id := range UintM {
		board.chessPieces[id] = &ChessPiece{
			Unit: NewChessPieceUint(id),
			X:    0,
			Y:    0,
		}
	}
	return board
}

type ChessBoard struct {
	chessPieces map[string]*ChessPiece
}

type ChessPieceUint struct {
	ID    string
	Name  string
	Color string
}

// ChessPiece 棋子
type ChessPiece struct {
	Unit *ChessPieceUint
	X    int
	Y    int
}

// Move 移动
func (cb *ChessBoard) Move(id string, x, y int) {
	cb.chessPieces[id].X = x
	cb.chessPieces[id].Y = y
}
