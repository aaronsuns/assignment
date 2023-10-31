package numrange

import (
	"testing"
)

func TestParseRange(t *testing.T) {
	// Test a valid range.
	rangeStr := "10-100"
	expectedRange := Range{Start: 10, End: 100}
	parsedRange, err := ParseRange(rangeStr)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if parsedRange != expectedRange {
		t.Errorf("ParseRange(%s) = %v, expected %v", rangeStr, parsedRange, expectedRange)
	}

	// Test an invalid range.
	rangeStr = "invalid-range"
	_, err = ParseRange(rangeStr)
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}

func TestParseRanges(t *testing.T) {
	// Test valid range strings.
	rangeStr := "10-100,200-300,400-500"
	expectedRanges := []Range{
		{Start: 10, End: 100},
		{Start: 200, End: 300},
		{Start: 400, End: 500},
	}
	parsedRanges, err := ParseRanges(rangeStr)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !compareRanges(parsedRanges, expectedRanges) {
		t.Errorf("ParseRanges(%s) = %v, expected %v", rangeStr, parsedRanges, expectedRanges)
	}

	// Test an invalid range string.
	rangeStr = "10-100,invalid-range,400-500"
	_, err = ParseRanges(rangeStr)
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}

func TestFormatRanges(t *testing.T) {
	ranges := []Range{
		{Start: 10, End: 100},
		{Start: 200, End: 300},
		{Start: 400, End: 500},
	}
	expectedFormatted := "10-100, 200-300, 400-500"
	formatted := FormatRanges(ranges)
	if formatted != expectedFormatted {
		t.Errorf("FormatRanges(%v) = %s, expected %s", ranges, formatted, expectedFormatted)
	}
}

func compareRanges(a, b []Range) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
