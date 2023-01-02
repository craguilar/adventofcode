package problem4

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func FindInclusiveRanges(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	fullyContained := 0
	for scanner.Scan() {
		ranges := strings.Split(scanner.Text(), ",")
		range1 := stringToIntArray(strings.Split(ranges[0], "-"))
		range2 := stringToIntArray(strings.Split(ranges[1], "-"))
		if isFullyContainedRange(range1, range2) || isFullyContainedRange(range2, range1) {
			fullyContained++
		}
	}
	return fullyContained
}

func FindUnionRanges(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	overlap := 0
	for scanner.Scan() {
		ranges := strings.Split(scanner.Text(), ",")
		range1 := stringToIntArray(strings.Split(ranges[0], "-"))
		range2 := stringToIntArray(strings.Split(ranges[1], "-"))
		if overlapRange(range1, range2) || overlapRange(range2, range1) {
			overlap++
		}
	}
	return overlap
}

func overlapRange(range1, range2 []int) bool {
	return (range2[0] >= range1[0] && range2[0] <= range1[1]) || (range2[1] >= range1[0] && range2[1] <= range1[1])
}

// Does range1 contains range2?
func isFullyContainedRange(range1, range2 []int) bool {
	return range2[0] >= range1[0] && range2[1] <= range1[1]
}

func stringToIntArray(arr []string) []int {
	vals := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		vals[i] = getInt(arr[i])
	}
	return vals
}

func getInt(value string) int {
	intVal, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		log.Fatalf("Problem parsing %s", value)
	}
	return int(intVal)
}
