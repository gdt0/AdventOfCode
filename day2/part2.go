package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	// "strconv"
	"time"
)

// Create a struct
type Game struct {
	Id    int
	red   int
	blue  int
	green int
}

var (
	SumOfPower *int
)

func main() {
	SumOfPower = new(int)
	start := time.Now()
	readFile()

	fmt.Println("SumOfPower: ", *SumOfPower)
	fmt.Println("time: ", time.Since(start))
}

func readFile() {
	fileName := "Input.txt"
	file, err := os.Open(fileName)

	defer file.Close()
	scanner := bufio.NewScanner(file)


	for scanner.Scan() {
		text := scanner.Text()
		getGameData(text)
		// fmt.Println("scanner.Text(): ", text)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func getGameData(line string) {
	idPos := 5
	game := new(Game)
	red := "red"
	strlen := len(line) - 1
	blue := "blue"
	green := "green"
	fmt.Println("red: ", game.red)
	for index, val := range line {
		if val == ':' {
			game.Id, _ = strconv.Atoi(line[idPos:index])
			fmt.Println(line[idPos:index])
		}
		if index < strlen-1 && strings.Contains(line[index:index+3], red) {
			temp, _ := strconv.Atoi(strings.TrimSpace(line[index-3 : index-1]))
				if game.red <temp{
					game.red = temp
					// fmt.Printf("\nnew red: %v", game.red)
				}
		}
		if index < strlen-2 && strings.Contains(line[index:index+4], blue) {
			temp, _ := strconv.Atoi(strings.TrimSpace(line[index-3 : index-1]))
				if game.blue <temp{
					game.blue = temp
					// fmt.Printf("\nnew blue: %v", game.blue)
				}
		}
		if index < strlen-3 && strings.Contains(line[index:index+5], green) {
			temp, _ := strconv.Atoi(strings.TrimSpace(line[index-3 : index-1]))
				if game.green <temp{
					game.green = temp
					// fmt.Printf("\nnew green: %v", game.green)
				}
		}

	}
	fmt.Println("game: ", game)
	*SumOfPower += game.red * game.blue * game.green
}
