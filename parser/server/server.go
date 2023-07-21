package server

import (
	"encoding/json"
	"fmt"
	"github.com/InstIDEA/ddjj/parser/extract"
	"net/http"
	"os"
	"path"
	"strings"
)

type FileRequest struct {
	File struct {
		Path string `json:"path"`
	} `json:"file"`
}

func handleFile(filePath string) extract.ParserData {

	if strings.HasPrefix("http", filePath) {
		return extract.CreateError("Downloading files is not yet supported")
	}

	var finalPath = path.Join(getEnv("PARSER_DEFAULT_DIR", "."), filePath)
	dat, err := os.Open(finalPath)
	fmt.Printf("Parsing %s\n", finalPath)

	if err != nil {
		fmt.Println("The file can't be found", err)
		return extract.CreateError(fmt.Sprint("File ", filePath, " not found. "))
	}

	return extract.ParsePDF(dat)

}

func InitServer() {
	// Define the route handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is POST
		if r.Method == http.MethodPost {
			// Parse the POST data (assuming it's a JSON object)
			var request FileRequest
			err := json.NewDecoder(r.Body).Decode(&request)
			if err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			// Convert the response to JSON format
			var response = handleFile(request.File.Path)
			jsonResponse, err := json.Marshal(response)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			if response.Status != 0 {
				w.WriteHeader(400)
			}

			// Set the Content-Type header to application/json
			w.Header().Set("Content-Type", "application/json")

			// Write the JSON response back to the client
			w.Write(jsonResponse)
		} else {
			// Handle other HTTP methods (GET, PUT, DELETE, etc.)
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}

	// Register the route handler function with the http package
	http.HandleFunc("/", handler)

	// Start the server and listen on port 8080
	var port = getEnv("PARSER_PORT", "8080")
	fmt.Printf("Server is running on http://localhost:%s/\n", port)
	fmt.Printf("Serving files from '%s'\n", os.Getenv("PARSER_DEFAULT_DIR"))
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
