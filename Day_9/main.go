package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func factorial() (answer int64) {

	var i int64
	var z int64
	var fact int64

	z = 1
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fact, _ = strconv.ParseInt(scanner.Text(), 10, 64)

	for i = 1; i <= fact; i++ {
		z = z * i
		answer = z
	}

	fmt.Println(answer)
	return answer

}

func main() {

	// var fact int64

	// var answer int64

	factorial()
}
