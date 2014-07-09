package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/spark"
	"time"
)

func main() {
	gbot := gobot.NewGobot()
	api.NewAPI(gbot).Start()

	sparkCore := spark.NewSparkCoreAdaptor("spark", "device_id", "access_token")
	led := gpio.NewLedDriver("led", sparkCore, "D7")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("spark",
		[]gobot.Connection{sparkCore},
		[]gobot.Device{led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
