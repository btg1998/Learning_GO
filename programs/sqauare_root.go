
// Function to calculate square root of a given number using newton's iterative meth0od
//  and also to print the value of the approximate square root at each iteration
package main

import (
	
)

func sqr(x float64){
	app:=0.5*x
	for i:=0;i<10;i++{
		 temp := 0.5 * (app + x/app)
        	 app = temp
    		 print(app)
	}
}
func main() {
	b:=64.0
	sqr(b)
}

