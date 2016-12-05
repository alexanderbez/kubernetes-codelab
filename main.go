// Kubernetes codelab tutorial
//
// @author: Aleksandr Bezobchuk
// @date: 12/04/2016
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Response is a dummy struct that is used for an API response.
type Response struct {
	Msg       string `json:"msg"`
	Timestamp string `json:"timestamp"`
}

// TODO
func HomeRoute(rw http.ResponseWriter, r *http.Request) {
	resp := &Response{Msg: "Hello, world", Timestamp: time.Now().Format(time.RFC822)}
	body, _ := json.MarshalIndent(resp, "", "    ")

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(body)
}

func main() {
	p := 8080

	fmt.Println("Booting server...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("Server shutting down...")
		os.Exit(1)
	}()

	defer func() {
		http.HandleFunc("/", HomeRoute)
		http.ListenAndServe(fmt.Sprintf(":%d", p), nil)
	}()

	fmt.Println("Server running port", p)
}
