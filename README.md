# Tire Pressure Management System

A Go library to read the tire pressures and temperatures from the [ZEEPIN TPMS Sensor Bluetooth Low Energy Tire Pressure Monitoring System](https://www.amazon.com/gp/product/B079JXMM2P/ref=oh_aui_detailpage_o02_s00?ie=UTF8&psc=1).

## installation

	go get github.com/ricallinson/tpms

## Usage

The API is simply creating an instance of Tmps, start monitoring and read the values. The values are checked ever second but it depends on the sensors being active. In testing it took 10 minutes to retrieve all four sets of data when there were no pressure changes.

	import("github.com/ricallinson/tpms")

	tires, err := tpms.NewTpms()
	tires.StartMonitoring()
	defer tires.StopMonitoring()
	sensors := tires.Read()

There is an option for logging the raw data from the sensors. This will keep logging until `.StopMonitoring()` is called or the processes is exited.

	tires, err := tpms.NewTpms()
	tires.Log("/path/to/file")
