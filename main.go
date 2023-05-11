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
		writer.Header().Set("Content-Type", "text/html")
		writer.WriteHeader(http.StatusOK)
		host := request.Host
		fmt.Fprintf(writer, `
Welcome to GitLab workspace demo Go HTTP app! <br/><br/>
You can browse <br/>
- <a href="https://%[1]s/text">https://%[1]s/text</a> for Textual response <br/>
- <a href="https://%[1]s/json">https://%[1]s/json</a> for JSON response <br/>
`, host)
	}
}

func handleTextEndpoint() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html")
		response := ""
		for k, v := range request.Header {
			response += fmt.Sprintf("%s : %v <br/>", k, v)
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(response))
	}
}

func handleJsonEndpoint() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		headers := make(map[string]interface{})
		for k, v := range request.Header {
			headers[k] = v
		}

		response, err := json.MarshalIndent(headers, "", "  ")
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusOK)
		writer.Write(response)
	}
}
