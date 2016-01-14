package fizzbuzz

import (
	"fmt"
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

func fizzbuzz() {

	fmt.Println("derp")





}


/*)	for i := 0; i <= 1000000; i++ {
		switch {
				case i%3 == 0 && i%5 == 0:
					fmt.Println("FizzBuzz")
				case i%3 == 0 && i%5 != 0:
					fmt.Println("Fizz")
				case i%3 != 0 && i%5 == 0:
					fmt.Println("Buzz")
				default:
					fmt.Println(i)
		}
	}
*/