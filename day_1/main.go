package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getDist(l int, r int) int {
	if l == r {
		return 0
	} else if r > l {
		return r - l
	} else {
		return l - r
	}
}

func main() {
	// Read input.txt into memory
	f, err := os.Open("./input.txt")
	check(err)

	defer f.Close()
	reader := bufio.NewReader(f)
	left := []int{}
	right := []int{}
	simScore := make(map[int]int)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		elems := strings.Fields(string(line))
		asInt, err := strconv.Atoi(elems[0])
		check(err)
		left = append(left, asInt)
		simScore[asInt] = 0

		asIntR, err := strconv.Atoi(elems[1])
		check(err)
		right = append(right, asIntR)
	}
	f.Close()


	sort.Ints(left)
	sort.Ints(right)

	dists := 0
	for i := 0; i < len(left); i++ {
		n := getDist(left[i], right[i])
		dists += n

		_, ok := simScore[right[i]]
		if ok {
			simScore[right[i]] += right[i]
		}
	}


	// Calc similarity score
	res := 0
	for i:= 0; i < len(left); i++ {
		res += simScore[left[i]]
	}
	fmt.Println(res)
}