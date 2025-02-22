package atbl

import (
	"testing"
)

func TestBasicBorder(t *testing.T) {
    widths := []int{3, 4, 5}
    
    tests := []struct {
        name     string
        border   BorderStyle
        expected string
    }{
        {
            name:   "BasicBorder Top",
            border: BasicBorder,
            expected: "+-----+------+-------+\n",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := tt.border.Top(widths)
            if result != tt.expected {
                t.Errorf("Expected %q, got %q", tt.expected, result)
            }
        })
    }
}

func TestUnicodeBorder(t *testing.T) {
    widths := []int{3, 4, 5}
    
    tests := []struct {
        name     string
        border   BorderStyle
        expected string
    }{
        {
            name:   "UnicodeBorder Top",
            border: UnicodeBorder,
            expected: "╔═════╦══════╦═══════╗\n",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := tt.border.Top(widths)
            if result != tt.expected {
                t.Errorf("Expected %q, got %q", tt.expected, result)
            }
        })
    }
}