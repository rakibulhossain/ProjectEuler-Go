package main

import (
	"fmt"
	"sort"
)

const lim = 4*int64(1e16)

type query struct{
	ind int
	val int64
}
type Key []query
func (this Key) Len() int{
	return len(this)
}

func (this Key) Less(i,j int) bool{
	return this[i].val<this[j].val
}

func (this Key) Swap(i,j int) {
	this[i],this[j] = this[j],this[i]
}

func main(){
	x := make([]int64,int64(1000))

	var t int
	var n int64
	fmt.Scanf("%d",&t)
	arr :=make([]query,t)
	ans :=make([]int64,t+1)
	for ca:=1;ca<=t;ca++ {
		fmt.Scanf("%d",&n)
		arr[ca-1]=query{ca,n}
	}



	sort.Sort(Key(arr))
	//fmt.Println(arr)
	x[0]=1
	x[1]=1
	in:=0
	sum:=int64(0)
	for i:=2;x[i-1]+x[i-2]<=lim;i++{
		x[i]=x[i-1]+x[i-2]
		if x[i]% 2 == 0 {
			for ;in<t && x[i]>arr[in].val; {
				ans[arr[in].ind] = sum

				in++
			}
			sum+=x[i]
		}
	//	fmt.Println(sum,in)
	}

	for ;in<t; {
		ans[arr[in].ind] = sum

		in++
	}

	for i:=1 ;i<=t;i++ {
		fmt.Println(ans[i])
	}
}