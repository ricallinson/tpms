package main

import (
	"flag"
	"fmt"
	"github.com/ricallinson/tpms"
	"log"
	"time"
)

var (
	duration = flag.Duration("duration", 0, "monitoring duration Xs, 0 for indefinitely")
)

func main() {
	flag.Parse()
	tires, err := tpms.NewTpms()
	if err != nil {
		log.Fatal(err)
	}
	tires.Log("./log")
	tires.StartMonitoring()
	start := time.Now()
	defer tires.StopMonitoring()
	for *duration == 0 || time.Now().Sub(start) < *duration {
		for _, sensor := range tires.Read() {
			if sensor != nil {
				fmt.Printf("Sensor: %d, kPa: %d, °C: %d\n", sensor.Id, sensor.Kilopascal, sensor.Celsius)
			}
		}
		time.Sleep(5 * time.Second)
	}
}
