# Tire Pressure Management System

A Go library to read the tire pressures and temperatures from the [ZEEPIN TPMS Sensor Bluetooth Low Energy Tire Pressure Monitoring System](https://www.amazon.com/gp/product/B079JXMM2P/ref=oh_aui_detailpage_o02_s00?ie=UTF8&psc=1).

## Install Example

You must first install the `Go` environment for your operating system from [golang.org](https://golang.org/dl/). Once installed open your terminal program.

* MAC: Terminal. To open the terminal, `apple+space` and type `terminal`.
* Windows: Command Prompt. To open the command prompt, `windows+r` and type `cmd`.
* Linux: If you don't know you shouldn't be using Linux.

You should now see your command line interface. From here you will install the program and execute it. Enter the following at the command line;

	go get github.com/ricallinson/tpms
	go install github.com/ricallinson/tpms
	cd $GOPATH/src/github.com/ricallinson/tpms/examples/tpmssh
	go install
	tpmssh -duration 10s

The `tpmssh` should now run for ten seconds.

For subsequent uses you only need to start the program. No need to install it again.

	tpmssh -duration 10s

### Example Usage

The example can run indefinitely by excluding the `-duration` argument.

	tmpssh

Alternatively you can provide any number of seconds, minutes or hours for a monitoring duration.

	tmpssh -duration 30s
	tmpssh -duration 20m
	tmpssh -duration 24h

In all cases a log file will be written to `./log` containing the raw sensor data as a byte stream.

## API Usage

The API is simply creating an instance of Tpms, starts monitoring and reads the values. The values are checked every second but it depends on the sensors being active. In testing it took 10 minutes to retrieve all four sets of data when there were no pressure changes.

	package main

	import(
		"fmt"
		"github.com/ricallinson/tpms"
	)

	func main() {
		tires, err := tpms.NewTpms()
		if err != nil {
			fmt.Println(err)
			return
		}
		tires.StartMonitoring()
		defer tires.StopMonitoring()
		for tires.Read()[0] == nil {}
		sensors := tires.Read()
		fmt.Printf("%d kPa, °C %d\n", sensors[0].Kilopascal, sensors[0].Celsius)
	}

There is an option for logging the raw data from the sensors. This will keep logging until `.StopMonitoring()` is called or the process is exited.

	tires, _ := tpms.NewTpms()
	tires.Log("/path/to/file")

## Testing

	cd $GOPATH/src/github.com/ricallinson/tpms
	go test

## Coverage

	cd $GOPATH/src/github.com/ricallinson/tpms
	go test -covermode=count -coverprofile=count.out; go tool cover -html=count.out
