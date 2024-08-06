package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var i int64 = 4
	var f float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("INT : ")
	scanner.Scan()
	i2, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("FLOAT : ")
	scanner.Scan()
	f2, _ := strconv.ParseFloat(scanner.Text(), 64)
	fmt.Printf("STRING : ")
	scanner.Scan()
	s2 := scanner.Text()

	fmt.Println(i + i2)
	fmt.Printf("%f \n", f+f2)
	fmt.Println(s + s2)

}
