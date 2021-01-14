package main
import "fmt"
func main(){
	sum:=0
	a:=0
	fmt.Scanln(&a)
	for i:=0; i<=a; i++ {
		sum=sum+i
	}
	fmt.Println(sum)
}
