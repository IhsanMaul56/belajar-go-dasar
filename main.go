package main

// Source untuk hello world
// import "fmt"

// func main() {
//     fmt.Println("Aku Tampan!")
// }

// Source untuk route
// import (
//     "log"
//     "net/http"
// )

// func home(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello dari Aku Tampan"))
// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", home)

// 	log.Print("Starting serve on: 4000")

// 	err := http.ListenAndServe(":4000", mux)
// 	log.Fatal(err)
// }

// Berbeda route
import (
    "log"
    "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Ini adalah HOME page"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Ini adalah view snippest"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Ini adalah create snippest"))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/{$}", home)
    mux.HandleFunc("/snippet/view", snippetView)
    mux.HandleFunc("/snippet/create", snippetCreate)

    log.Print("starting server on :4000")

    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}