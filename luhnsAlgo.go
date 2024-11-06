package ccv

import(
	"fmt"
)


func LuhnCheck(intSlice []int) bool {

	var rOddSum , rEvenSum int

	fmt.Println("odd value sum list :")
	for i:= len(intSlice) -1 ; i >= 0 ;  i -= 2 {
		fmt.Println(intSlice[i])
		rOddSum += intSlice[i]
	}

	fmt.Println("Even value sum list :")

	for i , tempVar := len(intSlice) -2 , 0 ; i >= 0 ; i -= 2 {

		tempVar = intSlice[i]
		tempVar *= 2
		if tempVar > 9 {
			tempVar -= 9
		}
		fmt.Println(intSlice[i] , tempVar )
		rEvenSum += tempVar
	}
	fmt.Printf("oddsum is : %d , evensum is : %d \n" , rOddSum , rEvenSum )

	var resultBoolean = ( ( rOddSum + rEvenSum ) % 10 == 0 )
	return   resultBoolean
}