package atbl

import (
	"testing"
	"strings"
)

func TestNewTable(t *testing.T) {
	headers := []string{"Name", "Age", "Occupation"}
	table := New(headers...)

	if len(table.headers) != len(headers) {
		t.Errorf("Expected %d headers, got %d", len(headers), len(table.headers))
	}

	for i, header := range headers {
		if table.headers[i] != header {
			t.Errorf("Expected header %s, got %s", header, table.headers[i])
		}
	}

	for _, alignment := range table.alignments {
		if alignment != Left {
			t.Errorf("Expected default alignment to be %s, got %s", Left, alignment)
		}
	}
}

func TestSetAlignment(t *testing.T) {
	headers := []string{"Name", "Age", "Occupation"}
	table := New(headers...)

	table.SetAlignment(1, Center)

	if table.alignments[1] != Center {
		t.Errorf("Expected alignment to be %s, got %s", Center, table.alignments[1])
	}
}

func TestAddRow(t *testing.T) {
	headers := []string{"Name", "Age", "Occupation"}
	table := New(headers...)

	row := []string{"John Doe", "30", "Engineer"}
	table.AddRow(row...)

	if len(table.rows) != 1 {
		t.Errorf("Expected 1 row, got %d", len(table.rows))
	}

	for i, cell := range table.rows[0] {
		if cell != row[i] {
			t.Errorf("Expected cell %s, got %s", row[i], cell)
		}
	}
}

func TestRender(t *testing.T) {
	headers := []string{"Name", "Age", "Occupation"}
	table := New(headers...)

	table.AddRow("John Doe", "30", "Engineer")
	table.AddRow("Jane Doe", "25", "Designer")

	expected := `
+----------+-----+------------+
| Name     | Age | Occupation |
+----------+-----+------------+
| John Doe | 30  | Engineer   |
| Jane Doe | 25  | Designer   |
+----------+-----+------------+
`

	// Remove the first newline character for comparison
	expected = strings.TrimSpace(expected)
	actual := strings.TrimSpace(table.Render())

	if actual != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, actual)
	}
}