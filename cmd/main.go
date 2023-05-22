package main

import (
	"github.com/shukra-in-spirit/haste-scheduler/internal/crons"
	"github.com/shukra-in-spirit/haste-scheduler/internal/data"
	"github.com/shukra-in-spirit/haste-scheduler/internal/handler"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Run the cron job at a fixed time of 09:00 AM
	// fixedTime := 9 * time.Hour

	// // Run the cron job function
	// crons.RunDailyCron(fixedTime)
	// Run the cron job every second
	NewChronosInstance, err := handler.NewChronos()
	if err != nil {
		log.Fatalf("The creation of a Chronos Instance failed with error: ", err)
	}
	crons.RunEverySecondCron(&NewChronosInstance)
	// Create a new Gin router
	router := gin.Default()

	// Define the /version endpoint
	router.GET("/version", func(c *gin.Context) {
		c.JSON(200, gin.H{"version": "v0.1.0"})
	})
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "Service is healthy"})
	})
	router.POST("/schedule", func(c *gin.Context) {

		// Call the desired function, passing the additional variable
		ScheduleEndpointFunction(c, NewChronosInstance)
	})

	// Run the Gin server in the main thread
	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start the server:", err)
	}

}

func ScheduleEndpointFunction(ctx *gin.Context, chronos handler.Chronos) {
	scheduleItem := data.ScheduleItem{}
	if err := ctx.ShouldBindJSON(&scheduleItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
			"body":  err.Error(),
		})
		return
	}
	chronos.CreateSchedule(ctx.Request.Context(), scheduleItem)
	fmt.Printf(scheduleItem.Name, " item created at ", time.Now())
	ctx.JSON(http.StatusOK, nil)
	return
}
