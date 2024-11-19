package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"time"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)
var todoAlreadyExist bool

type Tab int

const (
	TableTab = iota
	AddTaskTab
)

type model struct {
	todos     *Todos
	table     table.Model
	activeTab Tab
	input     string
}

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
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.table.SetWidth(msg.Width - h)
		m.table.SetHeight(msg.Height - v*2)
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			m.activeTab = (m.activeTab + 1) % 2
			m.input = ""
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	switch m.activeTab {

	case TableTab:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "d":
				index := m.table.Cursor()
				m.todos.delete(index)
				m.updateTable()
			case "enter":
				index := m.table.Cursor()
				m.todos.toggle(index)
				m.updateTable()
			}
		}
		var cmd tea.Cmd
		m.table, cmd = m.table.Update(msg)
		return m, cmd

	case AddTaskTab:
		if keyMsg, ok := msg.(tea.KeyMsg); ok {
			switch keyMsg.String() {
			case "enter":
				if err := m.todos.add(m.input); err != nil {
					todoAlreadyExist = true
				}
				// m.table, _ = m.table.Update(msg)
				m.updateTable()
				m.input = ""
			case "backspace":
				if len(m.input) > 0 {
					m.input = m.input[:len(m.input)-1]
				}
			default:
				m.input += keyMsg.String()
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	var tabs string
	switch m.activeTab {
	case TableTab:
		tabs = "[ Table ]  Add Task"
	case AddTaskTab:
		tabs = "Table [ Add Task ]"
	}

	var content string
	switch m.activeTab {
	case TableTab:
		content = m.table.View()
	case AddTaskTab:
		if todoAlreadyExist {
			content = fmt.Sprintf("New task: %s", m.input) + "\nToDo with that name already exist"
			todoAlreadyExist = false
		} else {
			content = fmt.Sprintf("New task: %s", m.input) + "\nPress enter to save"
		}
	}

	return lipgloss.NewStyle().
		Bold(true).Foreground(lipgloss.Color("34")).Render(tabs) + "\n\n" + content
}
