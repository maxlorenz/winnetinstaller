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

var htmlHeader string = `
<html><head><meta http-equiv="content-type" content="text/html;charset=utf-8">
<meta http-equiv="refresh" content="3" ><style type="text/css">html{
	font-family: "Consolas";font-size: 12;background-color: black;color: white;white-space: pre;
	margin: 1em;}a:link, a:visited, a:hover, a:active{color: orange;}.green{color: green;}
	.orange{color: orange;}.red{color: red;}input{width: 80%;}</style></head><body>
	<br><br>
`

type command struct {
	Command, Result string
	FullCommand     []string
	Running, Error  bool
}

var commands []command = make([]command, 0)

func exitHandler(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

func htmlInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlHeader)
	for i := len(commands) - 1; i >= 0; i-- {
		command := commands[i]
		fmt.Fprintf(w, "<p href=\"/test\" class=\"")
		if command.Running {
			fmt.Fprintf(w, "orange")
		} else if command.Error {
			fmt.Fprintf(w, "red")
		} else {
			fmt.Fprintf(w, "green")
		}
		fmt.Fprintf(w, "\">"+command.Command+"</p>")
		fmt.Fprintf(w, "<small>"+strings.Join(command.FullCommand, " "))
		fmt.Fprintf(w, " | <a href=\"/html/result/"+strconv.Itoa(i)+"\">Result</a></small><br><br>")
	}
}

func plainInfoHandler(w http.ResponseWriter, r *http.Request) {
	for i := len(commands) - 1; i >= 0; i-- {
		command := commands[i]
		if command.Running {
			fmt.Fprintf(w, "Running: ")
		} else if command.Error {
			fmt.Fprintf(w, "Error: ")
		} else {
			fmt.Fprintf(w, "Successful (" + strconv.Itoa(i) + "): ")
		}
		fmt.Fprintf(w, command.Command + "\n")
		fmt.Fprintf(w, "%s\n\n", command.FullCommand)
	}
}

func htmlResultHandler(w http.ResponseWriter, r *http.Request) {
	commandNumber, _ := strconv.Atoi(r.URL.Path[len("/html/result/"):])
	if commandNumber < len(commands) && commandNumber >= 0 {
		fmt.Fprintf(w, htmlHeader)
		fmt.Fprintf(w, commands[commandNumber].Result)
	}
}

func plainResultHandler(w http.ResponseWriter, r *http.Request) {
	commandNumber, _ := strconv.Atoi(r.URL.Path[len("/plain/result/"):])
	if commandNumber < len(commands) && commandNumber >= 0 {
		fmt.Fprintf(w, commands[commandNumber].Result)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	args := strings.Split(r.URL.Path[1:], " ")

	place := len(commands)
	commands = append(commands, command{Command: args[0], FullCommand: args, Running: true})

	// args[0] is the command, args[1:] are the options.
	// e.g. ping www.google.de : args = ["ping", "www.google.de"]
	out, err := exec.Command(args[0], args[1:]...).Output()

	if err != nil {
		commands[place].Result = "Command did not execute correctly"
		commands[place].Error = true
	} else {
		commands[place].Result = string(out)
	}

	commands[place].Running = false
	fmt.Fprintf(w, commands[place].Result)
}

func main() {
	fmt.Println(net.InterfaceAddrs())

	http.HandleFunc("/exit/", exitHandler)
	http.HandleFunc("/html/info/", htmlInfoHandler)
	http.HandleFunc("/plain/info/", plainInfoHandler)
	http.HandleFunc("/html/result/", htmlResultHandler)
	http.HandleFunc("/plain/result/", plainResultHandler)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":12345", nil)

	fmt.Scanln()
}
