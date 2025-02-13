package handlers

import (
	"context"
	"infosphere-backend/internal/database"
	"infosphere-backend/pkg/utils"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
)

func SignupHandler(db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		user := r.Form.Get("username")
		email := r.Form.Get("email")
		pw := r.Form.Get("password")

		hashedPassword, err := utils.HashPassword(pw, utils.DefaultAPParams)
		if err != nil {
			apiErr := utils.NewAPIError(http.StatusInternalServerError, "Internal Server Error")
			apiErr.SendAPIErrorResponse(w)
			return
		}

		err = db.CreateUser(context.Background(), database.CreateUserParams{
			Username:     user,
			Email:        email,
			PasswordHash: hashedPassword,
		})

		if err != nil {
			if pgErr, ok := err.(*pgconn.PgError); ok {

				var statusCode int
				var msg string

				switch pgErr.Code {
				case "23505":
					statusCode = http.StatusConflict
					msg = "Username/Email Already In Use"
				}

				apiErr := utils.NewAPIError(statusCode, msg)
				apiErr.SendAPIErrorResponse(w)
			} else {
				apiErr := utils.NewAPIError(http.StatusInternalServerError, "Internal Server Error")
				apiErr.SendAPIErrorResponse(w)
			}

			return
		}

		jwt, err := utils.GenerateJWT(user)
		if err != nil {
			apiErr := utils.NewAPIError(http.StatusInternalServerError, "Internal Server Error")
			apiErr.SendAPIErrorResponse(w)
			return
		}

		utils.SetAuthCookie(&w, jwt)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Account successfully created"))
	})

}
