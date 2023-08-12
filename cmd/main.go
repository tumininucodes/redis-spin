package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)


var ctx = context.Background()

func main() {


	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	ping, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err.Error())
	}

	

	fmt.Println(ping)
}