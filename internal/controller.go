package internal

import (
	"fmt"
	"github.com/go-redis/redis"
)

func OpenDB() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	ping, err := client.Ping().Result()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("redis client >> ", ping)

	return client
}


func AddToDB(client *redis.Client, reminder *Reminder) {

	client.Set(reminder.ID, reminder.Value)
}
