package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Create a Redis client
	rd := redis.NewClient(&redis.Options{
		Addr:     "localhost:6377",
		Password: "", // If your Redis server has a password
		DB:       0,  // Use the default Redis database
	})

	// Create a context
	ctx := context.Background()

	// Key prefix for rate limiting
	rateLimitKey := "rate_limit:ping:"

	// User ID or IP address of the caller
	userID := "user123"

	// Check if the user has reached the rate limit
	count, err := rd.Incr(ctx, rateLimitKey+userID).Result()
	if err != nil {
		fmt.Println("Error checking rate limit:", err)
		return
	}

	// If the count is 1, set the expiry time to 60 seconds
	if count == 1 {
		err := rd.Expire(ctx, rateLimitKey+userID, 60*time.Second).Err()
		if err != nil {
			fmt.Println("Error setting rate limit expiry:", err)
			return
		}
	}

	// If the count is greater than 2, the user has exceeded the rate limit
	if count > 2 {
		fmt.Println("Rate limit exceeded")
		return
	}

	// The user is within the rate limit, proceed with the API logic
	fmt.Println("API /ping logic executed")

	// Close the Redis client
	err = rd.Close()
	if err != nil {
		fmt.Println("Error closing Redis client:", err)
		return
	}
}
