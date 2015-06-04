package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Name string
type NameList []string

//
// Satisfy sort.Interface for the NameList type
// Make it sortable via sort.Sort
//

func (a NameList) Len() int           { return len(a) }
func (a NameList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NameList) Less(i, j int) bool { return a[i] < a[j] }

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Value(name string, place int) int64 {
	sum := int64(0)

	for i := 0; i < len(name); i++ {
		sum += (int64(name[i]) - 64)
	}

	return int64(place) * sum
}

func main() {

	sum := int64(0)
	var names []string

	// Open the datafile
	file, err := os.Open("problem022_names.txt")
	check(err)

	// Initialize the filescanner
	scanner := bufio.NewScanner(file)

	// close on exit
	defer func() {
		err := file.Close()
		check(err)
	}()

	scanner.Scan()
	names = append(names, scanner.Text())

	//
	// Split into names, unquoted
	//
	names = strings.Split(names[0], ",")

	for i := 0; i < len(names); i++ {
		names[i], err = strconv.Unquote(names[i])
		//fmt.Println(names[i])
		check(err)
	}

	sort.Sort(NameList(names))

	for i := 0; i < len(names); i++ {
		sum += Value(names[i], i+1)
	}

	fmt.Println("The total of all the name scores is:", sum)
}
