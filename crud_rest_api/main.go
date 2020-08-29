package main

import (
  "encoding/json"
  "log"
  "net/http"
  "math/rand"
  "strconv"
  "github.com/gorilla/mux"
)

//MODELS
type Book struct {
  Id string `json:"id"`
  Isbn string `json:"isbn"`
  Title string `json:"title"`
  Author *Author `json:"author"`
}

type Author struct {
  Firstname string `json:"firstname"`
  Lastname string `json:"lastname"`
}

//slice of books
var books []Book


//FUNCTIONS
func getBooks(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r) //get params

  for _, item := range books {
    if item.Id == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }

  json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  var book Book
  _ = json.NewDecoder(r.Body).Decode(&book)
  book.Id = strconv.Itoa(rand.Intn(1000000))
  books = append(books, book)

  json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r) //get params

  for i, item := range books {
    if item.Id == params["id"] {
      books = append(books[:i], books[i+1:]...)
      var book Book
      _ = json.NewDecoder(r.Body).Decode(&book)
      book.Id = params["id"]
      books = append(books, book)

      json.NewEncoder(w).Encode(book)
      return
    }
  }
  json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r) //get params

  for i, item := range books {
    if item.Id == params["id"] {
      books = append(books[:i], books[i+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(books)
}

func main() {
  //init router
  r := mux.NewRouter()

  //mock data
  books = append(books, Book{Id: "1", Isbn: "3423423", Title:"Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
  books = append(books, Book{Id: "2", Isbn: "5346345", Title:"Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
  books = append(books, Book{Id: "3", Isbn: "5233533", Title:"Book Three", Author: &Author{Firstname: "Foo", Lastname: "Bar"}})
  //route handlers
  r.HandleFunc("/api/books", getBooks).Methods("GET")
  r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
  r.HandleFunc("/api/books", createBook).Methods("POST")
  r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
  r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8000", r))
}
