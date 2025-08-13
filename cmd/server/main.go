package main

import "go-simple-api/internal/routers"

func main() {
	r := routers.NewRouter()

	err := r.Run(":8080")

	if err != nil {
		panic(err)
	}
}
