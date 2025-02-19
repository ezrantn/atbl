package atbl

import "strings"

const (
	Left   = "left"
	Right  = "right"
	Center = "center"
)

type Table struct {
	headers    []string
	rows       [][]string
	alignments []string
	border     BorderStyle
	color      bool
}

// Creates a simple table with headers. Uses BasicBorder by default
func New(headers ...string) *Table {
	// Initialize all alignments to Left by default
	alignments := make([]string, len(headers))
	for i := range alignments {
		alignments[i] = Left
	}
	
	return &Table{
		headers:    headers,
		alignments: alignments,
		border:     BasicBorder,
	}
}

// Creates a table with right-aligned numeric columns
func NewNumeric(headers ...string) *Table {
	t := New(headers...)
	// Right-align the last column by default (usually contains numbers)
	if len(headers) > 0 {
		t.alignments[len(headers)-1] = Right
	}
	
	return t
}

// Simple functions with clear purposes
func (t *Table) SetBorder(border BorderStyle) {
	t.border = border
}

func (t *Table) SetAlignment(column int, alignment string) {
	if column < len(t.alignments) {
		t.alignments[column] = alignment
	}
}

func (t *Table) AddRow(cells ...string) {
	t.rows = append(t.rows, cells)
}

func (t *Table) EnableColor() {
	t.color = true
}

func (t *Table) Render() string {
	colWidths := calculateColumnWidths(t.headers, t.rows)
	var sb strings.Builder

	// Top Border
	sb.WriteString(t.border.Top(colWidths))

	// Print Headers
	sb.WriteString(t.border.Left)
	for i, header := range t.headers {
		sb.WriteString(formatCell(header, colWidths[i], t.alignments[i]) + t.border.Separator)
	}
	
	sb.WriteString("\n")

	// Middle Border (Separator between header and rows)
	sb.WriteString(t.border.Middle(colWidths))

	// Print Rows
	for _, row := range t.rows {
		sb.WriteString(t.border.Left)
		for i, cell := range row {
			sb.WriteString(formatCell(cell, colWidths[i], t.alignments[i]) + t.border.Separator)
		}
		sb.WriteString("\n")
	}

	// Bottom Border
	sb.WriteString(t.border.Bottom(colWidths))

	return sb.String()
}

func calculateColumnWidths(headers []string, rows [][]string) []int {
	widths := make([]int, len(headers))
	for i, header := range headers {
		widths[i] = len(header)
	}

	for _, row := range rows {
		for i, cell := range row {
			if len(cell) > widths[i] {
				widths[i] = len(cell)
			}
		}
	}

	return widths
}