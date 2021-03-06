package apis

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/shubhamagarwal003/go-blog/helper"
	"github.com/shubhamagarwal003/go-blog/models"
	"net/http"
	"reflect"
	"strings"
)

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	if user == nil {
		http.Redirect(w, r, "/assets/login.html", http.StatusFound)
	} else {
		err := r.ParseForm()
		if err != nil {
			fmt.Println(fmt.Errorf("Error: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userValue := reflect.ValueOf(user).Elem()
		uid := userValue.FieldByName("Id").Int()
		blog := models.Blog{}
		blog.Title = r.Form.Get("title")
		blog.Content = r.Form.Get("content")
		blog.Uid = uid
		tag := r.Form.Get("tag")
		tx, err := helper.Str.BeginTransaction()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer func() {
			if err != nil {
				fmt.Println("... Error")
				tx.Rollback()
				return
			}
			err = tx.Commit()
		}()
		blogId, err := helper.Str.CreateBlog(&blog)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var tempTag *models.Tag
		if len(tag) > 0 && blogId > 0 {
			tagValues := strings.Split(tag, ",")
			for _, tagValue := range tagValues {
				tempTag = &models.Tag{Value: tagValue, BlogId: blogId}
				err = helper.Str.CreateTag(tempTag)
				if err != nil {
					fmt.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func GetBlogs(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	if user == nil {
		http.Redirect(w, r, "/assets/login.html", http.StatusFound)
		return
	}

	// Everything else is the same as before
	blogs, err := helper.Str.GetBlogs()
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

func GetBlog(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	if user == nil {
		http.Redirect(w, r, "/assets/login.html", http.StatusFound)
		return
	}
	vars := mux.Vars(r)

	// Everything else is the same as before
	blog, err := helper.Str.GetBlog(vars["id"])
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	blogBytes, err := json.Marshal(blog)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(blogBytes)
}
