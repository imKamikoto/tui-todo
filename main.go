package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
	"time"
)

func main() {

	fileName := "storage.json"

	todos := Todos{}
	storage, error := NewStorage[Todos](fileName)
	if error != nil {
		fmt.Println("Cannot read data from file - " + fileName)
	}
	if err := storage.Load(&todos); err != nil {
		fmt.Println("Error while reading data - " + err.Error())
	}

	colums := []table.Column{
		{Title: "#", Width: 2},
		{Title: "Title", Width: 30},
		{Title: "Status", Width: 10},
		{Title: "Added At", Width: 40},
		{Title: "Completed At", Width: 40},
	}

	rows := make([]table.Row, len(todos))
	for i, todo := range todos {
		var status string
		if todo.Completed {
			status = "+"
			rows[i] = table.Row{fmt.Sprint(i), todo.Title, status, todo.AddedAt.Format(time.RFC1123), todo.CompletedAt.Format(time.RFC1123)}
		} else {
			status = "x"
			rows[i] = table.Row{fmt.Sprint(i), todo.Title, status, todo.AddedAt.Format(time.RFC1123), ""}
		}
	}

	myTable := table.New(
		table.WithColumns(colums),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(15),
	)

	style := table.DefaultStyles()
	style.Header = style.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	style.Selected = style.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	myTable.SetStyles(style)

	fmt.Print("before start")
	m := model{todos: &todos, table: myTable}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program - " + err.Error())
		os.Exit(1)
	}
	// if err := storage.Load(todos); err != nil {
	// 	fmt.Println("Error while saving data - " + err.Error())
	// }
}
