package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)


type Url struct {
	Url string `json:"url"`
}

func main() {

	urlMap := make(map[string]Url)

	r := chi.NewRouter()

	// Post ручка
	r.Post("/api/user", func(w http.ResponseWriter, r *http.Request) {
		var url Url

		shortUrl := generateShortUrl()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, error.Error(fmt.Errorf("not found")), http.StatusNotFound)
			return
		}

		err = json.Unmarshal([]byte(body), &url)
		if err != nil {
			http.Error(w, error.Error(fmt.Errorf("bad request")), http.StatusBadRequest)
			return
		}
		urlMap[shortUrl] = url
		fmt.Println(urlMap)
		w.Write([]byte(shortUrl))
	})

	// Get ручка
	r.Get("/api/user", func(w http.ResponseWriter, r *http.Request) {
		shortUrl := r.URL.Query().Get("shortUrl")
		if shortUrl == "" {
			http.Error(w, "Shortened key is missing", http.StatusBadRequest)
        return
		}

		if _, ok := urlMap[shortUrl]; !ok {
			http.Error(w, error.Error(fmt.Errorf("not found")), http.StatusNotFound)
			return
		}
		
		fmt.Println(urlMap)

		http.Redirect(w, r, shortUrl, http.StatusMovedPermanently)

	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func generateShortUrl() string {
	const code = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const urlLenght = 7

	rand.NewSource(time.Now().Unix())
	shortUrl := make([]byte, urlLenght)
	for i := range shortUrl {
		shortUrl[i] = code[rand.Intn(len(code))]
	}
	return string(shortUrl)
}


