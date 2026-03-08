package main

import (
	"context"

	"github.com/1kovalevskiy/math-trainer/cmd/app"
)


func main() {
	ctx := context.Background()

	app := app.InitApp(ctx)

	app.RunApp(ctx)

}