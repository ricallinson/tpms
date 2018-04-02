package tpms

import (
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
)

func TestTpms(t *testing.T) {

	var tires *Tpms

	BeforeEach(func() {
		tires = NewTpms()
	})

	AfterEach(func() {
		tires.Close()
	})

	Describe("Tpms()", func() {

		It("should return a Tpms object", func() {
			AssertEqual(reflect.TypeOf(tires).String(), "*tpms.Tpms")
		})

		It("should ", func() {
			tires.scan()
			tires.Update()
		})
	})

	Report(t)
}
