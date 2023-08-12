package internal

import (
	"fmt"
	"github.com/go-redis/redis"
)

func OpenDB() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ping, err := client.Ping().Result()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("redis client >> ", ping)

	return client
}


func GetEntry(id string, client *redis.Client) (*Reminder, error) {
	var reminder Reminder
	result, err := client.Get(id).Result()
	if err != nil {
		return nil, err
	}
	reminder.ID = id
	reminder.Value = result
	return &reminder, nil
}


func AddToDB(client *redis.Client, reminder *Reminder) (*Reminder, *error) {

	err := client.Set(reminder.ID, reminder.Value, 0).Err()
	if err != nil {
		return nil, &err
	}
	return reminder, nil

}
