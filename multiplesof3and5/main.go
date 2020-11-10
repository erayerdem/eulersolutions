package main

import "fmt"

func main() {
	toplam :=0
	for i := 0; i < 1000 ; i+=3 {
			toplam+=i;
	}

	for i := 0; i < 1000 ; i+=5 {
		if (i%3!=0) {
			toplam+=i;}
	}

	fmt.Println(toplam)

}
