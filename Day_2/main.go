package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var meal_cost float64
	var tip float64
	var tax float64

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	meal_cost, _ = strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	tip, _ = strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	tax, _ = strconv.ParseFloat(scanner.Text(), 64)

	tip_per := meal_cost * tip / 100
	tax_per := meal_cost * tax / 100

	total := meal_cost + tax_per + tip_per
	answer := int(total)

	z := total - float64(answer)

	if z > 0.49 {
		answer = answer + 1
		fmt.Println(int(answer))
	} else {
		fmt.Println(int(answer))
	}

	// fmt.Println(z)

}
