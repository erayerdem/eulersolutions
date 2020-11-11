package main

import "fmt"

func main() {

	var ints [1001]int
	sum:=0
	for i := 0; i < 1000 ; i+=3 {
			ints[i]=1
	}

	for i := 0; i < 1000 ; i+=5 {
		ints[i]=1
	}

	for i1, i2 := range ints {
		if i2==1 {
		sum+=i1
		}
	}

	fmt.Println(sum)
}
