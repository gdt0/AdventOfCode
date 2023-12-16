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
	GameIdCount *int
)

func main() {
	GameIdCount = new(int)
	start := time.Now()
	readFile()

	fmt.Println("GameIdCount: ", *GameIdCount)
	fmt.Println("time: ", time.Since(start))
}

func readFile() {
	fileName := "Input.txt"
	file, err := os.Open(fileName)

	defer file.Close()
	scanner := bufio.NewScanner(file)

	i := 5
	for scanner.Scan() {
		text := scanner.Text()
		if i>0{
			getGameData(text)
			fmt.Println("scanner.Text(): ", text)
			i--
		}

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
	for index, val := range line {
		if val == ':' {
			game.Id, _ = strconv.Atoi(line[idPos:index])
			fmt.Println(line[idPos:index])
		}
		if index < strlen-1 && strings.Contains(line[index:index+3], red) {
			temp, _ := strconv.Atoi(line[index-3 : index-1])
				if game.red <temp{
					game.red = temp
				}
		}
		if index < strlen-2 && strings.Contains(line[index:index+4], blue) {
			temp, _ := strconv.Atoi(line[index-3 : index-1])
				if game.blue <temp{
					game.blue = temp
				}
		}
		if index < strlen-3 && strings.Contains(line[index:index+5], green) {
			temp, _ := strconv.Atoi(line[index-3 : index-1])
				if game.green <temp{
					game.green = temp
				}
		}

	}
	fmt.Println("game: ", game)
}
