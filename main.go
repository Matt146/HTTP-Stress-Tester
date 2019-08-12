package main

import (
	"net/http"
	"io"
	"io/ioutil"
	"os"
	"fmt"
	"strconv"
	"time"
)

var requests int64
var startingTime int64

func kill(target string) {
	for {
		resp, err := http.Get(target)
		if err != nil {
			fmt.Printf("\n[Error]: Unable to perform GET request on target machine %s\n", target)
			fmt.Printf("\tHint: server is down, not connected to internet\n")
			fmt.Printf("\tError Code: %s\n\n", err.Error())
			break
		}
		io.Copy(ioutil.Discard, resp.Body)
		requests += 1
		currentTime := time.Now().Unix()
		if requests % 5 == 0 {
			fmt.Printf("\rRequests Per Second: %d | Status Code: %d\r", requests/(currentTime-startingTime), resp.StatusCode)
		}
	}
}

func kill_threaded(target string, threads int) {
	for x := 0; x < threads; x++ {
		go kill(target)
	}
}

func main() {
	requests = 0
	fmt.Println("[Info]: Initializing stress test!")
	if len(os.Args) > 2 {
		fmt.Printf("[Info]: Target=%s\n", os.Args[1])
		fmt.Printf("[Info]: Threads=%s\n", os.Args[2])
		fmt.Println("Note: press Ctrl+C to sigint/kill process")
		threadCount, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("[Error]: Second argument, '# of threads' must be an integer")
		}
		startingTime = time.Now().Unix()
		time.Sleep(time.Second * 1)
		kill_threaded(os.Args[1], threadCount-1)
		kill(os.Args[1])
	} else {
		fmt.Println("[Error]: Invalid usage!")
		fmt.Println("[Info]: {target} {# of threads}")
		fmt.Println("\tNote: include scheme and path, as a GET request will be performed on the given target and path")
	}
}
