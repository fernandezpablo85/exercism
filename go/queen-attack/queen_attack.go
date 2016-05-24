package queenattack

import (
	"errors"
	"strings"
)

var (
	SameSquare     = errors.New("queens can't be in the same square")
	BadCoordinates = errors.New("Invalid coordinates")
)

const (
	ValidColumns = "abcdef"
	ValidRows    = "12345678"
)

func CanQueenAttack(first, second string) (bool, error) {
	if first == second {
		return false, SameSquare
	}

	if len(first) != 2 || len(second) != 2 {
		return false, BadCoordinates
	}
	firstColumn, firstRow, err := extractCoordinates(first)
	if err != nil {
		return false, BadCoordinates
	}
	secondColumn, secondRow, err := extractCoordinates(second)
	if err != nil {
		return false, BadCoordinates
	}

	sameColumn := firstColumn == secondColumn
	sameRow := firstRow == secondRow
	inDiagonal := abs(firstColumn-secondColumn) == abs(firstRow-secondRow)

	return sameColumn || sameRow || inDiagonal, nil
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	} else {
		return a
	}
}

func extractCoordinates(position string) (int, int, error) {
	column := strings.Index(ValidColumns, string(position[0]))
	row := strings.Index(ValidRows, string(position[1]))

	if column >= 0 && row >= 0 {
		return column, row, nil
	} else {
		return -1, -1, BadCoordinates
	}
}
