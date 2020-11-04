package main

import (
	"network-buddy/backend/web"
	"os"
)

func main() {
	listen := os.Getenv("PORT")
	if listen == "" {
		listen = "8080"
	}
	listen = ":" + listen
	web.ListenAndServe(listen)
}
