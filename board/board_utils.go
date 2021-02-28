package board

import (
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// Val matrices for

var kingValBlack = [8][8]float64{
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, -0.5, 0, 0, 0, -0.5, 0},
}

var rookValBlack = [8][8]float64{
	{-5.0, -5.0, -5.0, -5.0, -5.0, -5.0, -5.0, -5.0},
	{-5.0, -5.0, -5.0, -5.0, -5.0, -5.0, -5.0, -5.0},
	{-5.0, -5.0, -5.0, -5.0, -5.0, -5.0, -5.0, -5.0},
	{-5.0, -5.0, -5.0, -5.0, -5.0, -5.0, -5.0, -5.0},
	{-5.0, -5.0, -5.0, -5.0, -5.0, -5.0, -5.0, -5.0},
	{-5.0, -5.0, -5.0, -5.0, -5.0, -5.0, -5.0, -5.0},
	{-5.0, -5.0, -5.0, -5.0, -5.0, -5.0, -5.0, -5.0},
	{-4.95, -4.95, -5.1, -5.1, -5.1, -5.1, -4.95, -4.95},
}

var knightValBlack = [8][8]float64{
	{-2.8, -2.9, -2.9, -2.9, -2.9, -2.9, -2.9, -2.8},
	{-2.9, -3.0, -3.0, -3.0, -3.0, -3.0, -3.0, -2.9},
	{-2.9, -3.0, -3.05, -3.05, -3.05, -3.05, -3.0, -2.9},
	{-2.9, -3.0, -3.2, -3.25, -3.25, -3.2, -3.0, -2.9},
	{-2.9, -3.0, -3.2, -3.25, -3.25, -3.2, -3.0, -2.9},
	{-2.9, -3.0, -3.05, -3.05, -3.05, -3.05, -3.0, -2.9},
	{-2.9, -3.0, -3.0, -3.0, -3.0, -3.0, -3.0, -2.9},
	{-2.8, -2.9, -2.9, -2.9, -2.9, -2.9, -2.9, -2.8},
}
var bishopValBlack = [8][8]float64{
	{-3.2, -3.2, -3.2, -3.2, -3.2, -3.2, -3.2, -3.2},
	{-3.2, -3.3, -3.3, -3.3, -3.3, -3.3, -3.3, -3.2},
	{-3.2, -3.3, -3.3, -3.3, -3.3, -3.3, -3.3, -3.2},
	{-3.2, -3.3, -3.3, -3.3, -3.3, -3.3, -3.3, -3.2},
	{-3.2, -3.3, -3.3, -3.3, -3.3, -3.3, -3.3, -3.2},
	{-3.2, -3.3, -3.3, -3.3, -3.3, -3.3, -3.3, -3.2},
	{-3.2, -3.3, -3.3, -3.3, -3.3, -3.3, -3.3, -3.2},
	{-3.2, -3.2, -3.2, -3.2, -3.2, -3.2, -3.2, -3.2},
}

var pawnValBlack = [8][8]float64{
	{-9.0, -9.0, -9.0, -9.0, -9.0, -9.0, -9.0, -9.0},
	{-2.0, -2.1, -2.1, -2.1, -2.1, -2.1, -2.1, -2.0},
	{-1.3, -1.5, -1.5, -1.5, -1.5, -1.5, -1.5, -1.3},
	{-1.1, -1.1, -1.1, -1.3, -1.3, -1.1, -1.1, -1.1},
	{-1.01, -1.04, -1.1, -1.3, -1.35, -1.1, -1.04, -1.01},
	{-1.01, -1.05, -1.0, -1.2, -1.2, -1.0, -1.05, -1.01},
	{-1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0},
	{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
}
var queenValBlack = [8][8]float64{
	{-9.0, -9.0, -9.0, -9.0, -9.0, -9.0, -9.0, -9.0},
	{-9.0, -9.0, -9.0, -9.0, -9.0, -9.0, -9.0, -9.0},
	{-9.0, -9.0, -9.0, -9.0, -9.0, -9.0, -9.0, -9.0},
	{-9.0, -9.0, -9.0, -9.0, -9.0, -9.0, -9.0, -9.0},
	{-9.0, -9.0, -9.0, -9.0, -9.0, -9.0, -9.0, -9.0},
	{-9.0, -9.0, -9.0, -9.0, -9.0, -9.0, -9.0, -9.0},
	{-9.0, -9.0, -9.0, -9.0, -9.0, -9.0, -9.0, -9.0},
	{-9.0, -9.0, -9.0, -9.5, -9.0, -9.0, -9.0, -9.0},
}

var kingValWhite = [8][8]float64{
	{0, 0, 0.5, 0, 0, 0, 0.5, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
}
var rookValWhite = [8][8]float64{
	{4.95, 4.95, 5.1, 5.1, 5.1, 5.1, 4.95, 4.95},
	{5.0, 5.0, 5.0, 5.0, 5.0, 5.0, 5.0, 5.0},
	{5.0, 5.0, 5.0, 5.0, 5.0, 5.0, 5.0, 5.0},
	{5.0, 5.0, 5.0, 5.0, 5.0, 5.0, 5.0, 5.0},
	{5.0, 5.0, 5.0, 5.0, 5.0, 5.0, 5.0, 5.0},
	{5.0, 5.0, 5.0, 5.0, 5.0, 5.0, 5.0, 5.0},
	{5.0, 5.0, 5.0, 5.0, 5.0, 5.0, 5.0, 5.0},
	{5.0, 5.0, 5.0, 5.0, 5.0, 5.0, 5.0, 5.0},
}

var knightValWhite = [8][8]float64{
	{2.8, 2.9, 2.9, 2.9, 2.9, 2.9, 2.9, 2.8},
	{2.9, 3.0, 3.0, 3.0, 3.0, 3.0, 3.0, 2.9},
	{2.9, 3.0, 3.05, 3.05, 3.05, 3.05, 3.0, 2.9},
	{2.9, 3.0, 3.15, 3.2, 3.2, 3.15, 3.0, 2.9},
	{2.9, 3.0, 3.15, 3.2, 3.2, 3.15, 3.0, 2.9},
	{2.9, 3.0, 3.05, 3.05, 3.05, 3.05, 3.0, 2.9},
	{2.9, 3.0, 3.0, 3.0, 3.0, 3.0, 3.0, 2.9},
	{2.8, 2.9, 2.9, 2.9, 2.9, 2.9, 2.9, 2.8},
}
var bishopValWhite = [8][8]float64{
	{3.2, 3.2, 3.2, 3.2, 3.2, 3.2, 3.2, 3.2},
	{3.2, 3.3, 3.3, 3.3, 3.3, 3.3, 3.3, 3.2},
	{3.2, 3.3, 3.3, 3.3, 3.3, 3.3, 3.3, 3.2},
	{3.2, 3.3, 3.3, 3.3, 3.3, 3.3, 3.3, 3.2},
	{3.2, 3.3, 3.3, 3.3, 3.3, 3.3, 3.3, 3.2},
	{3.2, 3.3, 3.3, 3.3, 3.3, 3.3, 3.3, 3.2},
	{3.2, 3.3, 3.3, 3.3, 3.3, 3.3, 3.3, 3.2},
	{3.2, 3.2, 3.2, 3.2, 3.2, 3.2, 3.2, 3.2},
}

var pawnValWhite = [8][8]float64{
	{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
	{1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0},
	{1.01, 1.05, 1.1, 1.2, 1.2, 1.1, 1.05, 1.01},
	{1.01, 1.04, 1.1, 1.3, 1.35, 1.1, 1.04, 1.01},
	{1.1, 1.1, 1.1, 1.3, 1.3, 1.1, 1.1, 1.1},
	{1.3, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.3},
	{2.0, 2.1, 2.1, 2.1, 2.1, 2.1, 2.1, 2.0},
	{9.0, 9.0, 9.0, 9.0, 9.0, 9.0, 9.0, 9.0},
}

var queenValWhite = [8][8]float64{
	{9.0, 9.0, 9.0, 9.5, 9.0, 9.0, 9.0, 9.0},
	{9.0, 9.0, 9.0, 9.0, 9.0, 9.0, 9.0, 9.0},
	{9.0, 9.0, 9.0, 9.0, 9.0, 9.0, 9.0, 9.0},
	{9.0, 9.0, 9.0, 9.0, 9.0, 9.0, 9.0, 9.0},
	{9.0, 9.0, 9.0, 9.0, 9.0, 9.0, 9.0, 9.0},
	{9.0, 9.0, 9.0, 9.0, 9.0, 9.0, 9.0, 9.0},
	{9.0, 9.0, 9.0, 9.0, 9.0, 9.0, 9.0, 9.0},
	{9.0, 9.0, 9.0, 9.0, 9.0, 9.0, 9.0, 9.0},
}

// ChessBoard defines the main chessboard interface
type ChessBoard interface {
	Init() error
	//InitAndGetMoves() []Move
	Move() FEN
	UnMove() error
	GetAllowedMoves() []Move
	GetCaptureMoves() []Move
	GetOpponentVision() [8][8]bool
	GetPieceMoves() []string
	Evaluate() float64
	Fen() string
}

// Move contains To and From fields relating to the Square number
type Move struct {
	From    [2]int8
	To      [2]int8
	Capture bool
	Check   bool
}

// FEN is a fen string
type FEN struct {
	Position          string
	White             bool
	Black             bool
	CastlePermissions string
	EnPassantSquare   [2]int8
	HalfMove          int8
	FullMove          int8
}

// Board contains pieces
type Board struct {
	Pieces             [8][8]int8
	OpponentVision     [8][8]bool
	OpponentPawnVision [8][8]bool
	Pinned             [8][8]bool
	PinBy              [8][8][2]int8
	CheckPath          [8][8]bool
	NumberOfChecks     uint8

	AllowedMoves []Move
	Color        int8
	White        bool
	Black        bool
	InCheck      bool
	InCheckMate  bool
	InStaleMate  bool
	Fen          FEN
}

// Str2FEN converts a fen string to a FEN object
func Str2FEN(fenStr string) FEN {
	fenParts := strings.Split(fenStr, " ")
	var enPassant [2]int8
	if fenParts[3] != "-" {
		enPassant = Square2Pos(fenParts[3])
	}
	fullMove, _ := strconv.Atoi(fenParts[5])
	return FEN{
		Position:          fenParts[0],
		White:             fenParts[1] == "w",
		Black:             fenParts[1] == "b",
		CastlePermissions: fenParts[2],
		EnPassantSquare:   enPassant,
		HalfMove:          int8(fenParts[4][0]) - '0',
		FullMove:          int8(fullMove),
	}
}

/****************************************** Interface ******************************************/

// Init to board turns a fen string to a Board object
func (b *Board) Init(fen FEN) error {
	b.Pieces = fenPos2Pieces(fen.Position)
	if fen.White {
		b.Color = 1
	} else {
		b.Color = -1
	}
	b.White = fen.White
	b.Black = fen.Black
	b.Fen = fen
	return nil
}

// GetOpponentVision gets the opponent field of vision
func (b *Board) GetOpponentVision() {
	var i int8
	var j int8
	for i = 0; i < 8; i++ {
		for j = 0; j < 8; j++ {
			if b.Pieces[i][j]*b.Color < 0 {
				switch math.Abs(float64(b.Pieces[i][j])) {
				case 1:
					b.kingVision(i, j)
				case 2:
					b.rookVision(i, j)
				case 3:
					b.knightVision(i, j)
				case 4:
					b.bishopVision(i, j)
				case 5:
					b.pawnVision(i, j)
				case 6:
					b.queenVision(i, j)
				}
			}
		}
	}
}

// GetPieceMoves returns a []string of a square's allowed moves.
func (b *Board) GetPieceMoves(fromSquare [2]int8) []string {
	var pieceMoves []string
	for i := 0; i < len(b.AllowedMoves); i++ {
		if b.AllowedMoves[i].From == fromSquare {
			pieceMoves = append(pieceMoves, Pos2Square(b.AllowedMoves[i].From)+"-"+Pos2Square(b.AllowedMoves[i].To))
		}
	}
	return pieceMoves
}

// GetAllowedMoves gets the allowed moves of the playing color in a given position
func (b *Board) GetAllowedMoves() {
	var i int8
	var j int8
	if b.NumberOfChecks > 1 {
		for i = 0; i < 8; i++ {
			for j = 0; j < 8; j++ {
				if math.Abs(float64(b.Pieces[i][j])) == 1 {
					b.kingMoves(i, j)
				}
			}
		}
	} else {
		for i = 0; i < 8; i++ {
			for j = 0; j < 8; j++ {
				if b.Pieces[i][j]*b.Color > 0 {
					switch math.Abs(float64(b.Pieces[i][j])) {
					case 1:
						b.kingMoves(i, j)
					case 2:
						b.rookMoves(i, j)
					case 3:
						b.knightMoves(i, j)
					case 4:
						b.bishopMoves(i, j)
					case 5:
						b.pawnMoves(i, j)
					case 6:
						b.queenMoves(i, j)
					}
				}
			}
		}
	}

	if len(b.AllowedMoves) == 0 {
		if b.InCheck {
			b.InCheckMate = true
		} else {
			b.InStaleMate = true
		}
	} else {
		// Sort allowed moves by
		sort.Slice(b.AllowedMoves, func(i, j int) bool {
			return b.approximateValue(b.AllowedMoves[i]) > b.approximateValue(b.AllowedMoves[j])
		})
	}
}

// GetCaptureMoves gets the capture moves of the playing color in a given position
func (b *Board) GetCaptureMoves() []Move {
	b.GetAllowedMoves()
	captureMoves := []Move{}
	for i := 0; i < len(b.AllowedMoves); i++ {
		if b.Pieces[b.AllowedMoves[i].To[0]][b.AllowedMoves[i].To[1]]*b.Color > 0 {
			captureMoves = append(captureMoves, b.AllowedMoves[i])
		}
	}
	return captureMoves
}

func (b *Board) approximateValue(move Move) float64 {
	var approx float64 = 0
	fromPiece := b.Pieces[move.From[0]][move.From[1]]
	toPiece := b.Pieces[move.To[0]][move.To[1]]
	if toPiece != 0 {
		approx += value(toPiece)*value(toPiece) - value(fromPiece)
	}
	if b.OpponentPawnVision[move.To[0]][move.To[1]] {
		approx -= value(fromPiece)
	}

	if fromPiece%5 == 0 && ((b.White && move.To[0] == 6) || (b.Black && move.To[0] == 1)) {
		approx += value(6)
	}

	if move.Check {
		approx += 10
	}
	return approx
}

// Move makes a move on the board and returns a fen object
func (b *Board) Move(move Move) FEN {
	// Make a copy of the FEN and for later use:
	var fullMove int8
	if b.Black {
		fullMove = b.Fen.FullMove + 1
	} else {
		fullMove = b.Fen.FullMove
	}
	afterMoveFEN := FEN{
		CastlePermissions: b.Fen.CastlePermissions,
		White:             !b.White,
		Black:             !b.Black,
		EnPassantSquare:   [2]int8{0, 0},
		HalfMove:          b.Fen.HalfMove,
		FullMove:          fullMove,
	}

	var copyOfPieces [8][8]int8 = b.Pieces

	piece := b.Pieces[move.From[0]][move.From[1]]
	// Set the 'from' square to empty
	copyOfPieces[move.From[0]][move.From[1]] = 0
	// Check for castling and en passant
	switch math.Abs(float64(piece)) {
	case 1:
		// king, need to check for castles and remove castling permissions
		if move.From[1]-move.To[1] == 2 || move.From[1]-move.To[1] == -2 {
			var rookToCol int
			var rookFromCol int
			if move.To[1] == 2 {
				rookToCol = int(move.To[1]) + 1
				rookFromCol = 0
			} else {
				rookToCol = int(move.To[1]) - 1
				rookFromCol = 7
			}
			var rook int8 = b.Pieces[move.From[0]][rookFromCol]
			copyOfPieces[move.To[0]][rookFromCol] = 0
			copyOfPieces[move.To[0]][rookToCol] = rook
		}

		if b.White {
			afterMoveFEN.CastlePermissions = strings.Replace(b.Fen.CastlePermissions, "K", "", -1)
			afterMoveFEN.CastlePermissions = strings.Replace(afterMoveFEN.CastlePermissions, "Q", "", -1)
		} else {
			afterMoveFEN.CastlePermissions = strings.Replace(b.Fen.CastlePermissions, "k", "", -1)
			afterMoveFEN.CastlePermissions = strings.Replace(afterMoveFEN.CastlePermissions, "q", "", -1)
		}
	case 2:
		// rook, need to remove castling permisions
		if move.From[1] == 0 {
			if b.White {
				afterMoveFEN.CastlePermissions = strings.Replace(b.Fen.CastlePermissions, "Q", "", -1)
			} else {
				afterMoveFEN.CastlePermissions = strings.Replace(b.Fen.CastlePermissions, "q", "", -1)
			}

		} else if move.From[1] == 7 {
			if b.White {
				afterMoveFEN.CastlePermissions = strings.Replace(b.Fen.CastlePermissions, "K", "", -1)
			} else {
				afterMoveFEN.CastlePermissions = strings.Replace(b.Fen.CastlePermissions, "k", "", -1)
			}
		}
	case 5:
		// pawn
		if move.From[1] != move.To[1] && b.Pieces[move.To[0]][move.To[1]] == 0 {
			copyOfPieces[move.From[0]][move.To[1]] = 0
		}
		if move.From[0]-move.To[0] == 2 || move.From[0]-move.To[0] == -2 {
			afterMoveFEN.EnPassantSquare[0] = move.To[0]
			afterMoveFEN.EnPassantSquare[1] = move.To[1]
		}

		if move.To[0] == 7 && b.White {
			piece = 6
		} else if move.To[0] == 0 && b.Black {
			piece = -6
		}

	}
	// Update the 'to' square
	copyOfPieces[move.To[0]][move.To[1]] = piece

	afterMoveFEN.Position = Pieces2FenPos(copyOfPieces)

	// Update Fen
	return afterMoveFEN
}

// Evaluate returns an evaluation of a given position, with >0 meaning white is winning
func (b *Board) Evaluate() float64 {
	eval := float64(0)
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			eval += value(b.Pieces[row][col], int8(row), int8(col))
		}
	}
	return eval * float64(b.Color)
}

/****************************************** Evaluation ******************************************/
func value(p ...int8) float64 {
	if len(p) == 1 {
		switch p[0] {
		case -2:
			return -5
		case -3:
			return -3
		case -4:
			return -3.3
		case -5:
			return -1
		case -6:
			return -9
		case 2:
			return 5
		case 3:
			return 3
		case 4:
			return 3.3
		case 5:
			return 1
		case 6:
			return 9
		default:
			return 0
		}
	} else {
		return getValue(p[0], p[1], p[2])
	}
}

func getValue(piece, row, col int8) float64 {
	switch piece {
	case -1:
		return kingValBlack[row][col]
	case -2:
		return rookValBlack[row][col]
	case -3:
		return knightValBlack[row][col]
	case -4:
		return bishopValBlack[row][col]
	case -5:
		return pawnValBlack[row][col]
	case -6:
		return queenValBlack[row][col]
	case 1:
		return kingValWhite[row][col]
	case 2:
		return rookValWhite[row][col]
	case 3:
		return knightValWhite[row][col]
	case 4:
		return bishopValWhite[row][col]
	case 5:
		return pawnValWhite[row][col]
	case 6:
		return queenValWhite[row][col]
	default:
		return 0
	}
}

/****************************************** Allowed Moves ******************************************/

// AddMove adds a move to the Board's AllowedMoves list
func (b *Board) AddMove(move Move) {
	b.AllowedMoves = append(b.AllowedMoves, move)
}

func (b *Board) pawnMoves(r, c int8) {
	forward := r + b.Color
	twoForward := r + 2*b.Color
	if isInBoard(forward, c) {
		if b.Pieces[forward][c] == 0 {
			if (!b.Pinned[r][c] || b.PinBy[r][c][1] == c) && (!b.InCheck || b.CheckPath[forward][c]) {
				b.AddMove(Move{
					From:  [2]int8{r, c},
					To:    [2]int8{forward, c},
					Check: b.willCheckOpponent(b.Pieces[r][c], forward, c),
				})
			}
		}
		if c != 7 {
			if b.Pieces[forward][c+1]*b.Color < 0 {
				if !b.Pinned[r][c] || (b.PinBy[r][c] == [2]int8{forward, c + 1}) {
					if !b.InCheck || b.CheckPath[forward][c+1] {
						b.AddMove(Move{
							From:  [2]int8{r, c},
							To:    [2]int8{forward, c + 1},
							Check: b.willCheckOpponent(b.Pieces[r][c], forward, c+1),
						})
					}
				}
			}
		}
		if c != 0 {
			if b.Pieces[forward][c-1]*b.Color < 0 {
				if !b.Pinned[r][c] || (b.PinBy[r][c] == [2]int8{forward, c - 1}) {
					if !b.InCheck || b.CheckPath[forward][c-1] {
						b.AddMove(Move{
							From:  [2]int8{r, c},
							To:    [2]int8{forward, c - 1},
							Check: b.willCheckOpponent(b.Pieces[r][c], forward, c-1),
						})
					}
				}
			}
		}
		if b.Fen.EnPassantSquare[0] == r && b.Fen.EnPassantSquare[1] == c+1 {
			if !b.Pinned[r][c] && (!b.InCheck || b.CheckPath[forward][c+1]) {
				b.AddMove(Move{
					From:  [2]int8{r, c},
					To:    [2]int8{forward, c + 1},
					Check: b.willCheckOpponent(b.Pieces[r][c], forward, c+1),
				})
			}
		} else if b.Fen.EnPassantSquare[0] == r && b.Fen.EnPassantSquare[1] == c-1 {
			if !b.Pinned[r][c] && (!b.InCheck || b.CheckPath[forward][c-1]) {
				b.AddMove(Move{
					From:  [2]int8{r, c},
					To:    [2]int8{forward, c - 1},
					Check: b.willCheckOpponent(b.Pieces[r][c], forward, c-1),
				})
			}
		}
	}
	if r == 1 || r == 6 {
		if isInBoard(twoForward, c) && b.Pieces[twoForward][c] == 0 && b.Pieces[forward][c] == 0 {
			//
			if !b.Pinned[r][c] && (!b.InCheck || b.CheckPath[twoForward][c]) {
				b.AddMove(Move{
					From:  [2]int8{r, c},
					To:    [2]int8{twoForward, c},
					Check: b.willCheckOpponent(b.Pieces[r][c], twoForward, c),
				})
			}
		}
	}
}

func (b *Board) rookMoves(r, c int8) {
	if b.Pinned[r][c] {
		if r == b.PinBy[r][c][0] {
			b.straightColMoves(r, c)
		} else if c == b.PinBy[r][c][1] {
			b.straightRowMoves(r, c)
		}
	} else {
		b.straightMoves(r, c)
	}
}

func (b *Board) bishopMoves(r, c int8) {
	if b.Pinned[r][c] {
		if r != b.PinBy[r][c][0] && c != b.PinBy[r][c][1] {
			var positiveDiagonal bool = (c > b.PinBy[r][c][1]) == (r > b.PinBy[r][c][0])
			if positiveDiagonal {
				b.diagonalPositiveMoves(r, c)
			} else {
				b.diagonalNegativeMoves(r, c)
			}
		}
	} else {
		b.diagonalMoves(r, c)
	}
}

func (b *Board) queenMoves(r, c int8) {
	if b.Pinned[r][c] {
		if r != b.PinBy[r][c][0] && c != b.PinBy[r][c][1] {
			var positiveDiagonal bool = (c > b.PinBy[r][c][1]) == (r > b.PinBy[r][c][0])
			if positiveDiagonal {
				b.diagonalPositiveMoves(r, c)
			} else {
				b.diagonalNegativeMoves(r, c)
			}
		} else if r == b.PinBy[r][c][0] {
			b.straightColMoves(r, c)
		} else if c == b.PinBy[r][c][1] {
			b.straightRowMoves(r, c)
		}
	} else {
		b.diagonalMoves(r, c)
		b.straightMoves(r, c)
	}
}

func (b *Board) kingMoves(r, c int8) {
	var row int8
	var col int8
	for row = r - 1; row < r+2; row++ {
		for col = c - 1; col < c+2; col++ {
			if isInBoard(row, col) && (r != row || c != col) {
				if !b.OpponentVision[row][col] && b.Pieces[row][col]*b.Color <= 0 {
					b.AddMove(Move{
						From: [2]int8{r, c},
						To:   [2]int8{row, col},
					})
				}
			}
		}
	}
	// Need Castling
	var castleShort bool
	var castleLong bool
	if b.White {
		castleShort = strings.Contains(b.Fen.CastlePermissions, "K")
		castleLong = strings.Contains(b.Fen.CastlePermissions, "Q")
	} else {
		castleShort = strings.Contains(b.Fen.CastlePermissions, "k")
		castleLong = strings.Contains(b.Fen.CastlePermissions, "q")
	}

	if castleShort {
		if !b.InCheck && !b.OpponentVision[r][6] && !b.OpponentVision[r][5] {
			if b.Pieces[r][6] == 0 && b.Pieces[r][5] == 0 {
				b.AddMove(Move{
					From: [2]int8{r, c},
					To:   [2]int8{r, 6},
				})
			}
		}
	}
	if castleLong {
		if !b.InCheck && !b.OpponentVision[r][2] && !b.OpponentVision[r][3] {
			if b.Pieces[r][1] == 0 && b.Pieces[r][2] == 0 && b.Pieces[r][3] == 0 {
				b.AddMove(Move{
					From: [2]int8{r, c},
					To:   [2]int8{r, 2},
				})
			}
		}
	}
}

func (b *Board) knightMoves(r, c int8) {
	if b.Pinned[r][c] {
		return
	}
	heights := [8]int8{1, 2, 2, 1, -1, -2, -2, -1}
	rows := [8]int8{-2, -1, 1, 2}
	var h int8
	var row int8
	for h = 0; h < 8; h++ {
		row = rows[h%4] + r
		if isInBoard(row, c+heights[h]) && b.Pieces[row][c+heights[h]]*b.Color <= 0 {
			if !b.InCheck || b.CheckPath[row][c+heights[h]] {
				b.AddMove(Move{
					From:  [2]int8{r, c},
					To:    [2]int8{row, c + heights[h]},
					Check: b.willCheckOpponent(b.Pieces[r][c], row, c+heights[h]),
				})
			}
		}
	}
}

/****************************************** Allowed Moves Utils ******************************************/

func (b *Board) straightMoves(r, c int8) {
	b.straightRowMoves(r, c)
	b.straightColMoves(r, c)
}

func (b *Board) straightRowMoves(r, c int8) {
	rowsToTop := makeRange(r+1, 8, 1)
	rowsToBottom := makeRange(r-1, -1, -1)
	colArrTop := []int8{}
	colArrBottom := []int8{}
	var k int8
	for k = 0; k < 8-(r+1); k++ {
		colArrTop = append(colArrTop, c)
	}
	for k = 0; k <= r-1; k++ {
		colArrBottom = append(colArrBottom, c)
	}

	if len(rowsToTop) != 0 && len(colArrTop) != 0 {
		b.getLine(r, c, rowsToTop, colArrTop)
	}
	if len(rowsToBottom) != 0 && len(colArrBottom) != 0 {
		b.getLine(r, c, rowsToBottom, colArrBottom)
	}
}

func (b *Board) straightColMoves(r, c int8) {
	colsToRight := makeRange(c+1, 8, 1)
	colsToLeft := makeRange(c-1, -1, -1)
	rowArrRight := []int8{}
	rowArrLeft := []int8{}
	var k int8
	for k = 0; k < 8-(c+1); k++ {
		rowArrRight = append(rowArrRight, r)
	}
	for k = 0; k <= c-1; k++ {
		rowArrLeft = append(rowArrLeft, r)
	}

	b.getLine(r, c, rowArrRight, colsToRight)
	b.getLine(r, c, rowArrLeft, colsToLeft)
}

func (b *Board) diagonalMoves(r, c int8) {
	b.diagonalPositiveMoves(r, c)
	b.diagonalNegativeMoves(r, c)
}

func (b *Board) diagonalPositiveMoves(r, c int8) {
	rowsToTop := makeRange(r+1, 8, 1)
	rowsToBottom := makeRange(r-1, -1, -1)
	colsToRight := makeRange(c+1, 8, 1)
	colsToLeft := makeRange(c-1, -1, -1)
	b.getLine(r, c, rowsToTop, colsToRight)
	b.getLine(r, c, rowsToBottom, colsToLeft)
}

func (b *Board) diagonalNegativeMoves(r, c int8) {
	rowsToTop := makeRange(r+1, 8, 1)
	rowsToBottom := makeRange(r-1, -1, -1)
	colsToRight := makeRange(c+1, 8, 1)
	colsToLeft := makeRange(c-1, -1, -1)
	b.getLine(r, c, rowsToTop, colsToLeft)
	b.getLine(r, c, rowsToBottom, colsToRight)
}

func (b *Board) getLine(r int8, c int8, rows, cols []int8) {
	var i int = 0
	var depth int
	if len(rows) < len(cols) {
		depth = len(rows)
	} else {
		depth = len(cols)
	}

	for ; i < depth; i++ {
		if b.Pieces[rows[i]][cols[i]]*b.Color > 0 {
			break
		} else if b.Pieces[rows[i]][cols[i]] != 0 {
			if !b.InCheck || b.CheckPath[rows[i]][cols[i]] {
				b.AddMove(Move{
					From:  [2]int8{r, c},
					To:    [2]int8{rows[i], cols[i]},
					Check: b.willCheckOpponent(b.Pieces[r][c], rows[i], cols[i]),
				})
			}
			break
		} else {
			if !b.InCheck || b.CheckPath[rows[i]][cols[i]] {
				b.AddMove(Move{
					From:  [2]int8{r, c},
					To:    [2]int8{rows[i], cols[i]},
					Check: b.willCheckOpponent(b.Pieces[r][c], rows[i], cols[i]),
				})
			}
		}
	}
}

/****************************************** Direct Vision ******************************************/
func (b *Board) willCheckOpponent(t, r, c int8) bool {
	switch int8(math.Abs(float64(t))) {
	case 2:
		// rook
		return b.willCheckRook(r, c)
	case 3:
		// knight
		heights := [8]int8{1, 2, 2, 1, -1, -2, -2, -1}
		rows := [8]int8{-2, -1, 1, 2}
		var h int8
		var row int8
		for h = 0; h < 8; h++ {
			row = rows[h%4] + r
			if isInBoard(row, c+heights[h]) {
				if b.Pieces[row][c+heights[h]]*b.Color == -1 {
					return true
				}
			}
		}
	case 4:
		// bishop
		return b.willCheckBishop(r, c)
	case 5:
		// pawn
		forward := r + b.Color
		if c != 7 {
			if b.Pieces[forward][c+1]*b.Color == -1 {
				return true
			}
		}
		if c != 0 {
			if b.Pieces[forward][c-1]*b.Color == -1 {
				return true
			}
		}
	case 6:
		// queen
		return b.willCheckRook(r, c) || b.willCheckBishop(r, c)
	}
	return false
}

func (b *Board) willCheckRook(r, c int8) bool {
	rowsToTop := makeRange(r+1, 8, 1)
	rowsToBottom := makeRange(r-1, -1, -1)
	colArrTop := []int8{}
	colArrBottom := []int8{}
	var k int8
	for k = 0; k < 8-r; k++ {
		colArrTop = append(colArrTop, c)
	}
	for k = 0; k < r; k++ {
		colArrBottom = append(colArrBottom, c)
	}

	/* if len(rowsToTop) != 0 && len(colArrTop) != 0 {
		b.getLineOfSight(r, c, rowsToTop, colArrTop)
	}
	if len(rowsToBottom) != 0 && len(colArrBottom) != 0 {
		b.getLineOfSight(r, c, rowsToBottom, colArrBottom)
	}*/

	colsToRight := makeRange(c+1, 8, 1)
	colsToLeft := makeRange(c-1, -1, -1)
	rowArrRight := []int8{}
	rowArrLeft := []int8{}
	for k = 0; k < 8-c; k++ {
		rowArrRight = append(rowArrRight, r)
	}
	for k = 0; k < c; k++ {
		rowArrLeft = append(rowArrLeft, r)
	}

	return (b.doesLineCauseCheck(rowsToTop, colArrTop) ||
		b.doesLineCauseCheck(rowsToBottom, colArrBottom) ||
		b.doesLineCauseCheck(rowArrRight, colsToRight) ||
		b.doesLineCauseCheck(rowArrLeft, colsToLeft))
}

func (b *Board) willCheckBishop(r, c int8) bool {
	rowsToTop := makeRange(r+1, 8, 1)
	rowsToBottom := makeRange(r-1, -1, -1)
	colsToRight := makeRange(c+1, 8, 1)
	colsToLeft := makeRange(c-1, -1, -1)
	return (b.doesLineCauseCheck(rowsToTop, colsToRight) ||
		b.doesLineCauseCheck(rowsToBottom, colsToLeft) ||
		b.doesLineCauseCheck(rowsToTop, colsToLeft) ||
		b.doesLineCauseCheck(rowsToBottom, colsToRight))
}

func (b *Board) doesLineCauseCheck(rows, cols []int8) bool {
	var i int = 0
	var depth int
	if len(rows) < len(cols) {
		depth = len(rows)
	} else {
		depth = len(cols)
	}

	for ; i < depth; i++ {
		if b.Pieces[rows[i]][cols[i]]*b.Color == 0 {
			continue
		} else if b.Pieces[rows[i]][cols[i]]*b.Color != -1 {
			return false
		} else {
			return true
		}
	}
	return false
}

/****************************************** Vision ******************************************/

func (b *Board) pawnVision(r, c int8) {
	forward := r - b.Color
	if c != 7 {
		if b.Pieces[forward][c+1]*b.Color == 1 {
			// Check
			b.InCheck = true
			b.NumberOfChecks++
			b.CheckPath[r][c] = true
		}
		b.OpponentPawnVision[forward][c+1] = true
		if !b.OpponentVision[forward][c+1] {
			b.OpponentVision[forward][c+1] = true
		}
	}
	if c != 0 {
		if b.Pieces[forward][c-1]*b.Color == 1 {
			// Check
			b.InCheck = true
			b.NumberOfChecks++
			b.CheckPath[r][c] = true
		}
		b.OpponentPawnVision[forward][c-1] = true
		if !b.OpponentVision[forward][c-1] {
			b.OpponentVision[forward][c-1] = true
		}
	}
}

func (b *Board) rookVision(r, c int8) {
	b.straightVision(r, c)
}

func (b *Board) bishopVision(r, c int8) {
	b.diagonalVision(r, c)
}

func (b *Board) queenVision(r, c int8) {
	b.diagonalVision(r, c)
	b.straightVision(r, c)
}

func (b *Board) kingVision(r, c int8) {
	var row int8
	var col int8
	for row = r - 1; row < r+2; row++ {
		for col = c - 1; col < c+2; col++ {
			if isInBoard(row, col) && (r != row || c != col) {
				if !b.OpponentVision[row][col] {
					b.OpponentVision[row][col] = true
				}
			}
		}
	}
}

func (b *Board) knightVision(r, c int8) {
	heights := [8]int8{1, 2, 2, 1, -1, -2, -2, -1}
	rows := [8]int8{-2, -1, 1, 2}
	var h int8
	var row int8
	for h = 0; h < 8; h++ {
		row = rows[h%4] + r
		if isInBoard(row, c+heights[h]) {
			if b.Pieces[row][c+heights[h]]*b.Color == 1 {
				// Check
				b.InCheck = true
				b.NumberOfChecks++
				b.CheckPath[r][c] = true
			}
			if !b.OpponentVision[row][c+heights[h]] {
				b.OpponentVision[row][c+heights[h]] = true
			}
		}
	}
}

/****************************************** Vision Utils ******************************************/

func (b *Board) straightVision(r, c int8) {
	rowsToTop := makeRange(r+1, 8, 1)
	rowsToBottom := makeRange(r-1, -1, -1)
	colArrTop := []int8{}
	colArrBottom := []int8{}
	var k int8
	for k = 0; k < 8-r; k++ {
		colArrTop = append(colArrTop, c)
	}
	for k = 0; k < r; k++ {
		colArrBottom = append(colArrBottom, c)
	}

	if len(rowsToTop) != 0 && len(colArrTop) != 0 {
		b.getLineOfSight(r, c, rowsToTop, colArrTop)
	}
	if len(rowsToBottom) != 0 && len(colArrBottom) != 0 {
		b.getLineOfSight(r, c, rowsToBottom, colArrBottom)
	}

	colsToRight := makeRange(c+1, 8, 1)
	colsToLeft := makeRange(c-1, -1, -1)
	rowArrRight := []int8{}
	rowArrLeft := []int8{}
	for k = 0; k < 8-c; k++ {
		rowArrRight = append(rowArrRight, r)
	}
	for k = 0; k < c; k++ {
		rowArrLeft = append(rowArrLeft, r)
	}

	b.getLineOfSight(r, c, rowArrRight, colsToRight)
	b.getLineOfSight(r, c, rowArrLeft, colsToLeft)
}

func (b *Board) diagonalVision(r, c int8) {
	rowsToTop := makeRange(r+1, 8, 1)
	rowsToBottom := makeRange(r-1, -1, -1)
	colsToRight := makeRange(c+1, 8, 1)
	colsToLeft := makeRange(c-1, -1, -1)
	b.getLineOfSight(r, c, rowsToTop, colsToRight)
	b.getLineOfSight(r, c, rowsToBottom, colsToLeft)
	b.getLineOfSight(r, c, rowsToTop, colsToLeft)
	b.getLineOfSight(r, c, rowsToBottom, colsToRight)
}

func (b *Board) getLineOfSight(r int8, c int8, rows, cols []int8) {
	var i int = 0
	var depth int
	var depthN int
	if len(rows) < len(cols) {
		depth = len(rows)
	} else {
		depth = len(cols)
	}

	for ; i < depth; i++ {
		if b.Pieces[rows[i]][cols[i]] != 0 {
			if b.Pieces[rows[i]][cols[i]]*b.Color < 0 {
				// Own piece
				if !b.OpponentVision[rows[i]][cols[i]] {
					b.OpponentVision[rows[i]][cols[i]] = true
				}
				//fmt.Println("ownpiece")
				break
			} else if b.Pieces[rows[i]][cols[i]]*b.Color == 1 {
				// Check
				b.InCheck = true
				b.NumberOfChecks++
				b.CheckPath[r][c] = true
				for j := 0; j < i+1; j++ {
					if !b.CheckPath[rows[j]][cols[j]] {
						b.CheckPath[rows[j]][cols[j]] = true
					}
				}
				if !b.OpponentVision[rows[i]][cols[i]] {
					b.OpponentVision[rows[i]][cols[i]] = true
				}

			} else if b.Pieces[rows[i]][cols[i]]*b.Color > 0 {
				// Opponent piece
				if i < depth-1 {

					if len(rows[:i]) < len(cols[i+1:]) {
						depthN = len(rows[i+1:])
					} else {
						depthN = len(cols[i+1:])
					}

					var j int = 0
					for j < depthN && i+j+1 < depth && b.Pieces[rows[i+j+1]][cols[i+j+1]] == 0 {

						j++
					}

					if i+j+1 < depth {
						if b.Pieces[rows[i+j+1]][cols[i+j+1]]*b.Color == 1 {
							b.Pinned[rows[i]][cols[i]] = true
							b.PinBy[rows[i]][cols[i]][0] = r
							b.PinBy[rows[i]][cols[i]][1] = c
						}
					}
				}
				if !b.OpponentVision[rows[i]][cols[i]] {
					b.OpponentVision[rows[i]][cols[i]] = true
				}
				break
			}
		} else {
			if !b.OpponentVision[rows[i]][cols[i]] {
				b.OpponentVision[rows[i]][cols[i]] = true
			}
		}
	}
}

/****************************************** Array Utils ******************************************/

func makeRange(start, end, step int8) []int8 {
	if step == 0 || start == end {
		return []int8{}
	}
	var s []int8
	for start != end {
		s = append(s, start)
		start += step
	}
	return s
}

func isInBoard(r, c int8) bool {
	return c < 8 && c >= 0 && r < 8 && r >= 0
}

/****************************************** FEN Utils ******************************************/

// FenStr2Board to board turns a fen string to a Board object
func FenStr2Board(fenStr string) Board {
	fen := Str2FEN(fenStr)
	var color int8
	if fen.White {
		color = 1
	} else {
		color = -1
	}
	return Board{
		Pieces: fenPos2Pieces(fen.Position),
		White:  fen.White,
		Black:  fen.Black,
		Fen:    fen,
		Color:  color,
	}
}

func Pieces2FenPos(pieces [8][8]int8) string {
	var fenStr string
	var fenRow strings.Builder
	var emptyNum int = 0
	var piece int8
	for r := 0; r <= 7; r++ {
		fenRow.Reset()
		fenRow.Grow(9)
		emptyNum = 0
		for c := 0; c <= 7; c++ {
			piece = pieces[r][c]
			if piece != 0 {
				if emptyNum != 0 {
					fenRow.WriteString(strconv.Itoa(emptyNum))
					emptyNum = 0
				}
				fenRow.WriteString(Piece2Str(piece))
			} else {
				if emptyNum != 0 {
					emptyNum++
				} else {
					emptyNum = 1
				}
			}
		}
		if emptyNum != 0 {
			fenRow.WriteString(strconv.Itoa(emptyNum))
		}

		fenRow.WriteString("/")
		fenStr = fenRow.String() + fenStr
	}
	return fenStr[0 : len(fenStr)-1]
}

func Fen2Str(fen FEN) string {
	var fenColor string
	if fen.White {
		fenColor = "w"
	} else {
		fenColor = "b"
	}
	var enPassantSquare string = Pos2Square(fen.EnPassantSquare)
	return strings.Join([]string{fen.Position, fenColor, fen.CastlePermissions, enPassantSquare, strconv.Itoa(int(fen.HalfMove)), strconv.Itoa(int(fen.FullMove))}, " ")
}

// Square2Pos converts a string representation of a square to a position
func Square2Pos(sq string) [2]int8 {
	return [2]int8{
		int8(sq[1]) - '1',
		int8(sq[0]) - 'a',
	}
}

func fenPos2Pieces(fenPos string) [8][8]int8 {
	var pieces [8][8]int8
	fenRows := strings.Split(fenPos, "/")
	var col int8
	for row := 0; row < 8; row++ {
		col = 0
		for _, c := range fenRows[7-row] {
			if unicode.IsDigit(c) {
				col += int8(c) - '0'
			} else {
				pieces[row][col] = Rune2Piece(c)
				col++
			}
		}
	}
	return pieces
}

// Rune2Piece turns a rune representation of a piece into a number of the same
func Rune2Piece(p rune) int8 {
	switch p {
	case 'K':
		return 1
	case 'R':
		return 2
	case 'N':
		return 3
	case 'B':
		return 4
	case 'P':
		return 5
	case 'Q':
		return 6
	case 'k':
		return -1
	case 'r':
		return -2
	case 'n':
		return -3
	case 'b':
		return -4
	case 'p':
		return -5
	case 'q':
		return -6
	}
	return 0
}

// Pos2Square returns a square from a position
func Pos2Square(p [2]int8) string {
	// row, col
	var cols string = "abcdefgh"
	return string(cols[p[1]]) + strconv.Itoa(int(p[0])+1)
}

// Piece2Str takes a int8 representation of a piece and returns a string
func Piece2Str(p int8) string {
	switch p {
	case 1:
		return "K"
	case 2:
		return "R"
	case 3:
		return "N"
	case 4:
		return "B"
	case 5:
		return "P"
	case 6:
		return "Q"
	case -1:
		return "k"
	case -2:
		return "r"
	case -3:
		return "n"
	case -4:
		return "b"
	case -5:
		return "p"
	case -6:
		return "q"
	}
	return ""
}
