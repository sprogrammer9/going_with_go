package main // variadic functions are function that take n args

import "fmt"

func main() {

	bestFinish := championshipFinishes(12, 5, 6, 4, 3, 3) //using a function to declare a variable

	fmt.Println(bestFinish)

}

func championshipFinishes(finishes ...int) int {

	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i

		}
	}

	return best

}
