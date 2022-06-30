package main

import (
	"context"
	"fmt"
	"go-client/api"
	"log"
	"time"
)

func main() {
	c, err := api.NewClientWithResponses("http://localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := c.PutOrderIdWithResponse(ctx, "test", api.PutOrderIdJSONRequestBody{
		Item:  api.TeaTableGreen,
		Price: 10,
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(response.Status(), string(response.Body))
}
