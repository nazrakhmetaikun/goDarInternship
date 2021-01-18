package main 
import "fmt"

func main(){
	a:=0
	b:=0
	c:=0
	d:=0
	fmt.Scanln(&a,&b,&c,&d)
	fmt.Println(min(min(a,b),min(c,d)))
}

func min(a int, b int)int{
	min:=0
	if a<b{
		min=a
	} else{
		min=b
	}
	return min
}