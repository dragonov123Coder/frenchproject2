package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Serve static files (HTML, CSS, JS) from ./static folder
	var fs http.Handler = http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Use Railway's PORT environment variable, fallback to 8080 locally
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}
