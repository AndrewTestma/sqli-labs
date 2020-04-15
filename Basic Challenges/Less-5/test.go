package main

import (
	"fmt"
	"time"
)

func main() {
	defer timeCost1()()
	ch := make(chan struct{})
	a := [...]string{"a", "b", "c", "d"}
	for i := 0; i < len(a); i++ {
		go pri(a[i])
	}
	<-ch
}

func pri(n string) {
	time.Sleep(3 * time.Second)
	fmt.Println(n)
	ch <- struct{}{}
}
