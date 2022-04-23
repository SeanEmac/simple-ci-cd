package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
  "encoding/json"
)

type Response struct {
  Colour string
  Greeting string
}

// Workaround to stop cors being blocked
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("Got Request\n")
  enableCors(&w)

  data := Response{getColour(), r.URL.Path[1:]}

  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  json.NewEncoder(w).Encode(data)
}

func getColour() string {
  colour, exists := os.LookupEnv("COLOUR")
  if !exists {
    colour = "PINK"
  }
  fmt.Printf("Colour is %s\n", colour)
  return colour
}

func main() {
  fmt.Printf("Starting Webserver\n")
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}