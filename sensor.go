package tpms

import (
	"github.com/go-ble/ble"
)

type Sensor struct {
	Id		 int
	Address  ble.Addr
	Pressure int
	Temp     int
}

func (this *Sensor) ParseData(b []byte) {
	
}
