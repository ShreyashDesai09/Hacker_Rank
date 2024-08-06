package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var i int64

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Number : ")
	scanner.Scan()
	num , _ := strconv.ParseInt(scanner.Text(),10,0)
	

	for i = 1 ; i<=10 ; i++{
		ans := num * i 
		fmt.Printf("%d x %d = %d \n" ,num,i,ans)
	}

}
