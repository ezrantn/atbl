package main

import (
	"fmt"
	"log"
	"github.com/ezrantn/atbl"
)

func main() {
	// Create a simple table
	table := atbl.New("#", "Name", "Phone", "Email", "Qty")

	// Add data
	table.AddRow("1", "Newton G. Goetz", "252-585-5166", "NewtonGGoetz@dayrep.com", "10")
	table.AddRow("2", "Rebecca R. Edney", "865-475-4171", "RebeccaREdney@armyspy.com", "12")

	fmt.Println(atbl.Colorize(table.Render(), atbl.Red))

	if err := table.ExportToMarkdown("output.md"); err != nil {
		log.Fatal("Failed to export to Markdown:", err)
	}

	fmt.Println("Table exported to output.md")

	if err := table.ExportToCSV("output.csv"); err != nil {
		log.Fatal("Failed to export to CSV:", err)
	}
	fmt.Println("Table exported to output.csv")
}
