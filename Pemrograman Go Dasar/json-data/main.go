package main

import "fmt"
import "encoding/json"

type User struct {
	FullName string `json:"Name"`
	Age int
}

func main() {
	// decode json string to struct
	var jsonString = `{"Name": "john wick", "Age": 27}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user name:", data.FullName)
	fmt.Println("user age:", data.Age)

	// decode json string to map string interface
	var dataMap map[string]interface{}
	json.Unmarshal(jsonData, &dataMap)

	fmt.Println("user name:", dataMap["Name"])
	fmt.Println("user age:", dataMap["Age"])

	// ditampung variable interface
	var dataInterface interface{}
	json.Unmarshal(jsonData, &dataInterface)

	var dataMapInterface = dataInterface.(map[string]interface{})
	fmt.Println("user name:", dataMapInterface["Name"])
	fmt.Println("user age:", dataMapInterface["Age"])

	// decode array json to array object
	var jsonArray = `[{"Name": "wick", "Age": 27}, {"Name": "ethan", "Age": 30}]`
	var dataArr []User

	var err1 = json.Unmarshal([]byte(jsonArray), &dataArr)
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}

	fmt.Println("user 1:", dataArr[0].FullName)
	fmt.Println("user 2:", dataArr[1].FullName)

	// encode objek ke json string
	var object = []User{{"john wick", 27}, {"ethan hunt", 30}}
	var jsonData1, err2 = json.Marshal(object)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	var jsonString1 = string(jsonData1)
	fmt.Println(jsonString1)
}