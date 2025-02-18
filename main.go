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

	// API routes
	app.GET("/greet", GreetHandler)

	// API Redis routes
	app.GET("/redis", RedisHandler)

	// API SQL routes
	app.POST("/customer/{name}", AddCustomerHandler)
	app.GET("/customers", ListCustomersHandler)

	app.Run()
}

func GreetHandler(ctx *gofr.Context) (any, error) {
	return "Hello, World!", nil
}

func RedisHandler(ctx *gofr.Context) (any, error) {
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
}

func AddCustomerHandler(ctx *gofr.Context) (any, error) {
	name := ctx.PathParam("name")

	_, err := ctx.SQL.ExecContext(ctx, "INSERT INTO customers (name) VALUES (?)", name)
	return nil, err
}

func ListCustomersHandler(ctx *gofr.Context) (any, error) {
	var customers []Customer

	rows, err := ctx.SQL.QueryContext(ctx, "SELECT id, name FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.ID, &customer.Name)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}
