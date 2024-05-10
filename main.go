package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"strings"
	"time"
)

func selectSort(arr []int, speed int) {
	drawLines := *draw(arr)
	importLines := []string{}
	var maxLenStr float64
	fmt.Println(arr)

	fmt.Println("Randomly generated array:")
	fmt.Println()
	for l := 0; l < len(arr); l++ {
		fmt.Println(drawLines[l])
		importLines = append(importLines, drawLines[l])
		maxLenStr = math.Max(float64(len(importLines[l])), maxLenStr)
	}

	fmt.Print("Press enter to sort: ")
	waitForUser()
	waitForUser()
	fmt.Print("\033[A")
	fmt.Printf("\r%s", "                       ")

	for j := 0; j < len(arr); j++ {
		// min := &arr[j]
		smol := &arr[j]
		smal := &importLines[j]
		for i := j + 1; i < len(arr); i++ {

			if arr[i] <= *smol {
				smol = &arr[i]
				smal = &importLines[i]
			}
			if i == len(arr)-1 {
				*smol, arr[j] = arr[j], *smol
				*smal, importLines[j] = importLines[j], *smal

				for range len(importLines) - j {
					fmt.Print("\033[A")
				}

				for o := j; o < len(importLines); o++ {

					fmt.Printf("\r%s", emptyStr(maxLenStr))
					fmt.Printf("\r%s", importLines[o])
					fmt.Print("\n")
					time.Sleep(time.Millisecond * time.Duration(speed))

				}
			}
		}
	}
	fmt.Println(arr)
	fmt.Print("Sort complete.")
}

func waitForUser() {
	enter := bufio.NewReader(os.Stdin)
	_, err := enter.ReadString('\n')
	if err != nil {
		fmt.Println("Error at input")
		waitForUser()
	}
}

func emptyStr(length float64) string {
	var spaceStr string
	spaceCount := 1
	for i := 0; i < int(length); i++ {
		spaceCount += 1
	}
	spaceStr = strings.Repeat(" ", spaceCount)
	return spaceStr
}

func draw(arr []int) *[]string {
	printedArr := []string{}

	for i := range arr {
		chars := 0
		num := arr[i]

		for j := 0; j < num; j++ {
			chars += 1
		}

		line := strings.Repeat("=", chars)
		printedArr = append(printedArr, line)
		// fmt.Println(printedArr)
	}

	return &printedArr
}

func main() {
	var arrLen int
	var speed int
	fmt.Println("This is a visual representation of a selection sorting algorithm")
	fmt.Printf("Enter array length you want to sort: ")
	fmt.Scan(&arrLen)
	fmt.Printf("Set speed (1-fast, 2-medium, 3-slow): ")
	fmt.Scan(&speed)
	speed = setSpeed(speed)

	selectSort(randArray(arrLen), speed)
}

func randArray(arrLen int) []int {
	if arrLen > 60 {
		arrLen = 60
	}

	newArr := []int{}

	for i := 0; i < arrLen; i++ {
		newNum := rand.IntN(80)
		if newNum == 0 {
			newNum = 1
		}
		newArr = append(newArr, newNum)
	}
	return newArr

}

func setSpeed(speed int) int {
	var speedIs int

	switch speed {
	case 0:
		speedIs = 1
	case 1:
		speedIs = 10
	case 3:
		speedIs = 70
	default:
		speedIs = 40
	}

	return speedIs
}
