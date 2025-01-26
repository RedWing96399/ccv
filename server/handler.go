package server

import (
	"fmt"
	"net/http"
)

// handler is still incomplete
func validate(w http.ResponseWriter, r *http.Request) {

	ccvNum, err := r.Body.Read([]byte{})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Fprintf(w, "the ccv number recieved through the request is : %d", ccvNum)
}
