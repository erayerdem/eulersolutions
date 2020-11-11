package main

import (
	"fmt"
	"math"
)

func main()  {
	 c:=0.0
	 okdur:=false
	for i := 1; i <10000 ; i++ {

		for j := 1; j <10000 ; j++ {
			power:=math.Pow(float64(i),2)+math.Pow(float64(j),2)
			c=math.Sqrt(power)
			if float64(i)+float64(j)+c==1000 && c!=0{
				fmt.Println(i*j*int (c))

				okdur=true
				break
			}

		}
		if okdur {
			break
		}

	}




	// it s complex must be simple

}
