package atbl

import "strings"

type BorderStyle struct {
	Top, Middle, Bottom    func([]int) string
	Left, Right, Separator string
}

// Basic ASCII border
var BasicBorder = BorderStyle{
	Top: func(widths []int) string {
		return "+" + strings.Join(generateLine(widths, "-"), "+") + "+\n"
	},
	Middle: func(widths []int) string {
		return "+" + strings.Join(generateLine(widths, "-"), "+") + "+\n"
	},
	Bottom: func(widths []int) string {
		return "+" + strings.Join(generateLine(widths, "-"), "+") + "+\n"
	},
	Left:      "|",
	Right:     "|",
	Separator: "|",
}

// Unicode Border
var UnicodeBorder = BorderStyle{
	Top: func(widths []int) string {
		return "╔" + strings.Join(generateLine(widths, "═"), "╦") + "╗\n"
	},
	Middle: func(widths []int) string {
		return "╠" + strings.Join(generateLine(widths, "═"), "╬") + "╣\n"
	},
	Bottom: func(widths []int) string {
		return "╚" + strings.Join(generateLine(widths, "═"), "╩") + "╝\n"
	},
	Left:      "║",
	Right:     "║",
	Separator: "║",
}

func generateLine(widths []int, char string) []string {
	var segments []string
	for _, width := range widths {
		segments = append(segments, strings.Repeat(char, width+2))
	}
	return segments
}
