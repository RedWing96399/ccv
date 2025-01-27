package ccv

import (
	"errors"
	"fmt"
)

var errNegativePower error = errors.New("exponent is a negative integer")

func sliceToInt(slic []int) int {
	var y int = 0

	for i, j := len(slic)-1, 0; i >= 0; i, j = i-1, j+1 {
		coeff, err := intPow(10, j)

		if err != nil {
			fmt.Println(err)
			y = 0
			break
		} else {
			y = y + slic[i]*coeff
		}
	}
	return y
}

func intToIntSlice(num int) []int {

	var intSlice []int = []int{}
	for num != 0 {
		intSlice = append(intSlice, num%10)
		num /= 10
	}
	var resultIntSlice []int = []int{}
	// fmt.Println(resultIntSlice)
	for i := len(intSlice) - 1; i >= 0; i-- {
		resultIntSlice = append(resultIntSlice, intSlice[i])
	}
	// fmt.Println("length bool :" , len(intSlice) ==  len(resultIntSlice))
	return resultIntSlice
}

func intPow(base, exponent int) (int, error) {

	if exponent < 0 {
		return 0, errNegativePower
	}

	var powResult int = 1
	for i := 0; i < exponent; i++ {
		powResult *= base
	}
	return powResult, nil
}
