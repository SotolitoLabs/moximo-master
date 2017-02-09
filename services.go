package main

import (
  "fmt"
  "net/http"
)



/* 
 * Index returns a flag showing that i'm the master
*/

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "MASTER_OK")
}

/*
 * NextIP returns the next ip in the list
*/

func NextIP(w http.ResponseWriter, r *http.Request) {
  ip := get_next_ip()
  fmt.Fprintln(w, ip)
}

/*
 * get_next_ip takes the last ip from the ip list and returns the next
 * available.
*/

func get_next_ip() string {
  return "10.253.0.2"
}
