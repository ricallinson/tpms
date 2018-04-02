package main

import (
	"fmt"
	"github.com/ricallinson/tpms"
	"time"
)

func main() {
	tires := tpms.NewTpms()
	tires.StartMonitoring()
	for {
		for _, sensor := range tires.Read() {
			if sensor != nil {
				fmt.Printf("ID: %d, kPa: %d, °C: %d\n", sensor.Id, sensor.Kilopascal, sensor.Celsius)
			}
		}
		fmt.Println()
		time.Sleep(5 * time.Second)
	}
	tires.StopMonitoring()
}
