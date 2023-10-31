package numrange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Example 1:
// Includes: 10-100
// Excludes: 20-30
// Output should be: 10-19, 31-100
func TestExample1(t *testing.T) {
	includeRanges := []Range{
		{Start: 10, End: 100},
	}

	excludeRanges := []Range{
		{Start: 20, End: 30},
	}

	result := ProcessNumberRanges(includeRanges, excludeRanges)

	expected := Ranges{
		{Start: 10, End: 19},
		{Start: 31, End: 100},
	}

	assert.Equal(t, expected, result)
}

// Example 2:
// Includes: 50-5000, 10-100
// Excludes: (none)
// Output: 10-5000
func TestExample2(t *testing.T) {
	includeRanges := []Range{
		{Start: 50, End: 5000},
		{Start: 10, End: 100},
	}

	excludeRanges := []Range{}

	result := ProcessNumberRanges(includeRanges, excludeRanges)

	expected := Ranges{
		{Start: 10, End: 5000},
	}

	assert.Equal(t, expected, result)
}

// Example 3:
// Includes: 200-300, 50-150
// Excludes: 95-205
// Output: 50-94, 206-300
func TestExample3(t *testing.T) {
	includeRanges := []Range{
		{Start: 200, End: 300},
		{Start: 50, End: 150},
	}

	excludeRanges := []Range{
		{Start: 95, End: 205},
	}

	result := ProcessNumberRanges(includeRanges, excludeRanges)

	expected := Ranges{
		{Start: 50, End: 94},
		{Start: 206, End: 300},
	}

	assert.Equal(t, expected, result)
}

// Example 4:
// Includes: 200-300, 10-100, 400-500
// Excludes: 410-420, 95-205, 100-150
// Output: 10-94, 206-300, 400-409, 421-500
func TestExample4(t *testing.T) {
	includeRanges := []Range{
		{Start: 200, End: 300},
		{Start: 10, End: 100},
		{Start: 400, End: 500},
	}

	excludeRanges := []Range{
		{Start: 410, End: 420},
		{Start: 95, End: 205},
		{Start: 100, End: 150},
	}

	result := ProcessNumberRanges(includeRanges, excludeRanges)

	expected := Ranges{
		{Start: 10, End: 94},
		{Start: 206, End: 300},
		{Start: 400, End: 409},
		{Start: 421, End: 500},
	}

	assert.Equal(t, expected, result)
}

// Invalid range '200-3' is skipped
func TestExample5(t *testing.T) {
	includeRanges := []Range{
		{Start: 200, End: 3},
		{Start: 10, End: 100},
		{Start: 400, End: 500},
	}

	excludeRanges := []Range{
		{Start: 410, End: 420},
		{Start: 95, End: 205},
		{Start: 100, End: 150},
	}

	result := ProcessNumberRanges(includeRanges, excludeRanges)

	expected := Ranges{
		{Start: 10, End: 94},
		{Start: 400, End: 409},
		{Start: 421, End: 500},
	}

	assert.Equal(t, expected, result)
}
