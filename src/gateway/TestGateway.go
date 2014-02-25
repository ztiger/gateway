package main

import (
	//"fmt"
	"log"
)

func main() {
	log.Println(15 % 10) // 5
	b := byte('0' + 15%10)
	log.Println(b)
	log.Println(4 / 10)
}
