package datagen

import (
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {

	type T struct {
		Values map[string]string
		Uid    string `datagen:"uuid"`
		Info   struct {
			Name    string `datagen:"name"`
			Age     int    `datagen:"integer|10,40"`
			Address string `datagen:"address"`
			Source  float64
		}
	}
	var testt T
	b, err := StructGen(testt, nil)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(string(b))

}
