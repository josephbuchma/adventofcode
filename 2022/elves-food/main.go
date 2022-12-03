package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const N = 3

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	curAcc := 0

	topn := make([]int, N)

	scanner := bufio.NewScanner(file)
	// Assuming that input ends with a newline
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			val, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			curAcc += val
			continue
		}
		for i := range topn {
			if topn[i] < curAcc {
				topn[i] = curAcc
				sort.Sort(sort.IntSlice(topn))
				break
			}
		}
		curAcc = 0
		continue
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, v := range topn {
		sum += v
	}
	fmt.Println(sum)
}
