package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("/Users/eray/dev/eulersolutions/13_largesum /file.txt")
	if err != nil {
		panic(err)
	}

	s := string(file)
	var  list[100] int
	split := strings.Split(s, "\n")

	for i, s2 := range split {
		atoi, err := strconv.Atoi(s2[0:11])
		if err != nil {
			panic(err)
		}
		list[i]=atoi
	}
	plus:=0
	for _, i2 := range list {
		plus+= i2
	}
	s2 := strconv.Itoa(plus)[0:10]
	fmt.Println(s2)
}
