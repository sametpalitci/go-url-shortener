package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"

	db "url-shortener/database"
	"url-shortener/utils"

	"github.com/joho/godotenv"
)

var (
	port = utils.GetEnvVar("PORT", "8080")
	host = utils.GetEnvVar("HOST", "localhost")
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found")
	}
}

func generateShortURL() (string, error) {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b)[:6], nil
}

func handleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortURL, err := generateShortURL()
	if err != nil {
		http.Error(w, "Error generating short URL", http.StatusInternalServerError)
		return
	}

	urlDoc := db.URL{
		ShortURL:    shortURL,
		OriginalURL: originalURL,
		CreatedAt:   time.Now(),
	}

	err = db.SaveURL(urlDoc)
	if err != nil {
		log.Printf("Error saving URL: %v", err)
		http.Error(w, "Error saving URL", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Shortened URL: http://%s:%s/%s", host, port, shortURL)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]
	if shortURL == "" {
		http.Error(w, "Short URL is required", http.StatusBadRequest)
		return
	}

	urlDoc, err := db.GetURL(shortURL)
	if err != nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, urlDoc.OriginalURL, http.StatusMovedPermanently)
}

func main() {
	db.ConnectDB()
	defer db.CloseDB()

	http.HandleFunc("/create", handleShorten)
	http.HandleFunc("/", handleRedirect)

	fmt.Printf("URL Shortener | Server starting on http://%s:%s\n", host, port)
	http.ListenAndServe(":"+port, nil)
}
