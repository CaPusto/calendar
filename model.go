package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table        table.Model
	configIdx    int
	year         int
	mondayFirst  bool
	outputToFile bool
	finalOutput  string
}

const goRow = 4

func (m model) Init() tea.Cmd { return nil }

func (m *model) handleFinalize() {
	l := getLocale(countryKeys[m.configIdx])
	header := l.HeaderSun
	if m.mondayFirst {
		header = l.HeaderMon
	}

	calendar := PrintCalendar(m.year, m.mondayFirst, !m.outputToFile, l.Months, header)

	if m.outputToFile {
		err := os.WriteFile("calendar.txt", []byte(calendar), 0644)
		if err != nil {
			m.finalOutput = fmt.Sprintf(l.ErrorMsg, err)
		} else {
			m.finalOutput = l.SuccessMsg
		}
	} else {
		m.finalOutput = calendar
	}
}

func (m *model) syncTable() {
	l := getLocale(countryKeys[m.configIdx])

	weekVal := l.SunVal
	if m.mondayFirst {
		weekVal = l.MonVal
	}
	outVal := l.ConsoleVal
	if m.outputToFile {
		outVal = l.FileVal
	}

	rows := []table.Row{
		{"", l.YearLabel, strconv.Itoa(m.year), ""},
		{"", l.WeekDayLabel, weekVal, ""},
		{"", l.OutputLabel, outVal, ""},
		{"", l.CountryLabel, countryKeys[m.configIdx], ""},
		{"", l.GoLabel, ">>>", ""},
	}
	m.table.SetRows(rows)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			if m.table.Cursor() == goRow {
				m.handleFinalize()
				return m, tea.Quit
			}

		case "left", "right":
			currRow := m.table.Cursor()
			isRight := msg.String() == "right"

			switch currRow {
			case 0: // год
				if isRight && m.year < 2100 {
					m.year++
				} else if !isRight && m.year > 1900 {
					m.year--
				}
			case 1: // день недели
				m.mondayFirst = !m.mondayFirst
			case 2: // вывод
				m.outputToFile = !m.outputToFile
			case 3: // страна
				delta := 1
				if !isRight {
					delta = -1
				}
				m.configIdx = (m.configIdx + delta + len(countryKeys)) % len(countryKeys)
				m.mondayFirst = getLocale(countryKeys[m.configIdx]).DefaultMon
			}
			m.syncTable()
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	l := getLocale(countryKeys[m.configIdx])

	// Считаем ширину для заголовка
	totalWidth := 0
	for _, col := range m.table.Columns() {
		totalWidth += col.Width + 2 // + padding
	}

	headerText := l.Title
	header := lipgloss.NewStyle().Bold(true).Align(lipgloss.Center).Width(totalWidth).Render(headerText)
	divider := lipgloss.NewStyle().Foreground(lipgloss.Color("99")).Render(strings.Repeat("═", totalWidth))

	// Убираем стандартный заголовок таблицы (пустая строка)
	lines := strings.Split(m.table.View(), "\n")
	tableView := strings.Join(lines[1:], "\n")

	combinedView := lipgloss.JoinVertical(lipgloss.Left, header, divider, tableView)

	return baseStyle.Render(combinedView) + "\n " + l.HelpMsg
}
