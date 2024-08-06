package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check (z int64) {

	if 	z <13 {
		fmt.Println("Your are Young")
	} else if z >= 13 && z <18 {
		fmt.Println("Your are a teenager")
	} else if z >= 18 {
		fmt.Println("Your are old")
	}
	return
}

func main() {

	numarr := []int64{}
	var add int64
	var i int64

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Array Length :")
	scanner.Scan()
	add, _ = strconv.ParseInt(scanner.Text(), 10, 64)

	for i = 0; i < add; i++ {
		fmt.Printf("Enter num for pos %d of array ", i)
		scanner.Scan()
		num, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		numarr = append(numarr, num)
	}

	fmt.Println(numarr)
	// fmt.Printf("%T",numarr)
	// fmt.Printf("%T",len(numarr))
	// numarr = int64(numarr[])

	for i = 0 ; i < add ; i++ {
		if numarr[i] < 0 {
			fmt.Println("Age is not valid, setting age to 0.")
			numarr[i] = 0
			z := numarr[i]
			check(z)
			z = z+3
			check(z)
		} else {
			z := numarr[i]
			check(z)
			z = z+3
			check(z)
		}
		fmt.Printf("\n\n")
	}

}
