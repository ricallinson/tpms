package tpms

import (
	"context"
	"fmt"
	"github.com/go-ble/ble"
	"github.com/go-ble/ble/examples/lib/dev"
	"strconv"
	"strings"
	"time"
	"log"
	"github.com/pkg/errors"
)

type Tpms struct {
	sensors [4]*Sensor
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

func (this *Tpms) Update() {
	for _, sensor := range this.sensors {
		this.updateSensor(sensor)
	}
}

func (this *Tpms) Close() {

}

func (this *Tpms) updateSensor(sensor *Sensor) {
	ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), 5*time.Second))
	client, err := ble.Dial(ctx, sensor.Address)
	checkBleError(err)

	fmt.Println(client.Name())

	done := make(chan struct{})
	go func() {
		<-client.Disconnected()
		fmt.Printf("[ %s ] is disconnected \n", client.Addr())
		close(done)
	}()

	client.CancelConnection()
	<-done
}

func (this *Tpms) scan() {
	filter := func(a ble.Advertisement) bool {
		return strings.HasPrefix(strings.ToUpper(a.LocalName()), "TPMS")
	}
	retry := 10
	for !this.gotSensors() || retry > 0 {
		retry--
		ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), 1*time.Second))
		err := ble.Scan(ctx, false, this.registerSensor, filter)
		checkBleError(err)
	}
}

func (this *Tpms) registerSensor(a ble.Advertisement) {
	pos, _ := strconv.Atoi(string(a.LocalName()[4]))
	if pos < 1 || pos > 4 {
		return
	}
	if this.sensors[pos-1] == nil {
		this.sensors[pos-1] = &Sensor{
			Address: a.Addr(),
		}
		fmt.Printf("Sensor %d found\n", pos)
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