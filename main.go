package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "host.docker.internal:6379",
		DB:   0,
	})

	go func() {
		subscription := rdb.Subscribe(ctx, "hotel.events")
		for msg := range subscription.Channel() {
			fmt.Printf("🔓 Unlocking door for: %s\n", msg.Payload)
		}
	}()

	http.HandleFunc("/render", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
