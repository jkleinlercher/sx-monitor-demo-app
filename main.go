package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/site/200", handle200)
	http.HandleFunc("/site/503", handle503)
	http.HandleFunc("/site/delay", handleDelay)
	http.HandleFunc("/site/high-cpu", handleHighCPU)
	http.HandleFunc("/site/high-memory", handleHighMemory)
	http.HandleFunc("/site/high-bandwidth", handleHighBandwidth)
	fmt.Println("Listening on :8080...")
	http.ListenAndServe(":8080", nil)
}

func handle200(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func handle503(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusServiceUnavailable)
	w.Write([]byte("Service Unavailable"))
}

func handleDelay(w http.ResponseWriter, r *http.Request) {
	sleepParam := r.URL.Query().Get("sleep")
	var sleep time.Duration = 0
	if sleepParam != "" {
		sleepInt, err := strconv.Atoi(sleepParam)
		if err == nil {
			sleep = time.Duration(sleepInt) * time.Second
		}
	}
	time.Sleep(sleep)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Response delayed by %d seconds", sleep/time.Second)))
}

func handleHighCPU(w http.ResponseWriter, r *http.Request) {
	iterations := 1000000000
	result := 0
	for i := 0; i < iterations; i++ {
		result += i
	}
	fmt.Fprintf(w, "Result: %d\n", result)
}

func handleHighMemory(w http.ResponseWriter, r *http.Request) {
	n := 100000000
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}
	fmt.Fprintln(w, "Allocated memory:", n*4, "bytes")
}

func handleHighBandwidth(w http.ResponseWriter, r *http.Request) {
	data := make([]byte, 1024*1024*100)
	for i := 0; i < len(data); i++ {
		data[i] = byte(i % 256)
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
	w.Write(data)
}
