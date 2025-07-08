package main

import (
	"fmt"
	"time"

	"github.com/Sgudkov/memcache_loader_go/log_generator"
	memcache "github.com/bradfitz/gomemcache/memcache"
)

func main() {

	var mc = memcache.New("127.0.0.1:11211")

	mc.Timeout = 10 * time.Second

	if log_generator.Is_file_empty("test.log") {
		log_generator.JsonLog("test.log")
	}

	log := log_generator.Log{}

	log_list := log.Get_file_data("test.log")

	for _, v := range log_list {
		err := mc.Set(&memcache.Item{Key: v.Level, Value: []byte(v.Message)})
		if err != nil {
			fmt.Println(err)
		}

	}

	for _, v := range log_list {
		fmt.Println(mc.Get(v.Level))

	}

}
