package main

import (
	"context"

	"github.com/1kovalevskiy/math-trainer/cmd/app"
)

func main() {
	ctx := context.Background()

	application, err := app.InitApp(ctx)
	if err != nil {
		panic(err)
	}

	if err := application.RunApp(ctx); err != nil {
		panic(err)
	}
}
