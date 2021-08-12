package main

import (
	"fmt"
)
func total(n int) int {
	n*=n+1
	n/=2
	return n
}
func main(){
	var n,t int
	_,_ = fmt.Scanf("%d",&t)
	for ca := 1;ca<=t;ca++ {
		_,_ =fmt.Scanf("%d",&n)
		n--
		sum := 0
		sum += 3*total(n/3)
		sum+=5*total(n/5)
		sum-=15*total(n/15)
		fmt.Println(sum)
	}
}