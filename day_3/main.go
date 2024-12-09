package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"os"
)

func main() {
	// Read input.txt into memory
	// f, err := os.ReadFile("./test.txt")
	f, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(f)
	input = "do()" + input

	// Find do() || dont()
	reDo := regexp.MustCompile(`don't\(\)|do\(\)`)
	cmds := reDo.FindAllStringIndex(input, - 1)

	// fmt.Println("CMDS:", cmds)
	sections := make(map[string][]string)

	// Split input into sections
	for i := 0; i < len(cmds) - 1; i++ {
		// fmt.Println(input[cmds[i][0]:cmds[i][1]], "|||", input[cmds[i][1]:cmds[i + 1][0]])
		key := input[cmds[i][0]:cmds[i][1]]
		sections[key] = append(sections[key], input[cmds[i][1]:cmds[i + 1][0]])
	}
	// fmt.Println("|||", input[cmds[len(cmds) -1][1]:])
	key := input[cmds[len(cmds) -1][0]:cmds[len(cmds) -1][1]]
	sections[key] = append(sections[key],input[cmds[len(cmds) -1][1]:])

	// fmt.Println(sections)
	// Find operations in each section
	re := regexp.MustCompile(`mul\([-+]?\d{1,3},[-+]?\d{1,3}\)`)
	res := 0
	for _, matches := range(sections["do()"]) {
		for _, match := range re.FindAllStringIndex(matches, -1) {
			// Execute and sum
			s := strings.Split(matches[match[0] + 4:match[1] - 1], ",")
			fmt.Println("MATCH:", s)
	
			uno, _ := strconv.Atoi(s[0])
			dos,_ := strconv.Atoi(s[1])
			res += uno * dos
		}

	}

	// Output matches
	fmt.Println("Matches found:")
	fmt.Println(res)
}