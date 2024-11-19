package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"time"
)

type model struct {
	todos *Todos
	table table.Model
}

// init() tea.Cmd
// update(msg tea.Msg) (tea.Model, tea.Cmd)
// view() string

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) updateTable() {
	rows := make([]table.Row, len(*m.todos))
	for i, todo := range *m.todos {
		var status string
		if todo.Completed {
			status = "+"
			rows[i] = table.Row{fmt.Sprint(i), todo.Title, status, todo.AddedAt.Format(time.RFC1123), todo.CompletedAt.Format(time.RFC1123)}
		} else {
			status = "x"
			rows[i] = table.Row{fmt.Sprint(i), todo.Title, status, todo.AddedAt.Format(time.RFC1123), ""}
		}
	}
	m.table.SetRows(rows)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			if len(*m.todos) > 0 {
				index := m.table.Cursor()
				_ = m.todos.toggle(index)
				m.updateTable()
			}
		}
	}
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		Render(m.table.View()) + "\n"
}
