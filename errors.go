package main

import (
	"log"
)

/*
 * CheckError checks an error
 */

func CheckError(e error) {
	if e != nil {
		log.Printf("Error: %s", e)
		panic(e)
	}
}
