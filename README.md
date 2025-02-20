# [@ezrantn/atbl](https://github.com/ezrantn/atbl)

ATBL or ASCII Table is a simple and flexible table generator that supports both ASCII and Unicode borders, with customizable alignments, export options and colors.

## Features

- ASCII and Unicode border styles
- Customizable column alignments (Left, Right, Center)
- Export to Markdown and CSV formats
- Color support for terminal output
- Simple API with sensible defaults

## Installation

```bash
go get github.com/ezrantn/atbl
```

## Usage

To create a table, use the `New` function to define the headers, and then add rows of data using the `AddRow` method. Finally, render the table using the `Render` method.

**Example**

```go
// Create a simple table
table := atbl.New("#", "Name", "Phone", "Email", "Qty")

// Add data
table.AddRow("1", "Newton G. Goetz", "252-585-5166", "NewtonGGoetz@dayrep.com", "10")
table.AddRow("2", "Rebecca R. Edney", "865-475-4171", "RebeccaREdney@armyspy.com", "12")

fmt.Println(table.Render())
```

**Output**

```bash
+ --- + ------------------ + -------------- + --------------------------- + ----- +
|  #  |  Name              |  Phone         |  Email                      |  Qty  | 
+ --- + ------------------ + -------------- + --------------------------- + ----- +
|  1  |  Newton G. Goetz   |  252-585-5166  |  NewtonGGoetz@dayrep.com    |  10   | 
|  2  |  Rebecca R. Edney  |  865-475-4171  |  RebeccaREdney@armyspy.com  |  12   | 
+ --- + ------------------ + -------------- + --------------------------- + ----- +
```

