package main

import "fmt"

func main() {

	for i := 2520; i >2519 ; i=i+2 {
		if check(i) {
			fmt.Print(i)
			break
		}
	}

}

func check(val int)  bool {
	for i := 1; i <20 ; i++ {
		if val%i!=0{
			return false
		}
	}

	return true
}
