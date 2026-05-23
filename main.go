package main

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TickMsg time.Time

func tick(level int) tea.Cmd {
	delay := time.Second / time.Duration(level+1)
	return tea.Tick(delay, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

// Main game model
type Model struct {
	board        [Height][Width]int
	currentPiece Piece
	nextPiece    Piece
	score        int
	level        int
	lines        int
	gameOver     bool
	paused       bool
}

func initialModel() Model {
	return Model{
		currentPiece: spawnPiece(),
		nextPiece:    spawnPiece(),
		level:        0,
		score:        0,
	}
}

func (m Model) Init() tea.Cmd {
	return tick(m.level)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "p":
			m.paused = !m.paused
			if !m.paused && !m.gameOver {
				return m, tick(m.level)
			}
			return m, nil
		}

		if m.gameOver || m.paused {
			return m, nil
		}

		switch msg.String() {
		case "left", "h":
			p := m.currentPiece
			p.x--
			if isValidPosition(&m.board, p) {
				m.currentPiece = p
			}
		case "right", "l":
			p := m.currentPiece
			p.x++
			if isValidPosition(&m.board, p) {
				m.currentPiece = p
			}
		case "down", "j":
			p := m.currentPiece
			p.y++
			if isValidPosition(&m.board, p) {
				m.currentPiece = p
			} else {
				m = lockPiece(m)
			}
		case "up", "k":
			p := rotatePiece(m.currentPiece)
			if isValidPosition(&m.board, p) {
				m.currentPiece = p
			}
		case " ":
			// hard drop (instant fall)
			for {
				p := m.currentPiece
				p.y++
				if isValidPosition(&m.board, p) {
					m.currentPiece = p
				} else {
					break
				}
			}
			m = lockPiece(m)
		}

	case TickMsg:
		if m.gameOver || m.paused {
			return m, nil
		}
		p := m.currentPiece
		p.y++
		if isValidPosition(&m.board, p) {
			m.currentPiece = p
		} else {
			m = lockPiece(m)
		}
		return m, tick(m.level)
	}

	return m, nil
}

func lockPiece(m Model) Model {
	mergePiece(&m.board, m.currentPiece)
	cleared := clearLines(&m.board)

	if cleared > 0 {
		m.lines += cleared
		points := 0
		switch cleared {
		case 1:
			points = 100
		case 2:
			points = 300
		case 3:
			points = 500
		case 4:
			points = 800
		}
		m.score += points * (m.level + 1)
		m.level = m.lines / 10
	}

	m.currentPiece = m.nextPiece
	m.nextPiece = spawnPiece()

	if !isValidPosition(&m.board, m.currentPiece) {
		m.gameOver = true
	}

	return m
}

func (m Model) View() string {
	board := renderBoard(m)
	sidebar := renderSidebar(m)
	return lipgloss.JoinHorizontal(lipgloss.Top, board, sidebar)
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
