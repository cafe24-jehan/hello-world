package main

import "fmt"

func repeatMe(words ...string){
	fmt.Println(words)

}

func main() {
	repeatMe("jiEun", "bumSu", "hyoungAe", "yeEun", "down")
}