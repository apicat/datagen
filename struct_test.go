package datagen

import (
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {

	type T struct {
		values map[string]string
		uid    string `datagen:"uuid"`
		info   struct {
			name    string `datagen:"name"`
			age     int    `datagen:"integer|10,40"`
			address string `datagen:"address"`
		}
	}
	var testt T
	b, err := StructGen(testt, nil)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(string(b))

}
