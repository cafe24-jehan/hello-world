package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	// 	defer : 코드가 끝나고 실행시킴
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, up := lenAndUpper("jieun")
	fmt.Println(totalLength, up)
}
