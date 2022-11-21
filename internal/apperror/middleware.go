package apperror

import "net/http"

func Middleware(h ahandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
