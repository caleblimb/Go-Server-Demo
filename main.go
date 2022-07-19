package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Handle simple staic page
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Check if page exists
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Check for server errors
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	// Print out contents of page
	fmt.Fprintf(w, "hello!")
}

// Handle Form input
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Print out response
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	age := r.FormValue("age")
	address := r.FormValue("address")

	// Conver Age to Int for comparison
	i, err := strconv.Atoi(age)
	if err != nil {
		fmt.Printf("Server error. Please try again later")
	}

	// Check if age is at least 18
	if i < 18 {
		fmt.Fprintf(w, "You must be 18 or older, %s is not old enough!\n", age)
	} else { // Print rest of response
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Age = %s\n", age)
		fmt.Fprintf(w, "Address = %s\n", address)
	}
}

func main() {
	// Start the Server
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// Indicate that the server is starting
	fmt.Printf("Starting Server at port 8080\n")

	// Wait for user action
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
