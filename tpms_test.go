package tpms

import (
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
)

func TestTpms(t *testing.T) {

	var tires *Tpms

	BeforeEach(func() {
		tires, _ = NewTpms()
	})

	AfterEach(func() {
		tires.StopMonitoring()
	})

	Describe("Tpms()", func() {

		It("should return a Tpms object", func() {
			AssertEqual(reflect.TypeOf(tires).String(), "*tpms.Tpms")
		})
	})

	Report(t)
}
