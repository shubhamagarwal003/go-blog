package apis

import (
	"crypto/rand"
	"fmt"
	"github.com/shubhamagarwal003/blog/helper"
	"github.com/shubhamagarwal003/blog/models"
	"net/http"
	"time"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.Username = r.Form.Get("username")
	user.Password = r.Form.Get("password")
	// The only change we made here is to use the `CreateBird` method instead of
	// appending to the `bird` variable like we did earlier
	err = helper.Str.CreateUser(&user)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/assets/login.html", http.StatusFound)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.Username = r.Form.Get("username")
	user.Password = r.Form.Get("password")
	userExists, userId := helper.Str.CheckUser(&user)

	if userExists {
		expires := time.Now().AddDate(1, 0, 0)
		token := tokenGenerator()
		fmt.Println(token)
		cookie := &http.Cookie{Name: "SessionId", Value: token, HttpOnly: false, Expires: expires}
		sessionToken := models.SessionToken{Token: token, Uid: userId}
		err = helper.Str.SaveToken(&sessionToken)
		if err != nil {
			http.Redirect(w, r, "/assets/login.html", http.StatusFound)
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/assets/login.html", http.StatusFound)
	}
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("SessionId")
	if err == nil {
		helper.Str.DeleteToken(cookie.Value)
	}
	expires := time.Now().AddDate(-1, 0, 0)
	cookie = &http.Cookie{Name: "SessionId", Value: "", HttpOnly: false, Expires: expires}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/assets/login.html", http.StatusFound)
}

func tokenGenerator() string {
	b := make([]byte, 20)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
