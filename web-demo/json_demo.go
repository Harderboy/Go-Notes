package main

import (
	"encoding/json"
	"fmt"
)

type Company struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Country string `json:"country"`
}

func main() {
	jsonStr:=`
	{
		"id": 123,
		"name": "google",
		"country": "USA"
	}
	`
	c:=Company{}
	fmt.Println(c)
	_ =json.Unmarshal([]byte(jsonStr),&c)
	fmt.Println(c)

	bytes,_:=json.Marshal(c)
	fmt.Println(string(bytes))

	bytes2,_:=json.MarshalIndent(c,"","  ")
	fmt.Println(string(bytes2))
}