package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var number uint64 = 0

// To test this, run: ab -n 10000 -c 100 http://localhost:3000/ (must have ab installed)
func main() {
	// Define an HTTP handler that increments the visit counter and returns the count
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Increment the counter using atomic operations for thread safety
		atomic.AddUint64(&number, 1)

		// Write the response with the current count
		w.Write([]byte(fmt.Sprintf("You have accessed this page %d times.", number)))
	})

	// Start the HTTP server on port 3000
	http.ListenAndServe(":3000", nil)
}
