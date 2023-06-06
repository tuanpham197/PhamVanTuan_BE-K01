import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	// Create a Redis client
	rd := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // If your Redis server has a password
		DB:       0,  // Use the default Redis database
	})

	// Create a context
	ctx := context.Background()

	// Key for storing HyperLogLog
	hllKey := "ping_callers"

	// Add caller IDs to the HyperLogLog
	callerIDs := []string{"caller1", "caller2", "caller3"}
	err := rd.PFAdd(ctx, hllKey, callerIDs...).Err()
	if err != nil {
		fmt.Println("Error adding caller IDs to HyperLogLog:", err)
		return
	}

	// Count the approximate number of unique callers
	count, err := rd.PFCount(ctx, hllKey).Result()
	if err != nil {
		fmt.Println("Error counting unique callers:", err)
		return
	}

	// Print the approximate count
	fmt.Println("Approximate number of unique callers:", count)

	// Close the Redis client
	err = rd.Close()
	if err != nil {
		fmt.Println("Error closing Redis client:", err)
		return
	}
}
