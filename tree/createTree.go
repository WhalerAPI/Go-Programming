package tree

import (
	"fmt"
	"golangtour/tree"

)

// Greatest common Divisor 

var p, q int


fmt.Println(p, q)


func Gcd(p, q) r float32 {
	if q == 0 {
		return q
	}
	else {
		r := q%p 
	}
	return Gcd(q, r)
}
