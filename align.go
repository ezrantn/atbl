package atbl

import "strings"

func formatCell(text string, width int, align string) string {
	padding := width - len(text)
	
	switch align {
	case Left:
		return " " + text + strings.Repeat(" ", padding) + " "
	case Right:
		return " " + strings.Repeat(" ", padding) + text + " "
	case Center:
		leftPad := padding / 2
		rightPad := padding - leftPad
		return " " + strings.Repeat(" ", leftPad) + text + strings.Repeat(" ", rightPad) + " "
	default:
		return " " + text + strings.Repeat(" ", padding) + " "
	}
}
