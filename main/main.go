package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	// Declare the static file directory and point it to the directory we just made
	staticFileDirectory := http.Dir("./blogs/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/blogs/" prefix when looking for files.
	// For example, if we type "/blogs/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for "./blogs/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/blogs", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/blogs/", instead of the absolute route itself
	r.PathPrefix("/blogs/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/blog", getBlogHandler).Methods("GET")
	r.HandleFunc("/blog", createBlogHandler).Methods("POST")
	return r
}

func main() {
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	// fmt.Fprintf(w, "Starting Application")
	r := newRouter()
	err := http.ListenAndServe(":8081", r)
	// fmt.Fprintf(w, err)
	if err != nil {
		panic(err.Error())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
