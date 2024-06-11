package main

import (
	"log"
	"time"
)

const timeout = 3 * time.Second

func main() {
	log.Println("Starting work.")

	// create new runner

	// add tasks

	// start runner

	log.Println("Process ended.")
}

func createTask() func(id int) {
	return func(id int) {
		log.Printf("Processor task: %d\n", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
