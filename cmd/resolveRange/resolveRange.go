package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aaronsuns/assignment/pkg/numrange"
)

func readInputFromStdin() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	var includeStr, excludeStr string
	flag.StringVar(&includeStr, "includes", "", "Include ranges (e.g., '200-300,10-100,400-500')")
	flag.StringVar(&excludeStr, "excludes", "", "Exclude ranges (e.g., '410-420,95-205,100-150')")
	flag.Parse()

	// Check if the user provided values for flags, and if not, prompt for input.
	if includeStr == "" {
		fmt.Print("Enter include ranges: ")
		includeStr = readInputFromStdin()
	}

	if excludeStr == "" {
		fmt.Print("Enter exclude ranges: ")
		excludeStr = readInputFromStdin()
	}

	includeRanges, err := numrange.ParseRanges(includeStr)
	if err != nil {
		log.Fatalf("Failed to parse include ranges: %v", err)
	}
	excludeRanges, err := numrange.ParseRanges(excludeStr)
	if err != nil {
		log.Fatalf("Failed to parse exclude ranges: %v", err)
	}

	outputRanges := numrange.ProcessNumberRanges(includeRanges, excludeRanges)
	fmt.Println("Output:", numrange.FormatRanges(outputRanges))
}
