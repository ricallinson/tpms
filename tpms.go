package tpms

import (
	"context"
	"fmt"
	"github.com/go-ble/ble"
	"github.com/go-ble/ble/examples/lib/dev"
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Tpms struct {
	sensors [4]*Sensor
	log     *os.File
}

func NewTpms() *Tpms {
	this := &Tpms{
		sensors: [4]*Sensor{},
	}
	d, err := dev.NewDevice("default")
	if err != nil {
		fmt.Printf("can't new device : %s", err)
	}
	ble.SetDefaultDevice(d)
	return this
}

func (this *Tpms) StartMonitoring() {
	filter := func(a ble.Advertisement) bool {
		return strings.HasPrefix(strings.ToUpper(a.LocalName()), "TPMS")
	}
	go func() {
		for {
			ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), 1*time.Second))
			err := ble.Scan(ctx, false, this.updateSensor, filter)
			checkBleError(err)
		}
	}()
}

func (this *Tpms) Log(file string) {
	this.log, _ = os.Create(file)
}

func (this *Tpms) StopMonitoring() {

}

func (this *Tpms) Read() [4]*Sensor {
	return this.sensors
}

func (this *Tpms) updateSensor(a ble.Advertisement) {
	pos, _ := strconv.Atoi(string(a.LocalName()[4]))
	if pos < 1 || pos > 4 {
		return
	}
	if this.sensors[pos-1] == nil {
		this.sensors[pos-1] = &Sensor{
			Id:      pos,
			Address: a.Addr(),
		}
		fmt.Printf("Sensor %d added.\n", pos)
	}
	if len(a.ManufacturerData()) > 0 {
		this.sensors[pos-1].ParseData(a.ManufacturerData())
		if this.log != nil {
			this.log.Write(a.ManufacturerData())
		}
	}
}

func (this *Tpms) gotSensors() bool {
	for _, sensor := range this.sensors {
		if sensor == nil {
			return false
		}
	}
	return true
}

func checkBleError(err error) bool {
	switch errors.Cause(err) {
	case nil:
	case context.DeadlineExceeded:
		return true
	case context.Canceled:
		return true
	default:
		log.Fatalf(err.Error())
	}
	return false
}
