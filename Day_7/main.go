package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	arr := []int64{}
	var anum int64
	var i int64
	arr2 := []int64{}

	// var z int64

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("ENTER ARRAY LENGTH : ")
	scanner.Scan()
	anum, _ = strconv.ParseInt(scanner.Text(), 10, 64)

	if anum > 0 {
		for i = 0; i < anum; i++ {
			fmt.Printf("Array no %d = ", i)
			scanner.Scan()
			num, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				fmt.Println("ERROR ", err)
			} else {
				arr = append(arr, num)
			}
		}
	}

	// arr = strconv.FormatInt(arr, 10, 64)
	fmt.Println("ARRAY LENGTH :", anum)
	fmt.Println("ARRAY : ", arr)

	// var j int64
	z := anum

	for i := z; i > 0; i-- {
		arr2 = append(arr2, arr[i])
	}

	fmt.Println(arr2)

}
