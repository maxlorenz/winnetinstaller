package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func exitHandler(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

func handler(w http.ResponseWriter, r *http.Request) {
	args := strings.Split(r.URL.Path[1:], " ")

	// args[0] is the command, args[1:] are the options.
	// e.g. ping www.google.de : args = ["ping", "www.google.de"]
	out, err := exec.Command(args[0], args[1:]...).Output()

	if err != nil {
		fmt.Fprintf(w, "Command did not execute correctly")
	} else {
		fmt.Fprintf(w, "%s", string(out))
	}
}

func main() {
	fmt.Println(net.InterfaceAddrs())

	http.HandleFunc("/exit", exitHandler)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":12345", nil)

	fmt.Scanln()
}