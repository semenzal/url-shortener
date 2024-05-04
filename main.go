package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type User struct {
	Name      string `json:"name"`
	Last_Name string `json:"last_name"`
	Age       int    `json:"age"`
	Verified  bool   `json:"verified"`
}


func main() {
	userMap := make(map[string]User)

	r := chi.NewRouter()

	// Get ручка
	r.Get("/api/user", func(w http.ResponseWriter, r *http.Request) {
		uuid := chi.URLParam(r, "uuid")
		
		if _, ok := userMap[uuid]; !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json, err := json.Marshal(userMap[uuid])
		if err != nil {
			fmt.Println(err)
		}
		w.Write(json)
	})

	// Post ручка
	r.Post("/api/user", func(w http.ResponseWriter, r *http.Request) {
		var user User
		uuid := uuid.NewString()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		
		err = json.Unmarshal([]byte(body), &user)
		if err != nil {
			fmt.Println(err)
		}
		userMap[uuid] = user
		w.Write([]byte(uuid))
	})

	// Delete ручка
	r.Delete("/api/user", func(w http.ResponseWriter, r *http.Request) {
		uuid := chi.URLParam(r, "")

		delete(userMap, uuid)
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}
}
