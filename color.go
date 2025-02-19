package atbl

import "fmt"

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Reset  = "\033[0m"
)

func Colorize(text, color string) string {
	return fmt.Sprintf("%s%s%s", color, text, Reset)
}
