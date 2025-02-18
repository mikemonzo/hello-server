package main

import (
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gofr.dev/pkg/gofr"
)

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	app := gofr.New()

	app.GET("/greet", func(ctx *gofr.Context) (any, error) {
		return "Hello, World!", nil
	})

	app.GET("/redis", func(ctx *gofr.Context) (any, error) {
		ctx.Redis.Set(ctx.Context, "greeting", "Hello, Redis!", 0)
		pong := ctx.Redis.Ping(ctx.Context)
		fmt.Println("Pong: ", pong)
		val, err := ctx.Redis.Get(ctx.Context, "greeting").Result()
		fmt.Println("Val: ", val)
		if err != nil {
			if !errors.Is(err, redis.Nil) {
				fmt.Println("Key not found")
				return "Hello, World!", nil
			}
			fmt.Println("Error: ", err)
			return nil, err
		}

		return val, nil
	})

	app.Run()
}
