package decimalutil

import (
	"github.com/shopspring/decimal"
	"reflect"
	"testing"
)

// use this decimal to test against
var dec = decimal.NewFromFloat(2020)

func TestCentsToDecimal(t *testing.T) {
	c := CentsToDecimal(2020)
	if c != dec {
		t.Error("Expected 20.20, recieved ", c)
	}
}

func TestDecimalToCents(t *testing.T) {
	c := DecimalToCents(dec)
	if reflect.TypeOf(c).Kind() != reflect.Int64 {
		t.Error("Expected type int64, recieved type: ", reflect.TypeOf(c))
	}
}

func TestDecimalToString(t *testing.T) {
	c := DecimalToString(dec)
	if reflect.TypeOf(c).Kind() != reflect.String {
		t.Error("Expected type String, recieved type: ", reflect.TypeOf(c))
	}

}
