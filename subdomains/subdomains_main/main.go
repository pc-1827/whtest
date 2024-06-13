package main

import (
	"fmt"
	"net/http"
	"subdomains"
)

func main() {
	subdomains.SetupRouter()

	fmt.Println("Server listening on :80")
	http.ListenAndServe(":80", nil)
}
