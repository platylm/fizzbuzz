package fizzBuzz

import (
	"strconv"
)

func say(number int) string {

	y := map[int]string{
		3:  "Fizz",
		5:  "Buzz",
		15: "FizzBuzz",
	}
	str := []int{15, 3, 5} //อยากให้ตัวไหน run ก่อนให้เอาตัวนั้นขึ้น
	for _, str := range str {
		if number%str == 0 {
			return y[str]
		}
	}
	return strconv.Itoa(number)
}
