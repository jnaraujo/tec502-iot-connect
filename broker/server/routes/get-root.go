package routes

import (
	"fmt"
	"net/http"
)

func GetRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
