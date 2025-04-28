package testchess

import (
	"testing"
)

func TestRookGoToPosition(t *testing.T) {
	start := NewChessField('1', 'a')
	rook := NewRook(start, White)
	// Корректное вертикальное движение
	target1 := NewChessField('8', 'a')
	if rook.GoToPosition(target1) == nil {
		t.Error("Ладья может ходить вертикально")
	}
	// Корректное горизонтальное движение
	target2 := NewChessField('8', 'h')
	if rook.GoToPosition(target2) == nil {
		t.Error("Ладья может ходить горизонтально")
	}
	// Некорректное движение по диагонали
	target3 := NewChessField('3', 'c')
	if rook.GoToPosition(target3) != nil {
		t.Error("Ладья не может ходить по диагонали")
	}
}

func TestBishopGoToPosition(t *testing.T) {
	start := NewChessField('1', 'c')
	bishop := NewBishop(start, White)
	// Корректное движение по диагонали
	target1 := NewChessField('3', 'e')
	if bishop.GoToPosition(target1) == nil {
		t.Error("Слон может ходить по диагонали")
	}
	// Некорректное горизонтальное движение
	target2 := NewChessField('8', 'c')
	if bishop.GoToPosition(target2) != nil {
		t.Error("Слон не может ходить горизонтально")
	}
	// Некорректное вертикальное движение
	target3 := NewChessField('1', 'a')
	if bishop.GoToPosition(target3) != nil {
		t.Error("Слон не может ходить вертикально")
	}
}

func TestChessFieldBusy(t *testing.T) {
	field := NewChessField('1', 'a')
	if field.IsBusy() != nil {
		t.Error("Поле не может быть занятым по умолчанию")
	}
	NewRook(field, White)
	if field.IsBusy() == nil {
		t.Error("Поле должно стать занятым после помещения в него фигуры")
	}
}

func TestGetPosition(t *testing.T) {
	field := NewChessField('2', 'b')
	rook := NewRook(field, White)
	if rook.GetPosition() != field {
		t.Error("GetPosition должен возвращать корректное поле")
	}
	bishop := NewBishop(field, White)
	if bishop.GetPosition() != field {
		t.Error("GetPosition должен возвращать корректное поле")
	}
}

func TestBlockedByAlly(t *testing.T) {
	start := NewChessField('1', 'a')
	target := NewChessField('2', 'a')
	rook := NewRook(start, White)
	NewRook(target, White)

	if rook.GoToPosition(target) != nil {
		t.Error("Ладья не может встать на поле, занятое союзной фигурой")
	}

	bishop := NewBishop(start, White)
	if bishop.GoToPosition(target) != nil {
		t.Error("Слон не может встать на поле, занятое союзной фигурой")
	}
}

func TestRookTakesEnemy(t *testing.T) {
	start := NewChessField('1', 'a')
	target := NewChessField('2', 'a')
	rook := NewRook(start, White)
	enemy := NewRook(target, Black)
	move := rook.GoToPosition(target)
	if move == nil {
		t.Fatal("Ладья должна съесть вражескую фигуру")
	}
	if move.Captured != &enemy.TChessman {
		t.Error("Съеденная фигура должна быть enemy")
	}
	if move.asString() != "Rxa2" {
		t.Errorf("Нотация движения должна быть Rxa2, вместо этого: %s", move.asString())
	}
}

func TestBishopTakesEnemy(t *testing.T) {
	start := NewChessField('1', 'c')
	target := NewChessField('3', 'e')
	rook := NewBishop(start, White)
	enemy := NewRook(target, Black)
	move := rook.GoToPosition(target)
	if move == nil {
		t.Fatal("Слон должен съесть вражескую фигуру")
	}
	if move.Captured != &enemy.TChessman {
		t.Error("Съеденная фигура должна быть enemy")
	}
	if move.asString() != "Bxe3" {
		t.Errorf("Нотация движения должна быть Bxe3, вместо этого: %s", move.asString())
	}
}
