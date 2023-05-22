package constants

import "time"

const (
	RedisAddress  = "localhost:6379"
	RedisPassword = "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81"
	RedisDB       = 0
	SecondsQueue  = "priority_vector_queue"
	ParallelCrons = 10
)

var (
	OriginDate = time.Date(2023, time.May, 19, 0, 0, 0, 0, time.UTC)
)
