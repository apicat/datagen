package datagen

import (
	"fmt"
	"testing"
)

func TestFunctionStr(t *testing.T) {

	fmt.Println(CallFunction("title(zh)|3,5"))
	fmt.Println(CallFunction("provinceorstatecity(zh)"))
	fmt.Println(CallFunction("street(zh)"))
	fmt.Println(CallFunction("boolean|true"))
	fmt.Println(CallFunction("float|1.0"))
	fmt.Println(CallFunction("float|1.0,10,2"))
	fmt.Println(CallFunction("string|5"))
}
