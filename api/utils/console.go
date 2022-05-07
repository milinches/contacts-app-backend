package utils

import (
	"encoding/json"
	"log"
)

func Prettier(data interface{}) {
	pretty, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(string(pretty))
}