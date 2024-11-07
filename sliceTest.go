/* change  intPow so that it gives  error instead of string. that will  help with cases
 where  we get negative expo in a stream of postive expo's */

package ccv

import(
	"fmt"
)


func ArrToInt(arr []int) int {
	var y int = 0

	for i,j := len(arr) - 1 , 0 ; i >= 0 ; i,j = i-1 , j+1 {
		coeff , err :=  intPow(10 , j)

		if err != "" {
			fmt.Println(err)
			y = 0
			break
		} else {
			y = y + arr[i] * coeff
		}
	}
	return y
}

func IntToIntSlice(num int) []int {

	var intSlice []int =  []int{}
	for num  != 0 {
		intSlice = append(intSlice , num % 10)
		num /= 10
	}
	var resultIntSlice []int =  []int{}
	fmt.Println(resultIntSlice)
	for i := len(intSlice) - 1; i >= 0 ; i-- {
		resultIntSlice = append(resultIntSlice , intSlice[i] )
	}
//	fmt.Println("length bool :" , len(intSlice) ==  len(resultIntSlice))
	return resultIntSlice
}


func intPow(base , exponent int) (int , string) {

	if exponent < 0 {
		return  0 , "exponent should be non negative integer"
	}

	var powResult int = 1
	for i := 0 ; i < exponent ; i++ {
		 powResult *= base
	}
	return powResult , ""
}
