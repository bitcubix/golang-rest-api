package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
)

type Book struct {
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}

type BookHandlers struct {
	sync.Mutex
	store map[string]Book
}

func (h *BookHandlers) books(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	case http.MethodPost:
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

func (h *BookHandlers) get(w http.ResponseWriter, r *http.Request) {
	books := make([]Book, len(h.store))

	h.Lock()
	i := 0
	for _, book := range h.store {
		books[i] = book
		i++
	}
	h.Unlock()

	jsonBytes, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *BookHandlers) post(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if r.Header.Get("content-type") != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte("need content-type 'application/json'"))
		return
	}

	var book Book
	err = json.Unmarshal(bodyBytes, &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	h.Lock()
	h.store[book.ISBN] = book
	defer h.Unlock()
}

func newBookHandlers() *BookHandlers {
	return &BookHandlers{
		store: map[string]Book{
			"id1": {
				ISBN:   "1234",
				Title:  "title1",
				Author: "author1",
				Price:  10.1,
			},
		},
	}
}

func main() {
	bookHandlers := newBookHandlers()
	http.HandleFunc("/books", bookHandlers.books)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
