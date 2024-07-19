package main

import (
	"example/zota_assignment/handler"
	"net/http"
)

func main() {
	handler := handler.NewHandler()
	http.Handle("/", handler)
	http.ListenAndServe(":8080", nil)
}
