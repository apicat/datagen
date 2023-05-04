package datagen

import (
	"fmt"
	"strings"
)

// City 生成一个城市
func City(local ...string) string {
	return GetLocale(local...).Address.Region.RandomCity()
}

// ProvinceorState 生成一个省/邦/州
func ProvinceorState(local ...string) string {
	return GetLocale(local...).Address.Region.RandomStateOrProvince()
}

// ZipCode 生成邮编
func ZipCode(local ...string) string {
	return NumberPattern(GetLocale(local...).Address.ZipCode)
}

// Street 生成一个街道
func Street(local ...string) string {
	return street(GetLocale(local...))
}

// Address 生成一个详细地址
func Address(local ...string) string {
	loc := GetLocale(local...)
	a := loc.Address
	var b strings.Builder
	if a.Region.Reverse {
		b.WriteString(street(loc))
		b.WriteString(a.Region.Split)
		b.WriteString(a.Region.RandomString())
	} else {
		b.WriteString(a.Region.RandomString())
		b.WriteString(a.Region.Split)
		b.WriteString(street(loc))
	}
	b.WriteByte(0x20)
	b.WriteString(NumberPattern(loc.Address.ZipCode))
	return b.String()
}

// Longitude 生成经度
func Longitude() float64 {
	return toFixed(randFloat64(-180, 180), 6)
}

// Latitude 生成维度
func Latitude() float64 {
	return toFixed(randFloat64(-90, 90), 6)
}

// LongitudeAndLatitude 生成经纬度坐标
func LongitudeAndLatitude() string {
	return fmt.Sprintf("%f, %f", Longitude(), Latitude())
}

func street(loc *locale) string {
	var b strings.Builder
	if loc.Address.Region.Reverse {
		b.WriteString(NumberPattern(pick(loc.Address.StreetNumber)))
		b.WriteString(loc.Text.RandomTitle(1, 2))
		b.WriteByte(0x20)
		b.WriteString(pick(loc.Address.StreetSuffix))
	} else {
		b.WriteString(loc.Text.RandomTitle(1, 2))
		b.WriteByte(0x20)
		b.WriteString(pick(loc.Address.StreetSuffix))
		b.WriteString(NumberPattern(pick(loc.Address.StreetNumber)))
	}
	return b.String()
}

type i18nAddress struct {
	Region       i18nRegion `json:"region"`
	StreetSuffix []string   `json:"streetSuffix"`
	StreetNumber []string   `json:"streetNumber"`
	ZipCode      string     `json:"zipcode"`
}

type i18nRegion struct {
	Data    []regionItem `json:"data"`
	Reverse bool         `json:"reverse"`
	Split   string       `json:"split"`
}

func (r *i18nRegion) RandomCity() string {
	return pick(pick(r.Data).Citys)
}

func (r *i18nRegion) RandomStateOrProvince() string {
	return pick(r.Data).Name
}

func (r *i18nRegion) RandomString() string {
	p := pick(r.Data)
	c := pick(p.Citys)
	if r.Reverse {
		return fmt.Sprintf("%s%s%s", c, r.Split, p.Name)
	} else {
		return fmt.Sprintf("%s%s%s", p.Name, r.Split, c)
	}
}

type regionItem struct {
	Name  string   `json:"id"`
	Citys []string `json:"citys"`
}
