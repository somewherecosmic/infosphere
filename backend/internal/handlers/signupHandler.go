package handlers

import (
	"fmt"
	"infosphere-backend/pkg/utils"
	"net/http"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := r.Form.Get("username")
	email := r.Form.Get("email")
	pw := r.Form.Get("password")

	hashedPassword, err := utils.HashPassword(pw, utils.DefaultAPParams)
	if err != nil {
		apiErr := utils.NewAPIError(500, "Internal Server Error")
		apiErr.SendAPIErrorResponse(w)
		return
	}

	// TODO
	// create user row within DB table
	// validation i.e does user already exist etc.

	fmt.Printf("%s, %s, %s\n", user, email, pw)
	w.Write([]byte("Reached signup"))
}
