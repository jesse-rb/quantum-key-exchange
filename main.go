package main

import (
	"math/rand"
	"time"
)



func init() {
	// Get new random seed for this run.
	rand.Seed(time.Now().UTC().UnixNano()) // use time as seed
}

func main() {
	
}