package tpms

import (
	"github.com/go-ble/ble"
)

type Sensor struct {
	Address  ble.Addr
	Pressure int
	Temp     int
}
