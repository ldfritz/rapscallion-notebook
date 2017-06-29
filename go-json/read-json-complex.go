package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type ResponseData struct {
	Data []Message `json:"data"`
}

type Message struct {
	Id   string
	From struct {
		Name string
		Id   string
	}
	Message string
	Actions []struct {
		Name string
		Link string
	}
	Type         string
	Created_time string
	Updated_time string
}

var json_data = `
{
   "data": [
      {
         "id": "X999_Y999",
         "from": {
            "name": "Tom Brady", "id": "X12"
         },
         "message": "Looking forward to 2010!",
         "actions": [
            {
               "name": "Comment",
               "link": "http://www.facebook.com/X999/posts/Y999"
            },
            {
               "name": "Like",
               "link": "http://www.facebook.com/X999/posts/Y999"
            }
         ],
         "type": "status",
         "created_time": "2010-08-02T21:27:44+0000",
         "updated_time": "2010-08-02T21:27:44+0000"
      },
      {
         "id": "X998_Y998",
         "from": {
            "name": "Peyton Manning", "id": "X18"
         },
         "message": "Where's my contract?",
         "actions": [
            {
               "name": "Comment",
               "link": "http://www.facebook.com/X998/posts/Y998"
            },
            {
               "name": "Like",
               "link": "http://www.facebook.com/X998/posts/Y998"
            }
         ],
         "type": "status",
         "created_time": "2010-08-02T21:27:44+0000",
         "updated_time": "2010-08-02T21:27:44+0000"
      }
   ]
}
`

func main() {
	data := ResponseData{}
	json.Unmarshal([]byte(json_data), &data)
	for _, message := range data.Data {
		fmt.Println(message)
	}
	json_dump, dump_err := json.Marshal(data)
	if dump_err != nil {
		log.Fatal(dump_err)
	}
	fmt.Println(string(json_dump))
}
