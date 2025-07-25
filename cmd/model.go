package cmd

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct{
	matrix [][]int
	cursor int
	selected map[int]struct{}
}

func InitialModel() model {
	return model{
		matrix: [][]int{
			{2, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
	}
}

func (m model) Init() tea.Cmd{
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			for col := 0; col < 4; col++ {
				for row := 1; row < 4; row++ {
					if m.matrix[row][col] != 0 {
						m.matrix[row-1][col] = m.matrix[row][col]
						m.matrix[row][col] = 0
					}
				}
			}

		case "down", "j":
			for col := 0; col < 4; col++ {
				for row := 2; row >= 0; row-- {
					if m.matrix[row][col] != 0 {
						m.matrix[row+1][col] = m.matrix[row][col]
						m.matrix[row][col] = 0
					}
				}
			}

		case "left", "h":
			for row := 0; row < 4; row++ {
				for col := 1; col < 4; col++ {
					if m.matrix[row][col] != 0 {
						m.matrix[row][col-1] = m.matrix[row][col]
						m.matrix[row][col] = 0
					}
				}
			}

		case "right", "l":
			for row := 0; row < 4; row++ {
				for col := 2; col >= 0; col-- {
					if m.matrix[row][col] != 0 {
						m.matrix[row][col+1] = m.matrix[row][col]
						m.matrix[row][col] = 0
					}
				}
			}
		}
	}
	return m, nil
}


func (m model) View() string {
	var out strings.Builder
	for i, row := range m.matrix {
		for _, cell := range row {
			fmt.Fprintf(&out, " %d", cell)
		}
		if i < len(m.matrix)-1 {
			out.WriteString("\n")
		}
	}
	return out.String()
}
