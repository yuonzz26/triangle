package main

import "fmt"

func main() {
	var i, j, k int
	//i,j는 두 개의 반복문에 쓰일 변수

	fmt.Scan(&k)

	for i = 1; i <= k; i++ {
		for j = 1; j <= i-1; j++ {
			fmt.Print("o ")
		}
		fmt.Print("* ")
		fmt.Println("")
	}
}
