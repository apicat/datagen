package datagen

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	fmt.Println(String("number"))

	tp := Typography()
	fmt.Println(tp.Word())
	fmt.Println(tp.Phrase())
	fmt.Println(tp.Title())
	fmt.Println(tp.Sentence())
	fmt.Println(tp.Paragraph())

	tp2 := Typography("zh")
	fmt.Println(tp2.Word())
	fmt.Println(tp2.Phrase())
	fmt.Println(tp2.Title())
	fmt.Println(tp2.Sentence())
	fmt.Println(tp2.Paragraph())
}

func TestInternet(t *testing.T) {
	fmt.Println(UUID())
	fmt.Println(URL())
	fmt.Println(Domain())
	fmt.Println(IPv4())
	fmt.Println(IPv6())
}

func TestPeople(t *testing.T) {
	// fmt.Println(Name())
	// fmt.Println(Name("zh"))
	// fmt.Println(Pick("男", "女"))
	// fmt.Println(Phone())
	// fmt.Println(Phone("zh"))
	// fmt.Println(Email())
	// fmt.Println(IDCard())
	// fmt.Println(IDCard("zh"))
}

func TestAddress(t *testing.T) {
	fmt.Println(ProvinceorState())
	fmt.Println(ProvinceorState("zh"))
	fmt.Println(City())
	fmt.Println(City("zh"))
	fmt.Println(Street())
	fmt.Println(Street("zh"))
	fmt.Println(Address())
	fmt.Println(Address("zh"))
	fmt.Println(ZipCode())
	fmt.Println(ZipCode("zh"))
}

func TestTime(t *testing.T) {
	fmt.Println(DateTime("YYYY-MM-dd HH:mm:ss"))
	fmt.Println(Time("m:s"))
	fmt.Println(Date("YY/M/d"))
	fmt.Println(Timestamp())
	fmt.Println(Now("HH"))
}

func TestColor(t *testing.T) {
	fmt.Println(Color("rgb"))
	fmt.Println(Color("rgba"))
	fmt.Println(Color("hsl"))
	fmt.Println(Color("hex"))
}

func TestImage(t *testing.T) {
	fmt.Println(ImageData())
	fmt.Println(ImageData(32))
	fmt.Println(ImageData(128, 48))
	fmt.Println(ImageURL())
	fmt.Println(ImageURL(200))
	fmt.Println(ImageURL(200, 100))
}
