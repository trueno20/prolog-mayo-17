package main

import "frm"

func main(){
	sum := 1
	for  sum < 1000 {
	sum += sum
	}
	fmt.Println(sum)
}