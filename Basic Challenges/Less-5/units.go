package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func letter() map[int]string {
	var letters = make(map[int]string)
	for i := 97; i < 123; i++ {
		letters[i] = string(i)
	}
	return letters
}

func payload(url string) {
	defer timeCost()()
	//ch <-struct {}{}
	//fmt.Println(url)
	respDataBase, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	respDataBase.Body.Close()
}

func timeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		second := tc.Seconds()
		if second < float64(3) {
			cost = false
		} else {
			cost = true
		}
	}
}
func timeCost1() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		second := tc.Seconds()
		fmt.Println(second)
	}
}
