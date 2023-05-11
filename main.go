package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRootEndpoint())
	http.HandleFunc("/text", handleTextEndpoint())
	http.HandleFunc("/json", handleJsonEndpoint())

	port := "3000"
	log.Printf("Starting app on port: %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleRootEndpoint() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		host := request.Host
		fmt.Fprintf(writer, `
Welcome to GitLab workspace demo Go HTTP app! <br/><br/>
- https://%s/text for Textual response
- https://%s/json for JSON response
`, host, host)
	}
}

func handleTextEndpoint() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprint(writer, "Hello from Starter Go App")
	}
}

func handleJsonEndpoint() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		response := make(map[string]string)
		response["status"] = "ok"

		raw, err := json.Marshal(response)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusOK)
		writer.Write(raw)
	}
}
