package main

import (
	"fmt"
	"net/http"
	"whtest"
)

func main() {
	whtest.SetupRouter()

	fmt.Println("Server listening on :80")
	http.ListenAndServe(":80", nil)
}
