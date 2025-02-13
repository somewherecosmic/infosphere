package handlers

import (
	"context"
	"infosphere-backend/internal/database"
	"infosphere-backend/pkg/utils"
	"log"
	"net/http"
)

func LoginHandler(db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		userIdentity := r.Form.Get("user-identity")
		passwordAttempt := r.Form.Get("password")

		userDetails, err := db.FindUserByHandle(context.Background(), userIdentity)
		if err != nil {
			apiErr := utils.NewAPIError(http.StatusUnauthorized, "Incorrect credentials. Invalid username/email and/or password")
			log.Println(err.Error())
			apiErr.SendAPIErrorResponse(w)
			return
		}

		match, err := utils.VerifyPassword(passwordAttempt, userDetails.PasswordHash)
		if err != nil {
			apiErr := utils.NewAPIError(http.StatusInternalServerError, err.Error())
			apiErr.SendAPIErrorResponse(w)
			return
		}

		if !match {
			apiErr := utils.NewAPIError(http.StatusUnauthorized, "Incorrect credentials. Invalid username/email and/or password")
			apiErr.SendAPIErrorResponse(w)
			return
		}

		token, err := utils.GenerateJWT(userDetails.Email)
		if err != nil {
		}

		utils.SetAuthCookie(&w, token)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
	})
}
