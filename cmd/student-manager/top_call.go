import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sort"
)

type CallerStats struct {
	CallerID string
	Count    int64
}

func main() {
	// Create a Redis client
	rd := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // If your Redis server has a password
		DB:       0,  // Use the default Redis database
	})

	// Create a context
	ctx := context.Background()

	// Key for storing caller stats
	statsKey := "caller_stats"

	// Increment the count for the caller
	callerID := "caller1"
	err := rd.ZIncrBy(ctx, statsKey, 1, callerID).Err()
	if err != nil {
		fmt.Println("Error incrementing caller count:", err)
		return
	}

	// Get the top 10 callers with the most API calls
	results, err := rd.ZRevRangeWithScores(ctx, statsKey, 0, 9).Result()
	if err != nil {
		fmt.Println("Error retrieving top callers:", err)
		return
	}

	// Process the results
	var callerStats []CallerStats
	for _, result := range results {
		callerStats = append(callerStats, CallerStats{
			CallerID: result.Member.(string),
			Count:    int64(result.Score),
		})
	}

	// Sort the caller stats by call count
	sort.Slice(callerStats, func(i, j int) bool {
		return callerStats[i].Count > callerStats[j].Count
	})

	// Print the top 10 callers
	for i, stat := range callerStats {
		fmt.Printf("Rank %d: Caller %s with %d API calls\n", i+1, stat.CallerID, stat.Count)
	}

	// Close the Redis client
	err = rd.Close()
	if err != nil {
		fmt.Println("Error closing Redis client:", err)
		return
	}
}
