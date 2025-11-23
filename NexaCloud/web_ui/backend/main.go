package main

import (
	"fmt"
	"net/http"
	"nexacloud/web_ui/handlers"
)

func main() {
	// API routes for package management MUST be registered before the file server.
	http.HandleFunc("/api/packages", handlers.GetPackages)
	http.HandleFunc("/api/packages/install", handlers.InstallPackage)
	http.HandleFunc("/api/packages/remove", handlers.RemovePackage)

	// Serve the frontend files.
	fs := http.FileServer(http.Dir("../../frontend"))
	http.Handle("/", fs)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
