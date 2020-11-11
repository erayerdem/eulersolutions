package main

import (
	"fmt"
	"strconv"
)

func  main()  {

	for i := 999999; i > 100000 ; i-- {
		if  isPalindrome(i) {
				if carpin(i) {
					fmt.Println(i)
					break
				}
		}

	}

}
func carpin(val int) bool{

	for i := 100; i <999 ; i++ {
			if val%i==0 && val /i <1000{
				return true
			}
	}

	return false
}


func reverse(val string)  (result string){
	for _, v := range val {
		result = string(v) + result
	}
	return
}

func isPalindrome(val int)  bool{
	str:= strconv.Itoa(val)
	if str == reverse(str) {
		return true
	}

	return false
}
