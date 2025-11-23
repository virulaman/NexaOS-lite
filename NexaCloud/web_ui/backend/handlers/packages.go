package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// Package represents a software package.
type Package struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

// GetPackages handles API requests to list all available packages.
func GetPackages(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for /api/packages")
	// In a real implementation, this would call `nexapkg` to get the list of packages.
	packages := []Package{
		{Name: "nginx", Version: "1.21.6", Description: "A high-performance web server."},
		{Name: "docker", Version: "20.10.7", Description: "A platform for developing, shipping, and running applications in containers."},
		{Name: "python", Version: "3.9.7", Description: "An interpreted, high-level, general-purpose programming language."},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(packages)
}

// InstallPackage handles API requests to install a new package.
func InstallPackage(w http.ResponseWriter, r *http.Request) {
	// In a real implementation, this would call `nexapkg install <package_name>`.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Package installed successfully (simulation)."))
}

// RemovePackage handles API requests to remove an existing package.
func RemovePackage(w http.ResponseWriter, r *http.Request) {
	// In a real implementation, this would call `nexapkg remove <package_name>`.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Package removed successfully (simulation)."))
}
