package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // alias to be displayed in display data
	Price    int
	Platform string
	Password string   `json:"-"` // field to be ommitted when fetched as json
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON lesson in Go")
	//EncodeJSON()
	decodeJSON()
}

func EncodeJSON() {
	lcoCourses := []course{
		{"ReactJS", 299, "LCO", "abc123", []string{"web dev", "JS"}},
		{"NextJS", 199, "LCO", "abc123", []string{"full stack", "TS"}},
		{"NodeJS", 239, "LCO", "owili123", nil},
	}

	// package this data as JSON data

	//finalJSON, err := json.Marshal(lcoCourses)
	finalJSON, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n ", finalJSON)

}

func decodeJSON() {
	jsonDataFromWeb := []byte(`
	{
                "coursename": "NextJS",
                "Price": 199,
                "Platform": "LCO",
                "tags": [
                        "full stack",
                        "TS"
                ]
        }
	`)
	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON NOT VALID!")
	}
	// some cases where you just want ot add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and the type is %T\n", k, v, v )
	}
}
