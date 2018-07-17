package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/shubhamagarwal003/blog/apis"
	"github.com/shubhamagarwal003/blog/helper"
	"github.com/shubhamagarwal003/blog/middlewares"
	"net/http"
	"reflect"
)

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	// r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/register", apis.CreateUser).Methods("POST")
	r.HandleFunc("/login", apis.LoginUser).Methods("POST")
	r.HandleFunc("/logout", apis.LogoutUser).Methods("GET")
	r.HandleFunc("/", middlewares.UserLogged(homePage)).Methods("GET")
	r.HandleFunc("/blog", middlewares.UserLogged(apis.CreateBlog)).Methods("POST")
	r.HandleFunc("/blogs", middlewares.UserLogged(apis.GetBlogs)).Methods("GET")
	r.HandleFunc("/blog/{id:[0-9]+}", middlewares.UserLogged(apis.GetBlog)).Methods("GET")
	return r
}

func homePage(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	userValue := reflect.ValueOf(user).Elem()
	u := userValue.FieldByName("Id").Int()
	fmt.Println(userValue, u)
}

func main() {
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	connString := "dbname=bird_encyclopedia sslmode=disable user=shubham password=qwerty1234"
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	helper.InitStore(&helper.DbStore{Db: db})
	r := newRouter()
	fmt.Println("Server Starting")
	http.ListenAndServe(":8080", r)
}
