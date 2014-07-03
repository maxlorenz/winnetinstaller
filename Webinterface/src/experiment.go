package main

import (
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

var commands []command

func exitHandler(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {

	for n, command := range commands {
		switch command.Running {
		case true:
			fmt.Fprintf(w, "(%d) Running:\t %v\n", n, command.Command)
		case false:
			fmt.Fprintf(w, "(%d) Completed:\t %v\n", n, command.Command)
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
	out, _ := exec.Command(args[0], args[1:]...).Output()
	commands[place].Result = string(out)

	commands[place].Running = false
	fmt.Fprintf(w, commands[place].Result)
}

func main() {
	fmt.Println(net.InterfaceAddrs())

	http.HandleFunc("/", execHandler)
	http.HandleFunc("/_info/", infoHandler)
	http.HandleFunc("/_result/", resultHandler)
	http.HandleFunc("/exit/", exitHandler)
	http.HandleFunc("/favicon.ico", nil)

	http.ListenAndServe(":12345", nil)

	fmt.Scanln()
}
