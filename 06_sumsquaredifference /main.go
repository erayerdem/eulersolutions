package main

import (
	"fmt"
	"math"
)

func main()  {
	sumsquare:=math.Pow(100*101/2,2)
	squaresum:=0.0

	for i := 1; i <101 ; i++ {
		squaresum+=math.Pow(float64(i),2)
	}

	fmt.Print(int64(sumsquare-squaresum))

}
