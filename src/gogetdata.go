// This file contains code for a simple static data web server
//
// Author: Josh McIntyre
//

package main

import (
	"fmt"
	"net/http"
	"os"
)

// Handle a simple request by reading a file and displaying contents
func request_handler(writer http.ResponseWriter, request *http.Request) {

	// RequestURI gives the URI after the host, including the root slash
	// ex: http://localhost/data.html is returned as /data.html
	// So, trim the / off using a slice to get the right filename
	file_requested := request.RequestURI[1:]
	fmt.Printf("%s\n", file_requested)

	// Read the requested file
	content, err := os.ReadFile(file_requested)
	if err != nil {
		// Write a simple error message if unable to read
		fmt.Fprintf(writer, "Unable to read requested file")
		return
	}
	
	// Write the data to the ResponseWriter
	fmt.Fprintf(writer, string(content))

}

// The main entry point for the program
func main() {
	
	// Define URL handlers
	http.HandleFunc("/", request_handler)

	// Start the web server
	http.ListenAndServe(":8080", nil)
}