package atbl

import "fmt"

func formatCell(text string, width int, align string) string {
	switch align {
	case Left:
		return fmt.Sprintf(" %-*s ", width, text)
	case Right:
		return fmt.Sprintf(" %*s ", width, text)
	case Center:
		padding := (width - len(text)) / 2
		return fmt.Sprintf(" %*s%-*s ", padding, text, padding, "")
	default:
		return fmt.Sprintf(" %-*s ", width, text)
	}
}
