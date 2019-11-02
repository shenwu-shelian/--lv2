package main

import (
	"fmt"
	"time"
)

var (
	g,m,v int
	channel = make(chan int)
)
func hh(n int)  {  //设计一个判断是否是素数用的协程，名字都是细节
	var j,k int
	OuterLoop:
		for k=2;k<n; k++ {
			j = n%k
			if j==0 {
				break OuterLoop
			}else  if k==n-1{
				m=n
			}
		}
	channel <- m
}
func main()  {
	fmt.Print("2")//实在是设计不出来2是素数……原谅我
	var i int
	for i=2;i<=10000;i++ {
		go hh(i)
		v=g
		g =<-channel
		if m!=0 && v!=g {
			fmt.Printf("%d\n",g)
		}
	}
	time.Sleep(10 * time.Second)
}