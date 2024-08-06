package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// var num int

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("ENTER NUMBER : ")
	scanner.Scan()
	num , _ := strconv.ParseInt(scanner.Text(), 10, 64)

	if 1 < num && num <= 5 {
		if num%2 == 0 {
			fmt.Println("Not Weird")
		} else {
			fmt.Println("Weird")
		}
	} else if num >= 6 && num <= 20 {
		if num %2 == 0 {
			fmt.Println("Weird")
		} else {
			fmt.Println("not weird")
		}
	} else if num >= 20{
		if num % 2 ==0 {
			fmt.Println("not weird")
		} else  {				 
				fmt.Println("weird")
		}
	} 


}
