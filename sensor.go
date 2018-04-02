package tpms

import (
	"github.com/go-ble/ble"
	// "fmt"
	"encoding/binary"
	// "bytes"
	// "math"
)

type Sensor struct {
	Id         int
	Address    ble.Addr
	Kilopascal int
	Celsius    int
}

func (this *Sensor) ParseData(b []byte) {
	// The first 2 bytes are unknown.
	// The next 8 bytes are unknown.
	// The next 4 bytes are pressure in kPa.
	// The last 4 bytes are temperature in Celsius.
	this.Kilopascal = int(binary.BigEndian.Uint32(b[9:]) / 100000000)
	this.Celsius = int(binary.BigEndian.Uint32(b[14:]) / 1000)
}
