package tpms

import (
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
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
			// 1
			sensor.ParseData([]byte{0, 1, 128, 234, 202, 16, 4, 102, 127, 99, 0, 0, 179, 7, 0, 0, 94, 1})
			sensor.ParseData([]byte{0, 1, 128, 234, 202, 16, 4, 102, 228, 70, 0, 0, 180, 7, 0, 0, 94, 1})
			// 2
			sensor.ParseData([]byte{0, 1, 129, 234, 202, 32, 2, 4, 200, 68, 0, 0, 177, 7, 0, 0, 100, 1})
			sensor.ParseData([]byte{0, 1, 129, 234, 202, 32, 2, 4, 244, 51, 0, 0, 178, 7, 0, 0, 100, 1})
			// 3
			sensor.ParseData([]byte{0, 1, 130, 234, 202, 48, 2, 237, 24, 70, 0, 0, 154, 7, 0, 0, 94, 1})
			sensor.ParseData([]byte{0, 1, 130, 234, 202, 48, 2, 237, 0, 37, 0, 0, 154, 7, 0, 0, 94, 1})
			sensor.ParseData([]byte{0, 1, 130, 234, 202, 48, 2, 237, 183, 24, 0, 0, 154, 7, 0, 0, 94, 1})
			// 4
			sensor.ParseData([]byte{0, 1, 131, 234, 202, 64, 1, 118, 7, 172, 0, 0, 182, 7, 0, 0, 100, 0})
			sensor.ParseData([]byte{0, 1, 131, 234, 202, 64, 1, 118, 127, 16, 0, 0, 178, 7, 0, 0, 100, 1})
		})
	})

	Report(t)
}
