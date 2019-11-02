package main
//这是求1到20阶乘的程序x
import (
	"fmt"
	"time"
)

var (
	myres = make(map[int]int, 20)
	j int
	channel = make(chan int)
)
func factorial(n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myres[n] = res
	channel <- res //将结果用管道发出。
}

func main() {
	var i  int
	for i = 1; i <= 20; i++ {
		go factorial(i)
		j =<-channel//传递数据，并使之按顺序输出。
		fmt.Printf("myres[%d] = %d\n", i, j)
	}
	time.Sleep(5 * time.Second)
}