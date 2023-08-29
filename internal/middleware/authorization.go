package middleware

import (
	"errors"
	"net/http"

	"github.com/leetcode-golang-classroom/go-api/api"
	"github.com/leetcode-golang-classroom/go-api/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error
		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrrorHandler(w, UnAuthorizedError)
			return
		}
		var database *tools.DatabaseInterface
		database, err = tools.NewDatebase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}
		loginDetails := (*database).GetUserLoginDetails(username)
		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrrorHandler(w, UnAuthorizedError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
