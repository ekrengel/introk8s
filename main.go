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
	podName := getEnvVar("POD_NAME")
	nodeName := getEnvVar("NODE_NAME")

	fmt.Fprintf(w, "Pod name: %s\n", podName)
	fmt.Fprintf(w, "Node name: %s\n", nodeName)
}

func getEnvVar(envVar string) string {
	v := os.Getenv(envVar)
	if v == "" {
		return "UNKNOWN"
	}
	return v
}