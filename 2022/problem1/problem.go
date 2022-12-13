package problem1

import (
	"bufio"
	"container/heap"
	"log"
	"os"
	"strconv"
)

func TopN(n int) int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	h := &IntHeap{}
	currentCount := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			heap.Push(h, int(currentCount))
			currentCount = 0
			continue
		}
		value, err := strconv.ParseInt(line, 10, 0)
		if err != nil {
			log.Fatalf("Problem parsing %d %s", value, err)
		}
		currentCount = value + currentCount
	}
	heap.Push(h, int(currentCount))

	sum := 0
	for i := 0; i < n; i++ {
		sum += heap.Pop(h).(int)
	}
	return sum
}
