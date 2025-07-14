package utils

import (
	"fmt"
	"net/http"
)

func RenderError(w http.ResponseWriter, statuscode int, msg string) {
	w.WriteHeader(statuscode)
	fmt.Fprintf(w, "%d - %s\n", statuscode, msg)
}
