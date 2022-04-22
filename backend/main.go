package main

import (
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
  return colour
}

func add(x, y int) int {
  return x + y
}

func main() {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}