package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {	
	
	reader := bufio.NewScanner(os.Stdin)
	h := reader.Scan()
	fmt.Println("Hello, World !")
	fmt.Printf("%v" , h)
}
