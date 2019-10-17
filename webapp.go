package main

import ("fmt"
				"net/http"
				)

// index_handler takes in two paramaters, response and request
func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "=============================")
	fmt.Fprintf(w,   "[ <h1>GO is awesome!</h1> ]")
	fmt.Fprintf(w, "=============================")
}

// firstgo_handler takes in two paramaters, response and request
func firstgo_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "==================================")
	fmt.Fprintf(w, "[Whoa, Check this out...awesome!]")
	fmt.Fprintf(w, "==================================")
}

// main ()is the only function that will run in the program and listens in local host
func main() {
		http.HandleFunc("/", index_handler)
		http.HandleFunc("/firstgo", firstgo_handler)
		http.ListenAndServe(":8000", nil)
}
