package apis

import (
	"encoding/json"
	"fmt"
	"github.com/shubhamagarwal003/go-blog/helper"
	"net/http"
)

func GetBlogsWithTag(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	if user == nil {
		http.Redirect(w, r, "/assets/login.html", http.StatusFound)
		return
	}

	tags, ok := r.URL.Query()["tag"]

	if !ok || len(tags[0]) < 1 {
		fmt.Println("Url Param 'tag' is missing")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	tag := tags[0]
	blogs, err := helper.Str.GetBlogsForTag(tag)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	blogBytes, err := json.Marshal(blogs)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(blogBytes)
}
