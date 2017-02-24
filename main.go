package main

import (
	"flag"
	"log"
	"net/http"
)

type config struct {
	nodes string
	debug bool
}

var Config config

func main() {
	nodesFlag := flag.String("nodes", "/etc/moximo/nodes", "Node config file")
	debugFlag := flag.Bool("d", false, "Enable debugging")
	flag.Parse()

	Config := &config{nodes: *nodesFlag, debug: *debugFlag}

	log.Printf("Starting moximo master with flags:\nNodes file: %s\nDebug mode:%b\n",
		Config.nodes, Config.debug)
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8081", router))
}
