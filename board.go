package main

// Board dimensions
const Width = 10
const Height = 20

func isValidPosition(board *[Height][Width]int, p Piece) bool {
	for r, row := range p.shape {
		for c, val := range row {
			if val != 0 {
				br := p.y + r
				bc := p.x + c
				if bc < 0 || bc >= Width || br >= Height {
					return false // out of bounds
				}
				if br >= 0 && board[br][bc] != 0 {
					return false // collision
				}
			}
		}
	}
	return true
}

func mergePiece(board *[Height][Width]int, p Piece) {
	for r, row := range p.shape {
		for c, val := range row {
			if val != 0 {
				br := p.y + r
				bc := p.x + c
				if br >= 0 && br < Height && bc >= 0 && bc < Width {
					board[br][bc] = p.id
				}
			}
		}
	}
}

func clearLines(board *[Height][Width]int) int {
	linesCleared := 0
	for r := Height - 1; r >= 0; r-- {
		full := true
		for c := 0; c < Width; c++ {
			if board[r][c] == 0 {
				full = false
				break
			}
		}

		if full {
			linesCleared++
			// Shift everything down
			for shiftR := r; shiftR > 0; shiftR-- {
				for c := 0; c < Width; c++ {
					board[shiftR][c] = board[shiftR-1][c]
				}
			}
			// Clear top row
			for c := 0; c < Width; c++ {
				board[0][c] = 0
			}
			r++ // check the same row again after shift
		}
	}
	return linesCleared
}
