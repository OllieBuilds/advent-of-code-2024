package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseReport(line []byte) []int {
	elems := strings.Fields(string(line))
	report := []int{}
	for i :=0; i < len(elems); i++ {
		n, err := strconv.Atoi(elems[i])
		check(err)
		report = append(report, n)
	}

	return report
}

func checkReportSafe(report []int) bool {
	increasing := report[1] > report[0]
	unsafeSeen := false
	for i := 1; i < len(report); i++ {
		if increasing {
			if report[i] <= report[i - 1] || report[i] - report[i - 1] > 3 {
				if unsafeSeen {
					return false
				}
				unsafeSeen = true
			}
		} else {
			if report[i] >=  report[i - 1] || report[i - 1] - report[i] > 3 {
				if unsafeSeen {
					return false
				}
				unsafeSeen = true
			}
		}
	}
	return true
}

func main() {
	// Read input.txt into memory
	f, err := os.Open("./input.txt")
	check(err)

	defer f.Close()
	reader := bufio.NewReader(f)

	safeCount := 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		report := parseReport(line)
		if checkReportSafe(report) {
			safeCount += 1
		}
	}

	fmt.Println(safeCount)
}