package ccv

import (

	"fmt"
//	"os"
)


func CheckNum (num int) (bool , string) {

	if  num < 0 {
		return  false , "number is negative, only positive numbers should be given as input"
	}

	var  numLen int = GetNoOfDigits(num)

	if numLen > 19 || numLen < 13 {
		return false , " the input should be a number with  13 to 19 digits . no more , no less "
	}

	return  true , "given number is a valid credit card number"
}


func GetNoOfDigits(num int) int {

	var  digitCount  int = 0
	if  num  == 0 {
		digitCount = 1
	} else {
		for num != 0 {
			fmt.Println(num)
			digitCount++
			num  /= 10
		}
	}

	return digitCount
}