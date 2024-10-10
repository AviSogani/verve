package cron

import (
	"github.com/robfig/cron"
	"log"
	"verve/controller"
)

// Init initiates all the crons in the app
func Init() {
	c := cron.New()

	// Schedule the function to run every 1 minute
	err := c.AddFunc("@every 1m", controller.LogRequest)
	if err != nil {
		log.Fatalf("Error adding cron job: %v", err)
	}

	// extension #3
	//// Schedule the function to run every 1 minute
	//err := c.AddFunc("@every 1m", controller.LogRequestExtension)
	//if err != nil {
	//	log.Fatalf("Error adding cron job: %v", err)
	//}

	c.Start()
}
