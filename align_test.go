package atbl

import "testing"

func TestFormatCell(t *testing.T) {
    tests := []struct {
        name     string
        text     string
        width    int
        align    string
        expected string
    }{
        {
            name:     "Left alignment",
            text:     "test",
            width:    8,
            align:    Left,
            expected: " test     ",
        },
        {
            name:     "Right alignment",
            text:     "test",
            width:    8,
            align:    Right,
            expected: "     test ",
        },
        {
            name:     "Center alignment",
            text:     "test",
            width:    8,
            align:    Center,
            expected: "   test   ",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := formatCell(tt.text, tt.width, tt.align)
            if result != tt.expected {
                t.Errorf("Expected %q, got %q", tt.expected, result)
            }
        })
    }
}