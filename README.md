# [@ezrantn/atbl](https://github.com/ezrantn/atbl)

> [!WARNING]
> Work in progress!

ATBL or ASCII Table is a simple and flexible table generator that supports both ASCII and Unicode borders, with customizable alignments, export options and colors.

## Features

- ASCII and Unicode border styles
- Customizable column alignments (Left, Right, Center)
- Export to Markdown and CSV formats
- Color support for terminal output
- Simple API with sensible defaults
- Dependency free

## Installation

```bash
go get github.com/ezrantn/atbl
```

## Usage

### Quick Start

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

### Customizing Column Alignment

By default, all columns are left-aligned. You can change the alignment of specific columns using the `SetAlignment` method.

**Example**

```go
// Create a table with headers
table := atbl.New("#", "Name", "Phone", "Email", "Qty")

// Set the alignment of the "Qty" column to right-aligned
table.SetAlignment(4, atbl.Right)

// Add rows of data
table.AddRow("1", "Newton G. Goetz", "252-585-5166", "NewtonGGoetz@dayrep.com", "10")
table.AddRow("2", "Rebecca R. Edney", "865-475-4171", "RebeccaREdney@armyspy.com", "12")

// Render and print the table
fmt.Println(table.Render())
```

**Output**

```bash
+ --- + ------------------ + -------------- + --------------------------- + ----- +
|  #  |  Name              |  Phone         |  Email                      |  Qty  | 
+ --- + ------------------ + -------------- + --------------------------- + ----- +
|  1  |  Newton G. Goetz   |  252-585-5166  |  NewtonGGoetz@dayrep.com    |    10 | 
|  2  |  Rebecca R. Edney  |  865-475-4171  |  RebeccaREdney@armyspy.com  |    12 | 
+ --- + ------------------ + -------------- + --------------------------- + ----- +
```

### Creating Numeric Tables

For tables with numeric data, you can use the `NewNumeric` function, which automatically right-aligns the last column.

**Example**

```go
// Create a numeric table with headers
table := atbl.NewNumeric("ID", "Name", "Age", "Salary")

// Add rows of data
table.AddRow("1", "John Doe", "30", "100000")
table.AddRow("2", "Jane Smith", "25", "80000")

// Render and print the table
fmt.Println(table.Render())
```

**Output**

```bash
+ ---- + ----------- + ----- + -------- +
|  ID  |  Name       |  Age  |  Salary  | 
+ ---- + ----------- + ----- + -------- +
|  1   |  John Doe   |   30  |  100000  | 
|  2   |  Jane Smith |   25  |   80000  | 
+ ---- + ----------- + ----- + -------- +
```

### Customizing Table Borders

You can change the border style of the table using the `SetBorder` method. The package provides predefined border styles like `BasicBorder` and `UnicodeBorder`.

**Basic Border Example**

```go
table.SetBorder(atbl.BasicBorder)
```

**Basic Border Output**

```bash
+ --- + ------------------ + -------------- + --------------------------- + ----- +
|  #  |  Name              |  Phone         |  Email                      |  Qty  | 
+ --- + ------------------ + -------------- + --------------------------- + ----- +
|  1  |  Newton G. Goetz   |  252-585-5166  |  NewtonGGoetz@dayrep.com    |  10   | 
|  2  |  Rebecca R. Edney  |  865-475-4171  |  RebeccaREdney@armyspy.com  |  12   | 
+ --- + ------------------ + -------------- + --------------------------- + ----- +
```

**Unicode Border Example**

```go
table.SetBorder(atbl.UnicodeBorder)
```

**Unicode Border Output**

```bash
╔ ═══ ╦ ══════════════════ ╦ ══════════════ ╦ ═══════════════════════════ ╦ ═════ ╗
║  #  ║  Name              ║  Phone         ║  Email                      ║  Qty  ║ 
╠ ═══ ╬ ══════════════════ ╬ ══════════════ ╬ ═══════════════════════════ ╬ ═════ ╣
║  1  ║  Newton G. Goetz   ║  252-585-5166  ║  NewtonGGoetz@dayrep.com    ║  10   ║ 
║  2  ║  Rebecca R. Edney  ║  865-475-4171  ║  RebeccaREdney@armyspy.com  ║  12   ║ 
╚ ═══ ╩ ══════════════════ ╩ ══════════════ ╩ ═══════════════════════════ ╩ ═════ ╝
```

### Enabling Color

If you want to add color to your table, use the `Colorize` method. This will enable ANSI color codes in the rendered output.

**Available Colors:**

- Text: `Black`, `Red`, `Green`, `Yellow`, `Blue`, `Magenta`, `Cyan`, `White`
- Bright Text: `BrightBlack`, `BrightRed`, `BrightGreen`, `BrightYellow`, `BrightBlue`, `BrightMagenta`, `BrightCyan`, `BrightWhite`
- Background: `BgBlack`, `BgRed`, `BgGreen`, `BgYellow`, `BgBlue`, `BgMagenta`, `BgCyan`, `BgWhite`
- Bright Background: `BgBrightBlack`, `BgBrightRed`, `BgBrightGreen`, `BgBrightYellow`, `BgBrightBlue`, `BgBrightMagenta`, `BgBrightCyan`, `BgBrightWhite`

**Example**

```go
fmt.Println(atbl.Colorize(table.Render(), atbl.Red))
```

```go
fmt.Println(atbl.Colorize(table.Render(), atbl.BrightRed))
```

```go
fmt.Println(atbl.Colorize(table.Render(), atbl.BgRed))
```

```go
fmt.Println(atbl.Colorize(table.Render(), atbl.BgBrightRed))
```

## License

This tool is open-source and available under the [MIT License](https://github.com/ezrantn/atbl/blob/main/LICENSE).

## Contributing

Contributions are welcome! Please feel free to submit a pull request. For major changes, please open an issue first to discuss what you would like to change.