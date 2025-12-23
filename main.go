package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	// Context is required for all redis/v9 calls
	ctx = context.Background()
	// Global variable to hold our Redis connection
	rdb *redis.Client
)

func main() {
	// 1. Initialize the Redis Client
	// In GCP, this address would be the IP of your Memorystore instance
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Address of the docker container
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	// 2. Test the connection immediately (Fail fast if Redis is down)
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("❌ Could not connect to Redis: %v", err)
	}
	fmt.Printf("✅ Connected to Redis! Response: %s\n", pong)

	// 3. Define HTTP routes
	http.HandleFunc("/health", handleHealth)

	// 4. Start the Server
	port := ":8080"
	fmt.Printf("🚀 Toy Bidder running on http://localhost%s\n", port)
	
	// This blocks forever until the server crashes
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// handleHealth checks if the server is up AND if Redis is writable
func handleHealth(w http.ResponseWriter, r *http.Request) {
	// Let's write a timestamp to Redis to prove we have "Write" access
	// In AdTech, if you can't write to Redis, you can't bid.
	err := rdb.Set(ctx, "last_health_check", time.Now().String(), 0).Err()
	if err != nil {
		http.Error(w, "Redis write failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK: Bidder is healthy and Redis is writable!"))
}