package fizzbuzz

import (
	"strconv"
)

func fbdetection(i int) string {
	switch {
		case i%3 == 0 && i%5 == 0:
			return "FizzBuzz"
		case i%3 == 0 && i%5 != 0:
			return "Fizz"
		case i%3 != 0 && i%5 == 0:
			return "Buzz"
		default:
			return strconv.Itoa(i)
	}

}

/*
func fizzbuzz() {
	fmt.Println("derp")
}
*/
