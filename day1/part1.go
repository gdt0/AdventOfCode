package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	totalCount *int
)

func GetTotalNum() {
	totalCount = new(int)
	fileName := "Input1.txt"
	file, err := os.Open(fileName)

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		text := scanner.Text()
		lineCount := FetchInt(text)

		*totalCount = *totalCount + lineCount
	}
	if err != nil {
		log.Fatal(err)
	}
}

func FetchInt(line string) int {

	var strlen int
	// var stopFirst bool = false
	// var stopLast bool = false
	var firstInt int = -1
	var lastInt int = -1
	var count int
	
	lineRune := []rune(line)
	// fmt.Println(lineRune)
	strlen = len(line) - 1
	for i := 0; i <= strlen; i++ {
		if firstInt == -1 {
			first := lineRune[i]
			val, err := strconv.Atoi(string(first)) 
			if err == nil {
				firstInt = val
			}
		}
		if lastInt == -1 {
			last := lineRune[strlen-i]
			val, err := strconv.Atoi(string(last))
			if err == nil {
				lastInt = val
			}
		}
	}
	count = firstInt*10 + lastInt
	return count
}

func main() {
	start := time.Now()

	GetTotalNum()
	// FetchInt("AB9CD1")
	fmt.Println("\nFinal Total Count: ", *totalCount)
	fmt.Println("Time taken: ", time.Since(start))
}
