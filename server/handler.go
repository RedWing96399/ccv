package server

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/RedWing96399/ccv/ccv"
)

func validate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "the ccv number recieved through the request ")
}

func valC(w http.ResponseWriter, r *http.Request) {

	var numString []byte

	numString, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "error occured : %v", err.Error())
		return
	}
	// // fmt.Printf("%v bytes read \n", n)
	// fmt.Fprintln(w, "numstring : ", numString)

	numInt, err := strconv.Atoi(string(numString))
	if err != nil {
		fmt.Fprintf(w, "error occured : %v\n %v", err.Error(), numString)
		return
	}

	resultBool, err := ccv.CheckNum(numInt)
	if err != nil {
		fmt.Fprintf(w, "error occured : %v", err.Error())
		return
	}
	fmt.Fprintf(w, "the given number is %v\n", numInt)
	fmt.Fprintf(w, "the result of given number is : %v", resultBool)
}
