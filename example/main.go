package main

import (
	"log"

	"github.com/xyproto/simpleredis"
)

func main() {
	// Create a connection pool, connect to the given redis server
	pool := simpleredis.NewConnectionPool()

	// Use this for connecting to a different redis host/port
	// pool := simpleredis.NewConnectionPoolHost("localhost:6379")

	// Close the connection pool when this function returns
	defer pool.Close()

	// Create a list named "greetings"
	list := simpleredis.NewList(pool, "greetings")

	// Add "hello" to the list
	err := list.Add("hello")
	if err != nil {
		log.Fatalln("Could not add an item to list! Is Redis up and running?")
	}

	// Get the last item in the list
	item, err := list.GetLast()
	if err != nil {
		log.Fatalln("Could not fetch the last item from the list!")
	}
	log.Println("The value of the stored item is:", item)

	// Remove the list
	err = list.Remove()
	if err != nil {
		log.Fatalln("Could not remove the list!")
	}
}