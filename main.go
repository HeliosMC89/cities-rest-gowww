package main

import "github.com/heliosmc89/api-rest-gowww/api"

func main() {
	srv := api.NewServer()
	srv.Run(":8000")
}
