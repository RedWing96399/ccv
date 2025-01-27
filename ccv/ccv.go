package ccv

import (
	"errors"
	// "os"
)

var errNegativeNum error = errors.New("number is negative, only positive numbers should be given as input")
var errNumLen error = errors.New("the input should be a number with  13 to 19 digits . no more , no less")

func CheckNum(num int) (bool, error) {

	if num < 0 {
		return false, errNegativeNum
	}

	var numLen int = getNoOfDigits(num)

	if numLen > 19 || numLen < 13 {
		return false, errNumLen
	}

	numSlice := intToIntSlice(num)
	resultBool := LuhnCheck(numSlice)

	return resultBool, nil
	//need to implement another function that takes this boolean to generate a customized message regarding the
	//validity of the number
}

func getNoOfDigits(num int) int {

	var digitCount int = 0
	if num == 0 {
		digitCount = 1
	} else {
		for num != 0 {
			//fmt.Println(num)
			digitCount++
			num /= 10
		}
	}

	return digitCount
}
