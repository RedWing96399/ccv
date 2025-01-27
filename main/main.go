package main

import (
	"flag"
	"fmt"

	"github.com/RedWing96399/ccv/ccv"
)

func main() {

	var ccn int
	flag.IntVar(&ccn, "num", 111, "input for credit card")
	flag.Parse()

	resultBool, err := ccv.CheckNum(ccn)
	if err != nil {
		fmt.Println(err.Error(), err)
	}
	fmt.Println(resultBool)

}
