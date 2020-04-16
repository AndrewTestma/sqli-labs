package main

import (
	"fmt"
)

type Info struct {
	name     string
	length   int
	subCount int
}

var cost bool = false
var tbNames = ""
var ch = make(chan struct{})

func main() {
	letters := letter()
	dbinfo := getDbAndTableCount(letters)
	fmt.Printf("db name length:%d;db name content:%s;db tables count:%d\n", dbinfo.length, dbinfo.name, dbinfo.subCount)
}
