package crons

import (
	"github.com/shukra-in-spirit/haste-scheduler/internal/constants"
	"github.com/shukra-in-spirit/haste-scheduler/internal/handler"
	"fmt"
	"log"
	"time"
)

func RunDailyCron(fixedTime time.Duration, chronos handler.Chronos) {
	now := time.Now()

	// Calculate the duration until the next desired run time
	// based on the fixed time provided
	desiredTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Add(fixedTime)
	duration := desiredTime.Sub(now)

	// If the desired run time has already passed for the day, add 24 hours to the duration
	if duration < 0 {
		duration += 24 * time.Hour
	}

	// Create a ticker that ticks at the desired run time
	ticker := time.NewTicker(duration)

	// Create a goroutine to run the job
	go func() {
		for {
			<-ticker.C // Wait for the ticker to tick
			// Run your desired task or function here
			currentScore := 0.0
			itemScore := 1.0
			for currentScore >= itemScore {
				// Run your desired task or function here
				scheduleItem, err := chronos.SortedSet.ZPEEKMIN()
				if err != nil {
					log.Fatalf("Peek failed during the cron job at time: ", time.Now().Format("15:04:05"), " with error ", err)
				}
				itemScore = scheduleItem.Priority
				currentScore = time.Since(constants.OriginDate).Seconds()
				if currentScore >= itemScore {
					expiredItem, err := chronos.SortedSet.ZPOPMIN()
					if err != nil {
						log.Fatalf("Pop failed during the cron job at time: ", time.Now().Format("15:04:05"), " with error ", err)
					}
					fmt.Println("Item ", &expiredItem, " expired at ", time.Now().Format("15:04:05"))
				}
			}
			// Calculate the duration until the next desired run time
			duration = 24 * time.Hour

			// Reset the ticker with the new duration
			ticker.Reset(duration)
		}
	}()
}
