package handler

import (
	"github.com/shukra-in-spirit/haste-scheduler/internal/constants"
	"github.com/shukra-in-spirit/haste-scheduler/internal/data"
	"github.com/shukra-in-spirit/haste-scheduler/internal/sorted_set"
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type ChronosHandlerInterface interface {
	CreateSchedule(ctx context.Context, request data.ScheduleItem) (*data.ScheduleResponse, error)
}

type Chronos struct {
	SortedSet *sorted_set.SortedSet
}

func NewChronos() (Chronos, error) {
	// Connect to Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddress,  // Replace with your Redis server address
		Password: constants.RedisPassword, // Set password if required
		DB:       constants.RedisDB,       // Use default Redis database
	})

	// Test the connection to Redis
	_, err := redisClient.Ping(redisClient.Context()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	// Create a new priority queue
	priorityQueue := sorted_set.NewSortedSet(redisClient, constants.SecondsQueue)

	newChronos := Chronos{
		SortedSet: priorityQueue,
	}
	return newChronos, nil
}

func (c *Chronos) CreateSchedule(ctx context.Context, request data.ScheduleItem) *data.ScheduleResponse {
	score := calculateScore(request.Period)
	//fmt.Printf("HALLELUJAH", score)
	scheduleItem := data.PriorityQueueItem{
		ItemName: request.Name,
		Message:  request.Message,
		Priority: score,
		Url:      request.Url,
	}
	go func() {
		c.SortedSet.ZADD(&scheduleItem)
	}()
	response := data.ScheduleResponse{
		StatusCode: 200,
	}
	return &response
}

func calculateScore(period float64) float64 {
	score := time.Since(constants.OriginDate).Seconds() + period
	//fmt.Printf("ItemSCORE", score)
	return score
}
