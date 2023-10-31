package numrange

import (
	"fmt"
	"sort"
)

// Range represents a numeric range.
type Range struct {
	Start int
	End   int
}

// Ranges represents a slice of numeric ranges.
type Ranges []Range

// IncludeRange adds a numeric range to the Ranges slice.
func (ranges *Ranges) IncludeRange(start, end int) {
	*ranges = append(*ranges, Range{Start: start, End: end})
}

// ExcludeRange removes a numeric range from the Ranges slice.
func (ranges *Ranges) ExcludeRange(start, end int) {
	exclude := Range{Start: start, End: end}
	newRanges := Ranges{}
	DebugPrintf("Excluding Range: %d-%d", start, end)
	for _, r := range *ranges {
		switch {
		case r.Start >= exclude.End || r.End <= exclude.Start:
			// No overlap, keep the original range.
			DebugPrintf("No overlap: %d-%d", r.Start, r.End)
			newRanges = append(newRanges, r)
		case r.Start < exclude.Start && r.End > exclude.End:
			// The exclude range is fully contained within the current range.
			DebugPrintf("Fully contained: %d-%d", r.Start, r.End)
			newRanges = append(newRanges, Range{Start: r.Start, End: exclude.Start - 1})
			newRanges = append(newRanges, Range{Start: exclude.End + 1, End: r.End})
		case r.Start < exclude.Start:
			// The current range partially overlaps with the left side of the exclude range.
			DebugPrintf("Left overlap: %d-%d", r.Start, r.End)
			newRanges = append(newRanges, Range{Start: r.Start, End: exclude.Start - 1})
		case r.End > exclude.End:
			// The current range partially overlaps with the right side of the exclude range.
			DebugPrintf("Right overlap: %d-%d", r.Start, r.End)
			newRanges = append(newRanges, Range{Start: exclude.End + 1, End: r.End})
		}
	}
	*ranges = newRanges
}

// SortAndMerge sorts the ranges and merges any overlapping ranges.
func (ranges *Ranges) SortAndMerge() {
	sort.Slice(*ranges, func(i, j int) bool {
		return (*ranges)[i].Start < (*ranges)[j].Start
	})
	DebugPrintf("After Sort: %v", ranges)
	mergedRanges := Ranges{}
	currentRange := (*ranges)[0]

	for i := 1; i < len(*ranges); i++ {
		if currentRange.End >= (*ranges)[i].Start {
			// Merge the two overlapping ranges.
			currentRange.End = (*ranges)[i].End
		} else {
			// Ranges don't overlap; add the current range and update it.
			mergedRanges = append(mergedRanges, currentRange)
			currentRange = (*ranges)[i]
		}
	}

	mergedRanges = append(mergedRanges, currentRange)
	*ranges = mergedRanges
}

func printInvalidRange(rangeType string, start, end int) {
	fmt.Printf("WARN: Invalid %s range: %d-%d (start is greater than end)", rangeType, start, end)
}
func ProcessNumberRanges(includes, excludes []Range) Ranges {
	result := Ranges{}

	for _, include := range includes {
		if include.Start > include.End {
			printInvalidRange("include", include.Start, include.End)
			continue
		}
		result.IncludeRange(include.Start, include.End)
	}

	for _, exclude := range excludes {
		if exclude.Start > exclude.End {
			printInvalidRange("exclude", exclude.Start, exclude.End)
			continue
		}
		result.ExcludeRange(exclude.Start, exclude.End)
	}

	result.SortAndMerge()
	DebugPrintf("Result: %v", result)
	return result
}
