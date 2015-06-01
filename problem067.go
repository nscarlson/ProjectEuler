package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	triangle [][]int
	maxsum   int64
	index    int
)

//
//	Check an error thrown
//

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func maxinslice(arr []int) int {
	max := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

//
//	Top-down calculation of maxpath. Bottom-up is more beautiful, but WHATEVER.
//	Increments each current cell by max(left, right) maxpath
//
func maxpath() int {
	for row := 1; row < len(triangle); row++ {
		for col := 0; col <= row; col++ {
			if col == 0 {
				triangle[row][col] += triangle[row-1][col]
			} else if col == row {
				triangle[row][col] += triangle[row-1][col-1]
			} else {
				triangle[row][col] += max(triangle[row-1][col], triangle[row-1][col-1])
			}
		}
	}
	return maxinslice(triangle[len(triangle)-1])
}

//
//	An older version of maxpath, left here for historical purposes.
//

/*
func maxpath(row, pos int, currsum int64) {
	if row > len(triangle)-1 {
		return
	}

	currsum += int64(triangle[row][pos])
	//fmt.Println("Current", currsum)

	if currsum > maxsum {
		maxsum = currsum
		fmt.Println("Max:", maxsum)
	}

	//fmt.Printf("[%v][%v]", row, pos)
	maxpath(row+1, pos, currsum)
	maxpath(row+1, pos+1, currsum)
}
*/

func main() {
	var lines []string

	// Open the datafile
	file, err := os.Open("problem067.txt")
	check(err)

	// Initialize the filescanner
	scanner := bufio.NewScanner(file)

	// close on exit
	defer func() {
		err := file.Close()
		check(err)
	}()

	for i := 0; scanner.Scan(); i++ {
		lines = append(lines, scanner.Text())
	}

	triangle = make([][]int, len(lines))

	//
	//	parse into a 2D array of integers
	//
	for i := 0; i < len(lines); i++ {
		result := strings.Split(lines[i], " ")
		triangle[i] = make([]int, len(result), len(result))

		for pos := 0; pos < len(result); pos++ {
			number, err := strconv.Atoi(result[pos])
			check(err)

			triangle[i][pos] = number
		}
	}

	fmt.Println("Maximum path sum is:", maxpath())
}
