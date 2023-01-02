package problem5

import (
	"bufio"
	"container/list"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func CraneSimulation(fileName string, multiPick bool) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var queues []list.List
	movesRegex, _ := regexp.Compile(`\d+`)
	for scanner.Scan() {
		line := scanner.Text()
		if queues == nil {
			queues = make([]list.List, 0, int(math.Ceil(float64(len(line))/4.0)))
		}
		// Initialize the queues according to the possible number of boxes
		if line == "" || strings.HasPrefix(line, " 1") {
			continue
		} else if strings.HasPrefix(line, "move") {
			queues = processInstructions(line, queues, movesRegex, multiPick)
		} else {
			queues = processBoxes(line, queues)
		}
	}
	str := ""
	for i := 0; i < len(queues); i++ {
		value := queues[i].Back()
		if value == nil {
			continue
		}
		str += value.Value.(string)
	}
	return str
}

func processBoxes(boxes string, queues []list.List) []list.List {
	current := 0
	colIndex := 0
	for i := 0; current+3 <= len(boxes); i++ {
		column := boxes[current : current+3]
		if len(queues) <= colIndex {
			queues = queues[:cap(queues)]
		}
		if !strings.Contains(column, " ") {
			queues[colIndex].PushFront(column[1:2])
		}
		current += 4
		colIndex++
	}
	return queues
}

func processInstructions(instruction string, queues []list.List, r *regexp.Regexp, multiPick bool) []list.List {
	details := r.FindAllString(instruction, len(instruction)+1)
	boxesToMove := getInt(details[0])
	from := getInt(details[1]) - 1
	to := getInt(details[2]) - 1
	// Process a single box a time
	if !multiPick || boxesToMove == 1 {
		for i := 0; i < boxesToMove; i++ {
			element := queues[from].Back()
			if element != nil {
				queues[to].PushBack(element.Value)
				queues[from].Remove(element)
			}
		}
	} else {
		tmp := list.New()
		for i := 0; i < boxesToMove; i++ {
			element := queues[from].Back()
			if element != nil {
				tmp.PushBack(element.Value)
				queues[from].Remove(element)
			}
		}
		for e := tmp.Back(); e != nil; e = e.Prev() {
			queues[to].PushBack(e.Value)
		}
	}

	return queues
}

func getInt(value string) int {
	intVal, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		log.Fatalf("Problem parsing %s", value)
	}
	return int(intVal)
}
