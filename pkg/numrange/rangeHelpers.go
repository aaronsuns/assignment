package numrange

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseRange(rangeStr string) (Range, error) {
	parts := strings.Split(rangeStr, "-")
	if len(parts) != 2 {
		return Range{}, fmt.Errorf("invalid range format: %s", rangeStr)
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return Range{}, fmt.Errorf("invalid start value in range: %s", rangeStr)
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return Range{}, fmt.Errorf("invalid end value in range: %s", rangeStr)
	}

	return Range{Start: start, End: end}, nil
}

// parseRanges parses a comma-separated list of range strings and returns a slice of Range objects.
func ParseRanges(input string) ([]Range, error) {
	var ranges []Range
	rangeStrings := strings.Split(input, ",")
	for _, rangeStr := range rangeStrings {
		r, err := ParseRange(rangeStr)
		if err != nil {
			return nil, err
		}
		ranges = append(ranges, r)
	}
	return ranges, nil
}

// formatRanges formats a slice of Range objects as a string.
func FormatRanges(ranges []Range) string {
	var formattedRanges []string
	for _, r := range ranges {
		formattedRanges = append(formattedRanges, fmt.Sprintf("%d-%d", r.Start, r.End))
	}
	return strings.Join(formattedRanges, ", ")
}
