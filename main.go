package main

import (
	"log"
	"time"

	"runner/runner"
)

const timeout = 3 * time.Second

func main() {
	log.Println("Starting work.")

	r := runner.New(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		log.Println("runner stopped with error:", err)
	}

	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d\n", id)
		time.Sleep(time.Duration(id+1) * time.Second)
	}
}
