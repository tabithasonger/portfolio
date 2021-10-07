package controllers

import(
	"portfolio/views"
	"net/http"
	"fmt"
	// "github.com/gorilla/schema"
)

//used  to create new users controller. panic if err
func NewUsers() *Users{
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.html"),
	}
}

type Users struct {
	NewView *views.View
}

type SignupForm struct {
	Email 	string	`schema:"email"`
	Password string	`schema:"password"`
}

// New is used to render the form
// GET/signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}
// Create is used to process the signup form when user sumbits.
// New user account
// POST/signup
func (u *Users) Create(w http.ResponseWriter,  r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}