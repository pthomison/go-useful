package utilkit

import "net/http"

func Redirect(w http.ResponseWriter, location string) {
	w.Header().Add("Location", location)
	w.WriteHeader(http.StatusFound)
}
