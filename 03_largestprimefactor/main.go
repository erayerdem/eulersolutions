package main

import "fmt"

func main() {
	 prime := 600851475143
	var ints [] int
	bolecek :=2


	for prime>1 {

		if prime%bolecek==0 {
			prime=prime/bolecek
			ints = append(ints, bolecek)
			bolecek=1
		}

			bolecek++

	}

	for i := len(ints)-1; i >=0 ; i-- {

		if 	 isPrime(ints[i]) {
			fmt.Println(ints[i])
			break
		}
	}
}

func isPrime(val int)  bool {
	for i := 2; i <val/2 ; i++ {
		if val%i ==0 {
			return false
		}
	}
	return true
}
