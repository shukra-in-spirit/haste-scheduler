package crons

import (
	"github.com/shukra-in-spirit/haste-scheduler/internal/constants"
	"github.com/shukra-in-spirit/haste-scheduler/internal/handler"
	"fmt"
	"log"
	"time"
	"sync"
)

func RunEverySecondCron(chronos *handler.Chronos) {
	ticker := time.NewTicker(time.Second)
	go func() {
		for range ticker.C {
			var allExpiredWG sync.WaitGroup
			currentScore := 1.1
			itemScore := 1.0
			timeAtStart := time.Now()
			for i:=0; i<25; i++ {
				allExpiredWG.Add(1)
				go func() {
					defer allExpiredWG.Done()
					// Run your desired task or function here
					countOfItems, err := chronos.SortedSet.ZCARD()
					if err != nil {
						log.Fatalf("Failed getting a count of items in db with error: ", err)
					}
					if countOfItems == 0 {
						fmt.Println("The Queue is empty")
						return
					} else {
						scheduleItem, err := chronos.SortedSet.ZPEEKMIN()
						scheduleItemDirect := *scheduleItem
						// fmt.Printf("TADA", *scheduleItem)
						if err != nil {
							log.Fatalf("Peek failed during the cron job at time: ", time.Now().Format("15:04:05"), " with error ", err)
						}
						itemScore = scheduleItemDirect.Priority
						currentScore = time.Since(constants.OriginDate).Seconds()
						if currentScore >= itemScore {
							// fmt.Printf("This is the current score: ", currentScore)
							// fmt.Printf("This is the item score: ", itemScore)
							expiredItem, err := chronos.SortedSet.ZPOPMIN()
							if err != nil {
								log.Fatalf("Pop failed during the cron job at time: ", time.Now().Format("15:04:05"), " with error ", err)
							}
							fmt.Println("Item ", *expiredItem, " expired at ", time.Now())
						}
					}
				}()
			}
			// to make sure that processing time + sleep time = 1 second
			timeAtEnd := time.Now()
			timeToSleep := time.Second - timeAtEnd.Sub(timeAtStart)
			fmt.Printf(timeToSleep.String())
			time.Sleep(timeToSleep)
		}
	}()
}
