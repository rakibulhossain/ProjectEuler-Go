package main

import (
	"fmt"
	"math"
)

func main(){
	var (
		t int
		n int64
	)

	fmt.Scanf("%d",&t)

	for ca:=1;ca<=t;ca++ {
		fmt.Scanf("%d",&n)
		ans:=int64(0)
		for i:=int64(2) ; i<=int64(math.Sqrt(float64(n))) ;i++ {
			if n % i ==0 {
				for ;n%i==0; {
					n/=i
					ans=i
				}
			}
		}
		if n>1 {
			ans=n
		}
		fmt.Println(ans)
	}
}
