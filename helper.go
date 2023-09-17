package main

import (
	"log"
)

// handle error func
func handleErr(err error, fatal bool, msg string) error {
	// log error and exit if fatal is true
	if err != nil {
		if !fatal {
			log.Println(msg)
		} else {
			log.Fatal(msg)
		}
		return err
	}

	// return nil if no error
	return nil
}
