package main

import (
	"./broadcast"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type command struct {
	Command, Result string
	FullCommand     []string
	Running, Error  bool
}

func (c command) String() string {
	if !c.Error {
		return fmt.Sprintf("%v\nFull Command: %v\nRunning: %v\n", c.Command, c.FullCommand, c.Running)
	} else {
		return fmt.Sprintf("%v\nFull Command: %v\nError!\n", c.Command, c.FullCommand)
	}
}

var commands []command

func exitHandler(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Completed processes:\n--------------------\n\n")
	for n, command := range commands {
		if !command.Running {
			fmt.Fprintf(w, "(%d): \t%v \nError: %v\n\n", n, command.Command, command.Error)
		}
	}
}

func runningHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Running processes:\n------------------\n\n")
	for _, command := range commands {
		if command.Running {
			fmt.Fprintf(w, "%v\n%v\n\n", command.Command, command.FullCommand)
		}
	}
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	commandNumber, _ := strconv.Atoi(r.URL.Path[len("/_result/"):])
	if commandNumber < len(commands) && commandNumber >= 0 {
		fmt.Fprintf(w, commands[commandNumber].Result)
	}
}

func execHandler(w http.ResponseWriter, r *http.Request) {
	args := strings.Split(r.URL.Path[1:], " ")

	place := len(commands)
	commands = append(commands, command{Command: args[0], FullCommand: args, Running: true})

	// args[0] is the command, args[1:] are the options.
	// e.g. ping www.google.de : args = ["ping", "www.google.de"]
	out, err := exec.Command(args[0], args[1:]...).Output()
	commands[place].Result = string(out)

	if err != nil {
		commands[place].Error = true
		commands[place].Result = "Command did not execute correctly"
	}

	commands[place].Running = false
	fmt.Fprint(w, commands[place].Result)
}

func main() {
	ips, _ := net.InterfaceAddrs()

	for i, ip := range ips {
		fmt.Printf("IP %d:\t %v:12345\n", i, ip)

		broadcastIP := broadcast.IpToBroadcast(ip.String())
		broadcast.Broadcast(broadcastIP + ":12345")
	}

	http.HandleFunc("/", execHandler)
	http.HandleFunc("/_info/", infoHandler)
	http.HandleFunc("/_result/", resultHandler)
	http.HandleFunc("/_running/", runningHandler)
	http.HandleFunc("/exit/", exitHandler)
	http.HandleFunc("/favicon.ico", nil)

	http.ListenAndServe(":12345", nil)

	fmt.Scanln()
}
