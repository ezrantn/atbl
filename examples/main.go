package main

import (
	"fmt"

	"github.com/ezrantn/atbl"
)

func main() {
	// Create a simple table
	table := atbl.New("#", "Name", "Phone", "Email", "Qty")

	// Add data
	table.AddRow("1", "Newton G. Goetz", "252-585-5166", "NewtonGGoetz@dayrep.com", "10")
	table.AddRow("2", "Rebecca R. Edney", "865-475-4171", "RebeccaREdney@armyspy.com", "12")

	fmt.Println(table.Render())
	fmt.Println(atbl.Colorize(table.Render(), atbl.Red))

	// Example with numeric table
	// numericTable := atbl.NewNumeric("ID", "Item", "Price")
	// numericTable.AddRow("1", "Widget", "9.99")
	// numericTable.AddRow("2", "Gadget", "24.99")

	// fmt.Println(numericTable.Render())
}
