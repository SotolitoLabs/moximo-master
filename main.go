package main

import (
	"flag"
	"log"
	"net/http"
)

type Config struct {
	Nodes string
	Debug bool
}

var Cfg Config

func main() {
	nodesFlag := flag.String("nodes", "/etc/moximo/nodes", "Node Config file")
	debugFlag := flag.Bool("d", false, "Enable debugging")
	flag.Parse()

	Cfg = Config{Nodes: *nodesFlag, Debug: *debugFlag}

	log.Printf("Starting moximo master with flags:\nNodes file: %s\nDebug mode:%b\n",
		Cfg.Nodes, Cfg.Debug)
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8081", router))
}
