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
	// fmt.Printf("Data for sensor %d\n", this.Id)
	// fmt.Println(b)
	// fmt.Println(b[8:12])
	// fmt.Println("?", binary.BigEndian.Uint64(b[0:]))
	// fmt.Println("p", binary.BigEndian.Uint32(b[8:]) / 100000000)
	// fmt.Println("?", binary.BigEndian.Uint32(b[12:]) / 100000000)
	// fmt.Println("t", binary.BigEndian.Uint16(b[16:]) / 1000)

	this.Kilopascal = int(binary.BigEndian.Uint32(b[9:]) / 10000000)
	this.Celsius = int(binary.BigEndian.Uint32(b[14:]) / 1000)
}
