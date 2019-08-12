# HTTP-Stress-Tester
A quick HTTP stress tester. This allows you to choose how many goroutines are spamming GET requests to a specific path on a server.

# Usage
1. Build
`go build main.go`
2. Run
	- Linux: `./main {target} {# of threads}`
	- Windows: `main.exe {target} {# of threads}`
	- NOTE: target variable should be a url, # of threads should be an integer

# Features
1. Easy-To-Use
2. Allows you to choose the number of goroutines attacking a server
3. Shows real-time statistics on command line

# Screenshot
![Alt text](/master/screenshot.png "")
