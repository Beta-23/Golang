package main

import ("fmt"
				"net/http"
				"html/template")

type newsBlogPage struct {
	Title string
	News string

}

func newsBblogHandler(w http.ResponseWriter, r *http.Request) {
	p := newsBlogPage{Title: "My first GO Blog", News: "AL LEARNS GO!!!"}
	t, _ := template.ParseFiles("basictemplating.html")
	
	fmt.Println(t.Execute(w, p))
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "=============================")
	fmt.Fprintf(w,      "[ GO is awesome! ]      ")
	fmt.Fprintf(w, "=============================")
}

// main ()is the only function that will run in the program and listens in local host
func main() {
		http.HandleFunc("/", index_handler)
		http.HandleFunc("/blog/", newsBblogHandler)
		http.ListenAndServe(":8000", nil)
}