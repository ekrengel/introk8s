package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/kubernetes", kubernetes)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello ACT-W!")
}

func kubernetes(w http.ResponseWriter, r *http.Request) {
	podName := os.Getenv("POD_NAME")
	nodeName := os.Getenv("NODE_NAME")

	if podName == "" {
		podName = "UNKNOWN"
	}
	fmt.Fprintf(w, "Pod name: %s\n", podName)

	if nodeName == "" {
		nodeName = "UNKNOWN"
	}
	fmt.Fprintf(w, "Node name: %s\n", nodeName)
}