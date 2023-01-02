package problem3

import (
	"bufio"
	"log"
	"os"
)

func FindReorganization(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	priority := 0
	for scanner.Scan() {
		line := scanner.Text()
		r := []rune(line)
		n := len(r)
		left := r[0 : n/2]
		right := r[n/2 : n]
		set := map[rune]bool{}
		for i := 0; i < n/2; i++ {
			set[left[i]] = true
		}
		found := map[rune]bool{}
		for i := 0; i < n/2; i++ {
			if _, exists := set[right[i]]; exists {
				if _, counted := found[right[i]]; !counted {
					priority += runeValue(right[i])
					found[right[i]] = true
				}
			}
		}

	}
	return priority
}

func FindBadge(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	priority := 0
	count := 0
	set := map[rune]int{}
	for scanner.Scan() {
		line := scanner.Text()
		r := []rune(line)
		count++
		found := map[rune]bool{}
		for i := 0; i < len(r); i++ {
			if _, counted := found[r[i]]; !counted {
				found[r[i]] = true
			}
		}
		for key, _ := range found {
			set[key]++
		}

		if count == 3 {
			for key, value := range set {
				if value == 3 {
					priority += runeValue(key)
					break
				}
			}
			set = map[rune]int{}
			count = 0
		}
	}
	return priority
}

func runeValue(val rune) int {
	if val >= 65 && val <= 90 {
		return int(val - 'A' + 27)
	}

	return int(val - 'a' + 1)
}
