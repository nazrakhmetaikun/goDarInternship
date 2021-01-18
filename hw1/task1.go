package main 
import "fmt"

func main(){
	a:=0
	b:=0
	c:=0
	d:=0
	fmt.Scanln(&a,&b,&c,&d)
	fmt.Println(min(a,b,c,d))
}
func min(a int, b int, c int, d int) int{
	min:=0
	min1:=0
	min2:=0
	if a>b{
		min1=b
	} else{
		min1=a
	}
	if c>d{
		min2=d
	} else{
		min2=c
	}
	if min1>min2{
		min=min2
	} else{
		min=min1
	}

	return min
}