package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	ip := getNextIp()
	writeLine(Cfg.Nodes, ip)
	fmt.Fprintln(w, ip)
}

/*
 * getNextIp takes the last ip from the ip list and returns the next
 * available.
 */

func getNextIp() string {
	last := getLastLine(Cfg.Nodes)
	splitted := strings.Split(last, ".")
	next, _ := strconv.Atoi(splitted[3])
	next++
	splitted[3] = strconv.Itoa(next)
	newip := strings.Join(splitted, ".")
	if Cfg.Debug {
		log.Printf("getNextIP:: Last IP: %s", last)
	}
	log.Printf("Next IP: %s", newip)
	return newip
}

/*
 * getLastLine Returns the las line from a file
 *
 * TODO: read just the last line instead of the whole file
 */

func getLastLine(filename string) string {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Printf("Error Reading file %s: %s ", filename, err)
		panic(err)
	}

	file := make([]string, 0)
	input := bufio.NewScanner(f)

	//TODO check for error reading
	for input.Scan() {
		text := input.Text()
		// Ignore empty lines
		if text != "" {
			file = append(file, text)
		}
	}

	// Return the last line
	last := len(file) - 1
	return file[last]
}

/*
 * writeLine writes a line to the file
 *
 */

func writeLine(file string, line string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	CheckError(err)
	defer f.Close()
	f.Seek(0, 2)
	_, err = f.WriteString(line + "\n")
	CheckError(err)
}
