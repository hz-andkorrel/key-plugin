package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	redis := redis.NewClient(&redis.Options{
		Addr: "host.docker.internal:6379",
		DB:   0,
	})

	subscription := redis.Subscribe(ctx, "hotel.events")

	for msg := range subscription.Channel() {
		fmt.Printf("🔓 Unlocking door for: %s\n", msg.Payload)
	}
}
