package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var items []Item
var nextID = 1

func main2() {
	http.HandleFunc("/items", itemsHandler)
	http.HandleFunc("/debug", debugHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// ----------------------
// 🔄 CRUD Handler
// ----------------------
func itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(items)

	case http.MethodPost:
		var newItem Item
		json.NewDecoder(r.Body).Decode(&newItem)
		newItem.ID = nextID
		nextID++
		items = append(items, newItem)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newItem)

	case http.MethodPut:
		var updatedItem Item
		json.NewDecoder(r.Body).Decode(&updatedItem)

		for i := range items {
			if items[i].ID == updatedItem.ID {
				items[i].Name = updatedItem.Name
				json.NewEncoder(w).Encode(items[i])
				return
			}
		}
		http.Error(w, "Item not found", http.StatusNotFound)

	case http.MethodDelete:
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)

		for i := range items {
			if items[i].ID == id {
				items = append(items[:i], items[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		http.Error(w, "Item not found", http.StatusNotFound)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// ----------------------
// 🕵️ Debug Endpoint
// ----------------------
func debugHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("---- REQUEST DEBUG ----")

	// Request Line
	fmt.Println("Method:", r.Method)
	fmt.Println("URL:", r.URL.String())
	fmt.Println("Protocol:", r.Proto)

	// Headers
	fmt.Println("\nHeaders:")
	for k, v := range r.Header {
		fmt.Printf("%s: %v\n", k, v)
	}

	// Query parameters
	fmt.Println("\nQuery Parameters:")
	for k, v := range r.URL.Query() {
		fmt.Printf("%s: %v\n", k, v)
	}

	// Cookies
	fmt.Println("\nCookies:")
	for _, cookie := range r.Cookies() {
		fmt.Printf("%s = %s\n", cookie.Name, cookie.Value)
	}

	// Body
	fmt.Println("\nBody:")
	body, _ := io.ReadAll(r.Body)
	fmt.Println(string(body))

	// Respond with a full response
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Debug", "true")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"status":"debug completed"}`)
}
