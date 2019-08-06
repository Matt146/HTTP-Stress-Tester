package main

import (
	"net/http"
	"io"
	"io/ioutil"
	"os"
	"fmt"
)

func kill(target string) {
	for {
		resp, err := http.Get(target)
		if err != nil {
			break
		}
		io.Copy(ioutil.Discard, resp.Body)
	}
}

func kill_threaded(target string) {
	for x := 0; x < 1000; x++ {
		go kill(target)
	}
}

func main() {
	fmt.Println("[Info]: Initializing stress test!")
	if len(os.Args) > 1 {
		fmt.Printf("[Info]: Target=%s\n", os.Args[1])
		fmt.Println("Note: press Ctrl+C to sigint/kill process")
		kill_threaded(os.Args[1])
		kill(os.Args[1])
	} else {
		fmt.Println("[Error]: Invalid usage!")
		fmt.Println("[Info]: {target}")
		fmt.Println("\tNote: include scheme and path, as a GET request will be performed on the given target and path")
	}
	
}
