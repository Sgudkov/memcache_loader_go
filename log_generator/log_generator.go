package log_generator

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Log struct {
	Time    string `json:"time"`
	Level   string `json:"level"`
	Message string `json:"message"`
}

func JsonLog(file string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	for i := 0; i < 20; i++ {
		f.WriteString(fmt.Sprintf(`{"time": "%s", "level": "%s", "message": "test"}`, time.Now().Format(time.RFC3339), strconv.Itoa(i)) + "\n")
	}
}

func Is_file_empty(file string) bool {
	f, err := os.Open(file)
	if err != nil {
		return true
	}
	defer f.Close()
	_, err = f.Stat()
	if err != nil {
		return true
	}
	return false
}

func (l Log) Get_file_data(file string) []Log {

	file_json, err := os.Open(file)
	if err != nil {
		log.Println(err)

	}

	defer file_json.Close()

	decoder := json.NewDecoder(file_json)
	var items []Log
	for decoder.More() {
		var item Log
		err := decoder.Decode(&item)
		if err != nil {
			log.Println(err)
			continue
		}
		items = append(items, item)
	}

	return items
}
