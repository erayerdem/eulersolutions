package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)
var max = uint(0)

func main() {


	file, err := ioutil.ReadFile("/Users/eray/dev/eulersolutions/11_largestproductinagrid /file.txt")

	if err != nil {
		panic(err)
	}

	s := string(file)
	lines := strings.Split(s, "\n")
	arr := make([][]uint, 20)
	var strArr []string
	for i := range lines {
		arr[i] = make([]uint, 20)
		strArr = strings.Split(lines[i], " ")
		for j := range arr[i] {
			intvalue, _ := strconv.Atoi(strArr[j])
			arr[i][j] = uint(intvalue)
		}

	}

	// left to rigt
	for i := 0; i <17 ; i++ {
		for j := 0; j <17 ; j++ {
			temp := arr[i][j] *arr[i][j+1]*arr[i][j+2]*arr[i][j+3]
			setIfBigger(temp)
		}
	}
	// up to down
	for i := 0; i <17 ; i++ {
		for j := 0; j <17 ; j++ {
			temp := arr[j][i] *arr[j+1][i]*arr[j+2][i]*arr[j+3][i]
			setIfBigger(temp)
		}
	}


	for i := 0; i < 17; i++ {
		for j := 0; j < 17; j++ {
			temp := arr[i][j] * arr[i+1][j+1] * arr[i+2][j+2] * arr[i+3][j+3]
			setIfBigger(temp)
		//	00 11 22 33
		//	01 12 23 34
		//	01 12 23 35
		}
	}

	for i := 3; i < 20; i++ {
		for j := 0; j < 17; j++ {
			temp := arr[i][j] * arr[i-1][j+1] * arr[i-2][j+2] * arr[i-3][j+3]
			setIfBigger(temp)
		}
	}




	fmt.Println(max)

}
func setIfBigger (val uint) {
	if val > max {
		max = val
	}
}