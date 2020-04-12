package main

import "log"

func main() {
	log.Println(Add(12, 12))
}

func Add(x, y int) int {
	return x + y
}
