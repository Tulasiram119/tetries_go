package main

import (
	"math/rand"
)

// A Tetromino piece
type Piece struct {
	id    int
	shape [][]int // 2D grid of cells
	x, y  int     // position on board
	color string  // lipgloss color
}

var Tetrominoes = []Piece{
	{id: 1, shape: [][]int{{1, 1, 1, 1}}, color: "#00FFFF"},           // I
	{id: 2, shape: [][]int{{1, 1}, {1, 1}}, color: "#FFFF00"},         // O
	{id: 3, shape: [][]int{{0, 1, 0}, {1, 1, 1}}, color: "#800080"},   // T
	{id: 4, shape: [][]int{{1, 0}, {1, 1}, {0, 1}}, color: "#00FF00"}, // S
	{id: 5, shape: [][]int{{0, 1}, {1, 1}, {1, 0}}, color: "#FF0000"}, // Z
	{id: 6, shape: [][]int{{1, 0}, {1, 0}, {1, 1}}, color: "#FFA500"}, // L
	{id: 7, shape: [][]int{{0, 1}, {0, 1}, {1, 1}}, color: "#0000FF"}, // J
}

func spawnPiece() Piece {
	p := Tetrominoes[rand.Intn(len(Tetrominoes))]
	// Deep copy the shape
	newShape := make([][]int, len(p.shape))
	for i := range p.shape {
		newShape[i] = make([]int, len(p.shape[i]))
		copy(newShape[i], p.shape[i])
	}

	// Start at top center
	return Piece{
		id:    p.id,
		shape: newShape,
		x:     Width/2 - len(newShape[0])/2,
		y:     0,
		color: p.color,
	}
}

func rotatePiece(p Piece) Piece {
	// transpose + reverse rows to rotate 90°
	rows := len(p.shape)
	cols := len(p.shape[0])

	newShape := make([][]int, cols)
	for i := range newShape {
		newShape[i] = make([]int, rows)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			newShape[c][rows-1-r] = p.shape[r][c]
		}
	}

	return Piece{
		id:    p.id,
		shape: newShape,
		x:     p.x,
		y:     p.y,
		color: p.color,
	}
}
