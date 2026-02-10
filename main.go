package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {

	startKey := detectPreferredCountryKey()

	startIdx := 0
	for i, k := range countryKeys {
		if k == startKey {
			startIdx = i
			break
		}
	}

	l := getLocale(countryKeys[startIdx])
	columns := []table.Column{
		{Title: "", Width: 2},
		{Title: "", Width: 30},
		{Title: "", Width: 25},
		{Title: "", Width: 2},
	}

	m := model{
		configIdx:    startIdx,
		year:         2026,
		mondayFirst:  l.DefaultMon,
		outputToFile: false,
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	m.table = t
	m.syncTable()

	s := table.DefaultStyles()
	s.Header = lipgloss.NewStyle().Height(0).MaxHeight(0).Padding(0).Margin(0)
	s.Header.UnsetBorderStyle()
	s.Selected = s.Selected.Foreground(lipgloss.Color("229")).Background(lipgloss.Color("57")).Bold(false)
	m.table.SetStyles(s)

	p := tea.NewProgram(m)
	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if res, ok := finalModel.(model); ok && res.finalOutput != "" {
		fmt.Println(res.finalOutput)
	}
}