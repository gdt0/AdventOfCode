package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	totalCount *int
	spelled    = []string{
		"zero",
		"one", "two",
		"three", "four",
		"five", "six",
		"seven", "eight",
		"nine"}
)

func GetTotalNum() {
	totalCount = new(int)
	// fmt.Println("Reading Input File")
	fileName := "Input1.txt"
	file, err := os.Open(fileName)

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		text := scanner.Text()

		lineCount := GetIntArray(text)

		*totalCount = *totalCount + lineCount
		// fmt.Println("lineCount:= ", lineCount, " totalCount:= ", *totalCount)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func GetIntArray(line string) int {
	var first int = -1
	var last int = -1

	for index, _ := range line {
		strlen := len(line) - 1
		newStr := line[:index]
		lastStr := line[strlen-index:]

		if first == -1 {
			for key, val := range spelled {
				if strings.Contains(newStr, val) {
					first = key
				}
				if first == -1 {
					first = FetchInt(line[index])
				}
			}
		}
		if last == -1 {
			for key, val := range spelled {
				if strings.Contains(lastStr, val) {
					last = key
				}
				if last == -1 {
					last = FetchInt(line[strlen-index])
				}
			}
		}
	}
	return first*10 + last
}

func FetchInt(char byte) int {
	intval, err := strconv.Atoi(string(char))
	if err == nil {
		return intval
	}
	return -1
}

func main() {
	start := time.Now()
	totalCount = new(int)
	GetTotalNum()
	// FetchInt("AB9CD1")
	// GetIntArray("onetwo9", []int{})
	fmt.Println("\nFinal Total Count: ", *totalCount)
	fmt.Println("Time taken: ", time.Since(start))

}
