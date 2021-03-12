package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Blog struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var blogs []Blog

func getBlogHandler(w http.ResponseWriter, r *http.Request) {
	//Convert the "birds" variable to json
	blogListBytes, err := json.Marshal(blogs)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of birds to the response
	w.Write(blogListBytes)
}

func createBlogHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of Bird
	blog := Blog{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the bird from the form info
	blog.Title = r.Form.Get("title")
	blog.Description = r.Form.Get("description")

	// Append our existing list of birds with a new entry
	blogs = append(blogs, blog)

	//Finally, we redirect the user to the original HTMl page (located at `/assets/`)
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
