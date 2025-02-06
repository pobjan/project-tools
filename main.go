// main.go
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("./image.jpg")
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		defer file.Close()

		// Pobranie informacji o pliku
		info, err := file.Stat()
		if err != nil {
			http.Error(w, "Could not get file info", http.StatusInternalServerError)
			return
		}

		// Ustawienie nagłówków
		w.Header().Set("Content-Length", fmt.Sprintf("%d", info.Size()))
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Date", time.Now().Format(http.TimeFormat))

		// Wysłanie pliku jako odpowiedź
		http.ServeFile(w, r, "./image.jpg")
	})

	http.ListenAndServe(":8080", nil)
}
