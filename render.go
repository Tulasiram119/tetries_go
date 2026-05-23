package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	boardStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))

	sidebarStyle = lipgloss.NewStyle().
		MarginLeft(2)

	gameOverStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF0000")).
		Bold(true)

	titleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FFFF")).
		Bold(true)
)

func renderBoard(m Model) string {
	var b strings.Builder

	// Create a combined board with the current piece
	displayBoard := m.board
	for r, row := range m.currentPiece.shape {
		for c, val := range row {
			if val != 0 {
				br := m.currentPiece.y + r
				bc := m.currentPiece.x + c
				if br >= 0 && br < Height && bc >= 0 && bc < Width {
					displayBoard[br][bc] = m.currentPiece.id
				}
			}
		}
	}

	for r := 0; r < Height; r++ {
		for c := 0; c < Width; c++ {
			id := displayBoard[r][c]
			if id == 0 {
				b.WriteString(" .") // Empty space representation
			} else {
				// Find color
				var color string
				for _, p := range Tetrominoes {
					if p.id == id {
						color = p.color
						break
					}
				}
				style := lipgloss.NewStyle().Foreground(lipgloss.Color(color))
				b.WriteString(style.Render("██"))
			}
		}
		if r < Height-1 {
			b.WriteString("\n")
		}
	}

	return boardStyle.Render(b.String())
}

func renderSidebar(m Model) string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("TERMINAL TETRIS\n\n"))

	b.WriteString(fmt.Sprintf("Score: %d\n", m.score))
	b.WriteString(fmt.Sprintf("Level: %d\n\n", m.level))

	b.WriteString("Next Piece:\n")

	// Render next piece
	for r := 0; r < 4; r++ { // max height of piece is 4
		for c := 0; c < 4; c++ { // max width is 4
			if r < len(m.nextPiece.shape) && c < len(m.nextPiece.shape[r]) && m.nextPiece.shape[r][c] != 0 {
				style := lipgloss.NewStyle().Foreground(lipgloss.Color(m.nextPiece.color))
				b.WriteString(style.Render("██"))
			} else {
				b.WriteString("  ")
			}
		}
		b.WriteString("\n")
	}

	b.WriteString("\nControls:\n")
	b.WriteString("←/h : Move Left\n")
	b.WriteString("→/l : Move Right\n")
	b.WriteString("↓/j : Soft Drop\n")
	b.WriteString("↑/k : Rotate\n")
	b.WriteString("Space : Hard Drop\n")
	b.WriteString("p   : Pause\n")
	b.WriteString("q   : Quit\n")

	if m.gameOver {
		b.WriteString("\n" + gameOverStyle.Render("GAME OVER!") + "\nPress 'q' to quit.")
	} else if m.paused {
		b.WriteString("\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFF00")).Render("PAUSED") + "\nPress 'p' to resume.")
	}

	return sidebarStyle.Render(b.String())
}
