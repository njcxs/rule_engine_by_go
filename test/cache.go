package main

import "C"
import (
	"log"
	"time"

	cache "github.com/patrickmn/go-cache"
)

func main() {

	c := cache.New(30*time.Second, 10*time.Second)

	c.Set("Title", "Spring Festival", cache.DefaultExpiration)

	value, found := c.Get("Title")
	if found {
		log.Println("found:", value)
	} else {
		log.Println("not found")
	}

	time.Sleep(60 * time.Second)
	log.Println("sleep 60s...")
	value, found = c.Get("Title")
	if found {
		log.Println("found:", value)
	} else {
		log.Println("not found")
	}

}
