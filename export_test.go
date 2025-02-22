package atbl

import (
	"testing"
)

// func TestToMarkdown(t *testing.T) {
// 	table := New("Name", "Age")
// 	table.AddRow("John", "30")
	
// 	expected := "```\n" +
// 		"+------+-----+\n" +
// 		"| Name | Age |\n" +
// 		"+------+-----+\n" +
// 		"| John | 30  |\n" +
// 		"+------+-----+\n" +
// 		"```\n"
	
// 	result := table.ToMarkdown()
	
// 	if result != expected {
// 		t.Errorf("Expected:\n%q\nGot:\n%q", expected, result)
// 	}
// }

func TestToCSV(t *testing.T) {
	table := New("Name", "Age")
	table.AddRow("John", "30")
	
	expected := "Name,Age\nJohn,30\n"
	result := table.ToCSV()
	
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}