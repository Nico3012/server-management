package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	dir := os.Getenv("PUBLIC_DIR")
	if dir == "" {
		dir = "/var/www/public"
	}

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)

	log.Printf("Starte Webserver für %s auf Port 80...", dir)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
