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

func sortAndMergeRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return ranges
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	mergedRanges := []Range{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		currentRange := &mergedRanges[len(mergedRanges)-1]
		if currentRange.End >= ranges[i].Start {
			if currentRange.End < ranges[i].End {
				currentRange.End = ranges[i].End
			}
		} else {
			mergedRanges = append(mergedRanges, ranges[i])
		}
	}

	return mergedRanges
}

// SortAndMerge sorts the ranges and merges any overlapping ranges.
func (ranges *Ranges) SortAndMerge() {
	*ranges = sortAndMergeRanges(*ranges)
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
	result.SortAndMerge()
	DebugPrintf("includes before/after SortAndMerge: %v , %v", includes, result)

	excludesSortAndMerged := sortAndMergeRanges(excludes)
	DebugPrintf("excludes before/after SortAndMerge: %v , %v", excludes, excludesSortAndMerged)

	for _, exclude := range excludesSortAndMerged {
		if exclude.Start > exclude.End {
			printInvalidRange("exclude", exclude.Start, exclude.End)
			continue
		}
		result.ExcludeRange(exclude.Start, exclude.End)
	}
	DebugPrintf("Result after excludes: %v", result)

	result.SortAndMerge()
	DebugPrintf("Result: %v", result)
	return result
}
