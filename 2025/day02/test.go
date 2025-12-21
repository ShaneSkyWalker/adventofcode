// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

func main() {

	// process := func(itr int64, tms int64) int64 {
	// 	sum := itr
	// 	for j := int64(0); j < tms - int64(1); j++ {
	// 		sum = sum * 10 + itr
	// 	}
	// 	return sum
	// }
	
	// 121212, 123,123
	divisor := 3
	process := func(itr int64, lens int64) int64 {
		sum := itr
		quotient := lens / int64(divisor) // 6 / 2 = 3
		fmt.Println(sum)
		for j := int64(0); j < quotient - 1; j++ {
			sum = sum * int64(math.Pow(10.0, float64(divisor))) + itr
			fmt.Println(sum)
		}
		return sum
	}

	// main process: create Possible Invalid ID with repeated first digit.
	var firstDivisorInt int64 = 123
	var lenRunes int64 = 6
	pssblInvalidId := process(firstDivisorInt, lenRunes)
	fmt.Println(pssblInvalidId)
}

