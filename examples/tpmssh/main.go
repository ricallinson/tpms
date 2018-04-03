package main

import (
	"fmt"
	"github.com/ricallinson/tpms"
	"time"
)

func main() {
	tires := tpms.NewTpms()
	tires.Log("./log")
	tires.StartMonitoring()
	for {
		for _, sensor := range tires.Read() {
			if sensor != nil {
				fmt.Printf("Sensor: %d, kPa: %d, °C: %d\n", sensor.Id, sensor.Kilopascal, sensor.Celsius)
			}
		}
		time.Sleep(5 * time.Second)
	}
	tires.StopMonitoring()
}
