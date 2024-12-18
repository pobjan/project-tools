package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "image.jpg")  // Zakładając, że masz plik 'image.jpg' w tym samym katalogu
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Serwer uruchomiony na porcie 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
