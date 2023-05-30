package datagen

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type FuncHandler func(Param) any

var functions map[string]FuncHandler

func RegisterFunction(name string, handler FuncHandler) {
	if _, ok := functions[name]; ok {
		panic(fmt.Sprintf("func %s alreay exist", name))
	}
	functions[name] = handler
}

func CallFunction(fn string) any {
	f := ParseFunction(fn)
	if h, ok := functions[f.Name]; ok {
		return h(f.Param)
	}
	return nil
}

func init() {
	functions = map[string]FuncHandler{
		"string":              func(p Param) any { return String(p.Args.At(0), p.Args.SliceInt(1)...) },
		"boolean":             func(p Param) any { return Boolean(p.Args.SliceBool()...) },
		"integer":             func(p Param) any { return Integer(p.Args.SliceInt64()...) },
		"float":               func(p Param) any { return Float(p.Args.SliceFloat()...) },
		"regexp":              func(p Param) any { return Regexp(p.Args.At(0)) },
		"oneof":               func(p Param) any { return OneOf(p.Args.SliceAny()...) },
		"autoincrement":       func(p Param) any { return autoIncrement(p.Args.SliceInt64()...) },
		"numberpattern":       func(p Param) any { return NumberPattern(p.Args.At(0)) },
		"word":                func(p Param) any { return Typography(p.Loc...).Word(p.Args.SliceInt()...) },
		"title":               func(p Param) any { return Typography(p.Loc...).Title(p.Args.SliceInt()...) },
		"phrase":              func(p Param) any { return Typography(p.Loc...).Phrase(p.Args.SliceInt()...) },
		"sentence":            func(p Param) any { return Typography(p.Loc...).Sentence(p.Args.SliceInt()...) },
		"paragraph":           func(p Param) any { return Typography(p.Loc...).Paragraph(p.Args.SliceInt()...) },
		"markdown":            func(p Param) any { return Typography(p.Loc...).Markdown() },
		"name":                func(p Param) any { return People(p.Loc...).Name() },
		"firstname":           func(p Param) any { return People(p.Loc...).FirstName() },
		"lastname":            func(p Param) any { return People(p.Loc...).LastName() },
		"gender":              func(p Param) any { return People(p.Loc...).Gender() },
		"phone":               func(p Param) any { return People(p.Loc...).Phone() },
		"idcard":              func(p Param) any { return People(p.Loc...).IDCard() },
		"ipv4":                func(p Param) any { return IPv4() },
		"ipv6":                func(p Param) any { return IPv6() },
		"uuid":                func(p Param) any { return UUID() },
		"domain":              func(p Param) any { return Domain() },
		"url":                 func(p Param) any { return URL() },
		"httpcode":            func(p Param) any { return HTTPCode() },
		"httpmethod":          func(p Param) any { return HTTPMethod() },
		"email":               func(p Param) any { return Email() },
		"date":                func(p Param) any { return Date(p.Args...) },
		"time":                func(p Param) any { return Time(p.Args...) },
		"datetime":            func(p Param) any { return DateTime(p.Args...) },
		"timestamp":           func(p Param) any { return Timestamp() },
		"now":                 func(p Param) any { return Now(p.Args...) },
		"color":               func(p Param) any { return Color(p.Args.At(0)) },
		"imagedata":           func(p Param) any { return ImageData(p.Args.SliceInt()...) },
		"imageurl":            func(p Param) any { return ImageURL(p.Args.SliceInt()...) },
		"city":                func(p Param) any { return City(p.Loc...) },
		"provinceorstate":     func(p Param) any { return ProvinceorState(p.Loc...) },
		"provinceorstatecity": func(p Param) any { return ProvinceorStateCity(p.Loc...) },
		"street":              func(p Param) any { return Street(p.Loc...) },
		"zipcode":             func(p Param) any { return ZipCode(p.Loc...) },
		"address":             func(p Param) any { return Address(p.Loc...) },
		"longitude":           func(p Param) any { return Longitude() },
		"latitude":            func(p Param) any { return Latitude() },
		"longitudelatitude":   func(p Param) any { return LongitudeAndLatitude() },
	}
}

type Function struct {
	Name string
	Param
}

type Param struct {
	Loc  []string // locales
	Args Args
}

var parseReg = regexp.MustCompile(`(?P<func>\w+)(?P<loc>\((\w+)\))?(?P<args>\|(.*))?`)

// ParseFunction 解析function eg:funcname(locales...)|args...
func ParseFunction(function string) Function {
	var (
		ps = parseReg.FindStringSubmatch(function)
		ks = parseReg.SubexpNames()
		fn = Function{}
	)
	for index, v := range ks {
		switch v {
		case "func":
			fn.Name = ps[index]
		case "loc":
			for _, v := range strings.Split(ps[index+1], ",") {
				fn.Loc = append(fn.Loc, strings.TrimSpace(v))
			}
			// fn.Loc = strings.Split(ps[index+1],",")
		case "args":
			fn.Args = parseAargs(ps[index+1])
		}
	}
	return fn
}

func parseAargs(s string) Args {
	var strStart rune
	var j int
	var ps []string
	var ss = []rune(s)
	for i := 0; i < len(ss)-1; i++ {
		v := ss[i]
		switch v {
		case '"', '\'':
			if strStart == 0 {
				strStart = v
				j = i
			} else if strStart == v {
				strStart = 0
			}
		case '\\':
			i++
		case ',':
			if strStart == 0 {
				ps = append(ps, parseArgString(string(ss[j:i])))
				j = i + 1
			}
		}
	}
	if j < len(ss) {
		ps = append(ps, parseArgString(string(ss[j:])))
	}
	return ps
}

func parseArgString(s string) string {
	s = strings.TrimSpace(s)
	if len(s) > 0 {
		if p := s[0]; p == '"' || p == '\'' {
			return strings.ReplaceAll(s[1:len(s)-1], "\\"+string(p), string(p))
		}
	}
	return s
}

type Args []string

func (a Args) indexOk(i int) bool {
	return i >= 0 && i < len(a)
}

func (a Args) sliceStart(i ...int) int {
	var start int
	if len(i) != 0 && i[0] >= 0 {
		start = i[0]
	}
	return start
}

func (a Args) At(i int) string {
	if !a.indexOk(i) {
		return ""
	}
	return a[i]
}

func (a Args) Sub(i ...int) Args {
	start := a.sliceStart(i...)
	if !a.indexOk(start) {
		return nil
	}
	return a[start:]
}

func (a Args) Bool(i int) bool {
	if !a.indexOk(i) {
		return false
	}
	return a[i] == "true"
}

func (a Args) SliceBool(i ...int) []bool {
	s := a.Sub(i...)
	list := make([]bool, len(s))
	for k := range s {
		list[k] = s.Bool(k)
	}
	return list
}

func (a Args) Int(i int) int {
	if !a.indexOk(i) {
		return 0
	}
	v, _ := strconv.ParseInt(a[i], 10, 64)
	return int(v)
}

func (a Args) SliceInt(i ...int) []int {
	s := a.Sub(i...)
	list := make([]int, len(s))
	for k := range s {
		list[k] = s.Int(k)
	}
	return list
}

func (a Args) SliceInt64(i ...int) []int64 {
	s := a.Sub(i...)
	list := make([]int64, len(s))
	for k := range s {
		v, _ := strconv.ParseInt(s[k], 10, 64)
		list[k] = v
	}
	return list
}

func (a Args) Float(i int) float64 {
	if !a.indexOk(i) {
		return 0
	}
	v, _ := strconv.ParseFloat(a[i], 64)
	return v
}

func (a Args) SliceFloat(i ...int) []float64 {
	s := a.Sub(i...)
	list := make([]float64, len(s))
	for k := range s {
		list[k] = s.Float(k)
	}
	return list
}

func (a Args) SliceAny(i ...int) []any {
	s := a.Sub(i...)
	list := make([]any, len(s))
	for k, v := range s {
		if v == "true" || v == "false" {
			list[k] = v == "true"
			continue
		}
		if v[0] == '-' || v[0] > '0' && v[0] <= '9' {
			if strings.IndexByte(v, '.') != -1 {
				// float64 try
				x, err := strconv.ParseFloat(v, 64)
				if err == nil {
					list[k] = x
					continue
				}
			} else {
				x, err := strconv.ParseInt(v, 10, 64)
				if err == nil {
					list[k] = x
					continue
				}
			}
		}
		list[k] = v
	}
	return list
}
