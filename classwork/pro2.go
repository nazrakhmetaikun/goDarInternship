package main

import
	"fmt"


func main(){
	sum:=0
	a:=0
	fmt.Scanln(&a)
	if a>0{
		for i:=0; i<=int(a); i++ {
			sum+=i
		}
		fmt.Println(sum)
	} else{
		for i:=a; i<=1; i++ {
			sum+=i
		}
		fmt.Println(sum)
	}
}
