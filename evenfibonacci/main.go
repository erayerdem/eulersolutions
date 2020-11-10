package main

import "fmt"

func main()  {
	sum :=2
	left,right,temp:=1,2,0

	for right<4000000 {

		temp=right
		right=left+right
		left=temp
		if right%2==0 {
			sum+=right
		}

	}

fmt.Println(sum)




}

