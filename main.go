package main //como se fosse o namespace

import (
	"api-teste/routers"
	"fmt"

	api "github.com/valyala/fasthttp" //servidor http - bem mais rápida que a nativa
)

func main() {
	fmt.Println("Hello World")

	r := routers.Get()
	err := api.ListenAndServe(":8080", r) //função que sobe o serviço de HTTP

	if err != nil {
		fmt.Println(
			fmt.Printf("An inexpected error happened: %s", err),
		)
	}
}
