package main

import "fmt"

func main()  {
 	count :=0
	maxi := 0
	for count<10001 {
		maxi++
		if isPrime(maxi) {
			count++
		}

	}
	fmt.Print(maxi)
}


func isPrime(val int)  bool {
	for i := 2; i <val/2 ; i++ {
		if val%i ==0 {
			return false
		}
	}
	return true
}