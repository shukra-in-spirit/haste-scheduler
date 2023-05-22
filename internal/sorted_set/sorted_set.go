package sorted_set

import (
	"github.com/shukra-in-spirit/haste-scheduler/internal/data"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type SortedSet struct {
	redisClient *redis.Client
	queueName   string
}

func NewSortedSet(redisClient *redis.Client, setName string) *SortedSet {
	return &SortedSet{
		redisClient: redisClient,
		queueName:   setName,
	}
}

func (pq *SortedSet) ZCARD() (int64, error) {
	count, err := pq.redisClient.ZCard(pq.redisClient.Context(), pq.queueName).Result()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (pq *SortedSet) ZCOUNT(min, max float64) (int64, error) {
	result, err := pq.redisClient.ZCount(pq.redisClient.Context(), pq.queueName, fmt.Sprintf("%f", min), fmt.Sprintf("%f", max)).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (pq *SortedSet) ZPOPMIN() (*data.PriorityQueueItem, error) {
	result, err := pq.redisClient.ZPopMin(pq.redisClient.Context(), pq.queueName).Result()
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil // No items in the sorted set
	}

	item := &data.PriorityQueueItem{
		ItemName: result[0].Member.(string),
		Priority: result[0].Score,
	}

	return item, nil
}

func (pq *SortedSet) ZPOPMAX() (*data.PriorityQueueItem, error) {
	result, err := pq.redisClient.ZPopMax(pq.redisClient.Context(), pq.queueName).Result()
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil // No items in the sorted set
	}

	item := &data.PriorityQueueItem{
		ItemName: result[0].Member.(string),
		Priority: result[0].Score,
	}

	return item, nil
}

func (pq *SortedSet) ZADD(items ...*data.PriorityQueueItem) error {
	//fmt.Println("OPERATION SUCCESSFUL")
	var cmds []*redis.Z
	for _, item := range items {
		cmds = append(cmds, &redis.Z{
			Score:  item.Priority,
			Member: item.ItemName,
		})
	}
	//fmt.Printf("", (*cmds[0]))
	_, err := pq.redisClient.ZAdd(pq.redisClient.Context(), pq.queueName, cmds...).Result()
	//fmt.Println("REACHED HERE")
	if err != nil {
		return err
	}

	return nil
}

func (pq *SortedSet) ZPEEKMIN() (*data.PriorityQueueItem, error) {
	result, err := pq.redisClient.ZRangeWithScores(pq.redisClient.Context(), pq.queueName, 0, 0).Result()
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil // No items in the sorted set
	}

	item := &data.PriorityQueueItem{
		ItemName: result[0].Member.(string),
		Priority: result[0].Score, // You may need to retrieve the actual score from Redis separately
	}

	return item, nil
}

func (pq *SortedSet) ZPEEKMAX() (*data.PriorityQueueItem, error) {
	result, err := pq.redisClient.ZRevRange(pq.redisClient.Context(), pq.queueName, -1, -1).Result()
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil // No items in the sorted set
	}

	item := &data.PriorityQueueItem{
		ItemName: result[0],
		Priority: 0, // You may need to retrieve the actual score from Redis separately
	}

	return item, nil
}
