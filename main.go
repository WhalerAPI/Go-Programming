package main

import (
	"fmt"
	"time"
	"net/http"
	"github.com/whalelogic/Go-Programming/utils"
)



func main() {

	type T struct {a,b,c int}
	h := T{4,5,6}
	fmt.Println(h)
	c := utils.Collect(1,2,3,4,5)
	fmt.Println(c)
	x, y := utils.Swap("hello", 42)
	fmt.Println(x, y) 

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		h := r.Proto
		fmt.Fprintf(w, "Hello Protocol %s! Time is: %s\n", h, time.Now())
		// Response body: Hello Protocol HTTP/1.1! Time is: 2025-11-18 09:51:44.370398375 -0500 EST m=+30.972438668
	})

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		message := r.URL.Query().Get("message")
		if message == "" {
			message = "No message provided"
		}
		headers := r.Header
		for name, values := range headers {
			for _, value := range values {
				fmt.Fprintf(w, "Header: %s = %s\n", name, value)
			}
		}
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "Echo: %s", message)
	})

	http.ListenAndServe(":8080", nil)


}
