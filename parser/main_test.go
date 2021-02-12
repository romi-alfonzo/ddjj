package main

import (
	"reflect"
	"testing"
)

func TestMarioAbdo2016(t *testing.T) {

	data := handleSingleFile("./test_declarations/267948_MARIO_ABDO_BENITEZ.pdf")
	data.Print()

	if data.Data == nil {
		t.Errorf("Error parsing the document")
	}

	AssertEqual(t, "MARIO", data.Data.Nombre)
	AssertEqual(t, "2016-07-13", data.Data.Fecha.Format("2006-01-02"))
	AssertEqual(t, int64(3263852172), data.Data.Resumen.TotalActivo)
	AssertEqual(t, int64(241094919), data.Data.Resumen.TotalPasivo)
	AssertEqual(t, int64(3022757253), data.Data.Resumen.PatrimonioNeto)
}

// AssertEqual checks if values are equal
func AssertEqual(t *testing.T, want interface{}, got interface{}) {
	if want == got {
		return
	}
	// debug.PrintStack()
	t.Errorf("Received %v (type %v), expected %v (type %v)", got, reflect.TypeOf(got), want, reflect.TypeOf(want))
}
