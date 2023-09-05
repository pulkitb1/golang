package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	encodeJson()
	decodeJson()
}

func encodeJson() {
	lcoCourses := []course{
		{"ReactJsBootcamp", 299, "Youtube", "abc123", []string{"web-dev"}},
		{"AngularBootcamp", 199, "Youtube", "bcd123", nil},
	}

	// package this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))
}

func decodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJsBootcamp",
		"Price": 299,
		"website": "Youtube",
		"tags": [
				"web-dev"
		]
	}`)

	var lcoCourse course

	checkValidJson := json.Valid(jsonDataFromWeb)

	if checkValidJson {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Println(key, value)
	}
}
