package main

import (
	"strconv"
	"strings"
)

func modifySlice(s []int) {
	// Modifying the slice inside the function
	for i := range s {
		s[i] *= 2
	}
}

func main() {
	var builder strings.Builder
	for i := 0; i < 10; i++ {
		builder.WriteString(strconv.Itoa(i))
		builder.WriteString(",")
		println(builder.String())
	}
}
