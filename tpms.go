package tpms

import (
	"context"
	"fmt"
	"github.com/go-ble/ble"
	"github.com/go-ble/ble/examples/lib/dev"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"strings"
	"time"
)

type Tpms struct {
	sensors [4]*Sensor
	log     *os.File
	running bool
}

func NewTpms() (*Tpms, error) {
	this := &Tpms{
		sensors: [4]*Sensor{},
	}
	d, err := dev.NewDevice("")
	if err != nil {
		return nil, err
	}
	ble.SetDefaultDevice(d)
	return this, nil
}

func (this *Tpms) StartMonitoring() {
	this.running = true
	filter := func(a ble.Advertisement) bool {
		return strings.HasPrefix(strings.ToUpper(a.LocalName()), "TPMS")
	}
	go func() {
		for this.running {
			ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), 1*time.Second))
			err := ble.Scan(ctx, false, this.updateSensor, filter)
			checkBleError(err)
		}
	}()
}

func (this *Tpms) StopMonitoring() {
	this.running = false
	if this.log != nil {
		this.log.Close()
	}
}

func (this *Tpms) Read() [4]*Sensor {
	return this.sensors
}

func (this *Tpms) Log(file string) {
	this.log, _ = os.Create(file)
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
		fmt.Println(err.Error())
		return false
	}
	return false
}
