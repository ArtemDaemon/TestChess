package testchess

import "fmt"

type EChessmanType int
type ESide int

const (
	RookType    EChessmanType = 1
	BishopoType EChessmanType = 2
	White       ESide         = 0
	Black       ESide         = 1
)

type iChessman interface {
	GetPosition() *TChessField
	GoToPosition(target *TChessField) *TChessMove
}

type TChessman struct {
	Type     EChessmanType
	Position *TChessField
	Side     ESide
}

func (c *TChessman) GetPosition() *TChessField {
	return c.Position
}

func (c *TChessman) GoToPosition(target *TChessField) *TChessMove {
	return &TChessMove{}
}

type TChessField struct {
	row  rune
	col  rune
	busy *TChessman
}

func NewChessField(row, col rune) *TChessField {
	return &TChessField{row: row, col: col}
}

func (f *TChessField) GetRow() rune {
	return f.row
}

func (f *TChessField) GetCol() rune {
	return f.col
}

func (f *TChessField) IsBusy() *TChessman {
	return f.busy
}

type TChessMove struct {
	From     *TChessField
	To       *TChessField
	Mover    *TChessman
	Captured *TChessman
}

func (m *TChessMove) asString() string {
	pieceNotation := map[EChessmanType]string{
		RookType:    "R",
		BishopoType: "B",
	}
	capture := ""
	if m.Captured != nil {
		capture = "x"
	}
	return fmt.Sprintf("%s%s%s%d",
		pieceNotation[m.Mover.Type],
		capture,
		string(m.To.col),
		m.To.row-'0',
	)
}

// Классы для тестирования
type Rook struct {
	TChessman
}

func NewRook(pos *TChessField, side ESide) *Rook {
	r := &Rook{TChessman{Type: 1, Position: pos, Side: side}}
	pos.busy = &r.TChessman
	return r
}

func (r *Rook) GoToPosition(target *TChessField) *TChessMove {
	if target.IsBusy() != nil && target.IsBusy().Side == r.Side {
		return nil
	}
	if r.Position.row == target.row || r.Position.col == target.col {
		move := &TChessMove{
			From:     r.Position,
			To:       target,
			Mover:    &r.TChessman,
			Captured: target.IsBusy(),
		}
		r.Position.busy = nil
		target.busy = &r.TChessman
		r.Position = target
		return move
	}
	return nil
}

type Bishop struct {
	TChessman
}

func NewBishop(pos *TChessField, side ESide) *Bishop {
	b := &Bishop{TChessman{Type: 2, Position: pos, Side: side}}
	pos.busy = &b.TChessman
	return b
}

func (b *Bishop) GoToPosition(target *TChessField) *TChessMove {
	if target.IsBusy() != nil && target.IsBusy().Side == b.Side {
		return nil
	}
	if abs(int(b.Position.row)-int(target.row)) == abs(int(b.Position.col)-int(target.col)) {
		move := &TChessMove{
			From:     b.Position,
			To:       target,
			Mover:    &b.TChessman,
			Captured: target.IsBusy(),
		}
		b.Position.busy = nil
		target.busy = &b.TChessman
		b.Position = target
		return move
	}
	return nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
