package atbl

import (
	"os"
	"strings"
)

func (t *Table) ToMarkdown() string {
	var sb strings.Builder
	// Add proper spacing in markdown table
	headers := make([]string, len(t.headers))
	for i, h := range t.headers {
		headers[i] = strings.TrimSpace(h)
	}
	sb.WriteString("| " + strings.Join(headers, " | ") + " |\n")
	
	// Add alignment indicators
	alignRow := make([]string, len(t.headers))
	for i, align := range t.alignments {
		switch align {
		case Right:
			alignRow[i] = "---:"
		case Center:
			alignRow[i] = ":---:"
		default:
			alignRow[i] = "---"
		}
	}
	sb.WriteString("| " + strings.Join(alignRow, " | ") + " |\n")
	
	// Add data rows with proper spacing
	for _, row := range t.rows {
		cells := make([]string, len(row))
		for i, cell := range row {
			cells[i] = strings.TrimSpace(cell)
		}
		sb.WriteString("| " + strings.Join(cells, " | ") + " |\n")
	}
	return sb.String()
}

func (t *Table) ToCSV() string {
	var sb strings.Builder
	// Properly handle CSV escaping
	sb.WriteString(formatCSVRow(t.headers) + "\n")
	for _, row := range t.rows {
		sb.WriteString(formatCSVRow(row) + "\n")
	}
	return sb.String()
}

// New export to file methods
func (t *Table) ExportToMarkdown(filename string) error {
	content := t.ToMarkdown()
	return os.WriteFile(filename, []byte(content), 0644)
}

func (t *Table) ExportToCSV(filename string) error {
	content := t.ToCSV()
	return os.WriteFile(filename, []byte(content), 0644)
}

// Helper function to properly format CSV rows
func formatCSVRow(row []string) string {
	// Escape fields that contain commas, quotes, or newlines
	escapedFields := make([]string, len(row))
	for i, field := range row {
		if strings.ContainsAny(field, ",\"\n") {
			escapedFields[i] = `"` + strings.ReplaceAll(field, `"`, `""`) + `"`
		} else {
			escapedFields[i] = field
		}
	}
	return strings.Join(escapedFields, ",")
}
