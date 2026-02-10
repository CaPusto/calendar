package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func isLeap(y int) bool {
	return (y%4 == 0 && y%100 != 0) || (y%400 == 0)
}

func daysInMonth(m, y int) int {
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if m == 1 && isLeap(y) {
		return 29
	}
	return days[m]
}

func firstDayOfYear(year int) int {
	t := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	return int(t.Weekday()) // 0=Sunday … 6=Saturday
}

func PrintCalendar(year int, mondayFirst bool, useColors bool, months []string, header string) string {
	var sb strings.Builder

	weekendStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("204"))

	colWidth := 21
	separator := "   " // 3 пробела между месяцами
	totalWidth := 3*colWidth + 2*len(separator)

	// Заголовок года по центру
	yearStr := fmt.Sprintf("%d", year)
	padding := (totalWidth - len(yearStr)) / 2
	sb.WriteString("\n" + strings.Repeat(" ", padding) + yearStr + "\n\n")

	currentDow := firstDayOfYear(year)
	weekStart := 0
	if mondayFirst {
		weekStart = 1
	}

	for row := 0; row < 4; row++ {
		mIdx := []int{row * 3, row*3 + 1, row*3 + 2}

		// Названия месяцев
		sb.WriteString(fmt.Sprintf(" %-21s   %-21s   %-21s\n",
			months[mIdx[0]], months[mIdx[1]], months[mIdx[2]]))

		// Заголовок дней недели
		sb.WriteString(fmt.Sprintf(" %-21s   %-21s   %-21s\n", header, header, header))

		// Дни в месяцах и стартовые дни недели
		daysIn := make([]int, 3)
		startDow := make([]int, 3)
		tempDow := currentDow
		for col := 0; col < 3; col++ {
			daysIn[col] = daysInMonth(mIdx[col], year)
			startDow[col] = tempDow
			tempDow = (tempDow + daysIn[col]) % 7
		}

		// Печать 6 недель
		for w := 0; w < 6; w++ {
			for col := 0; col < 3; col++ {
				offset := (startDow[col] - weekStart + 7) % 7

				for d := 0; d < 7; d++ {
					dayNum := w*7 + d - offset + 1

					if dayNum < 1 || dayNum > daysIn[col] {
						sb.WriteString("   ")
					} else {
						isWeekend := false
						if mondayFirst {
							if d == 5 || d == 6 { // сб + вс
								isWeekend = true
							}
						} else {
							if d == 0 || d == 6 { // вс + сб
								isWeekend = true
							}
						}

						dayStr := fmt.Sprintf("%3d", dayNum)
						if useColors && isWeekend {
							sb.WriteString(weekendStyle.Render(dayStr))
						} else {
							sb.WriteString(dayStr)
						}
					}
				}

				if col < 2 {
					sb.WriteString(separator)
				}
			}
			sb.WriteString("\n")
		}

		sb.WriteString("\n")
		currentDow = tempDow
	}

	return sb.String()
}
