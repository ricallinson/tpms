package tpms

import (
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
	"os"
	"log"
	"fmt"
)

func TestSensor(t *testing.T) {

	var sensor *Sensor

	BeforeEach(func() {
		sensor = &Sensor{}
	})

	AfterEach(func() {

	})

	Describe("Tpms()", func() {

		It("should return a Sensor object", func() {
			AssertEqual(reflect.TypeOf(sensor).String(), "*tpms.Sensor")
		})

		It("should ", func() {
			file, err := os.Open("./fixtures/rawdata")
			if err != nil {
				log.Fatal(err)
			}
			data := make([]byte, 1872)
			file.Read(data)
			for i := 0; i <1872; i = i + 18 {
				// fmt.Println(data[i+1:i+4])
				// fmt.Println(data[i:i+18])
				// fmt.Println(data[i+8:i+18])
				sensor.ParseData(data[i:])
				fmt.Printf("kPa: %d, °C: %d\n", sensor.Kilopascal, sensor.Celsius)
			}
		})
	})

	Report(t)
}
