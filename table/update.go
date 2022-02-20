package table

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) moveHighlightUp() {
	m.rowCursorIndex--

	if m.rowCursorIndex < 0 {
		m.rowCursorIndex = len(m.rows) - 1
	}
}

func (m *Model) moveHighlightDown() {
	m.rowCursorIndex++

	if m.rowCursorIndex >= len(m.rows) {
		m.rowCursorIndex = 0
	}
}

func (m *Model) toggleSelect() {
	if !m.selectableRows || len(m.rows) == 0 {
		return
	}

	rows := make([]Row, len(m.rows))
	copy(rows, m.rows)

	rows[m.rowCursorIndex].selected = !rows[m.rowCursorIndex].selected

	m.rows = rows

	m.selectedRows = []Row{}

	for _, row := range m.rows {
		if row.selected {
			m.selectedRows = append(m.selectedRows, row)
		}
	}
}

// Update responds to input from the user or other messages from Bubble Tea.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if !m.focused {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.RowDown):
			m.moveHighlightDown()

		case key.Matches(msg, m.keyMap.RowUp):
			m.moveHighlightUp()

		case key.Matches(msg, m.keyMap.RowSelectToggle):
			m.toggleSelect()
		}
	}

	return m, nil
}