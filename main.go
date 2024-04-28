package main

import (
	"context"
	"fmt"

	"github.com/barisdevjs/orders-api/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
