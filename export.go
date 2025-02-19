package atbl

import "strings"

func (t *Table) ToMarkdown() string {
	var sb strings.Builder
	sb.WriteString("| " + strings.Join(t.headers, " | ") + " |\n")
	sb.WriteString("|" + strings.Repeat("---|", len(t.headers)) + "\n")
	for _, row := range t.rows {
		sb.WriteString("| " + strings.Join(row, " | ") + " |\n")
	}
	return sb.String()
}

func (t *Table) ToCSV() string {
	var sb strings.Builder
	sb.WriteString(strings.Join(t.headers, ",") + "\n")
	for _, row := range t.rows {
		sb.WriteString(strings.Join(row, ",") + "\n")
	}
	return sb.String()
}
