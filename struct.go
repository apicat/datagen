package datagen

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type GenOption struct {
	DatagenKey string
	SkipError  bool
	// json/xml defaut json
	// OutputFormat string
}

func createGenOption(key string) *GenOption {
	return &GenOption{
		DatagenKey: key,
	}
}

func StructGen(v any, opt *GenOption) ([]byte, error) {
	if opt == nil {
		opt = createGenOption("datagen")
	}
	b := &structBuilder{
		opt: opt,
	}
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("the value must be a struct")
	}
	result, err := b.gen(t, "")
	if err != nil {
		return nil, err
	}
	return json.Marshal(result)
}

type structBuilder struct {
	opt *GenOption
}

func (s *structBuilder) toNumber(fn string) int64 {
	var v int64
	if fn != "" {
		if ret := CallFunction(fn); ret != nil {
			switch x := ret.(type) {
			case int64:
				v = x
			case float64:
				v = int64(x)
			case bool:
				if x {
					v = 1
				}
			default:
				pv, err := strconv.ParseInt(fmt.Sprintf("%v", ret), 10, 64)
				if err == nil {
					v = pv
				}
			}
		}
	} else {
		v = Integer()
	}
	return v
}

func (s *structBuilder) tagName(field reflect.StructField) string {
	v := field.Tag.Get("json")
	ps := strings.Split(v, ",")
	if ps[0] != "" {
		return ps[0]
	}
	return field.Name
}

func (s *structBuilder) gen(tt reflect.Type, fn string) (any, error) {
	switch tt.Kind() {
	case reflect.String:
		if fn != "" {
			if v := CallFunction(fn); v != nil {
				if x, ok := v.(string); ok {
					return x, nil
				}
				return fmt.Sprintf("%v", v), nil
			}
		}
		return String("letter", 10, 20), nil
	case reflect.Float32, reflect.Float64:
		var v float64
		if fn != "" {
			if ret := CallFunction(fn); ret != nil {
				switch x := ret.(type) {
				case float64:
					v = x
				case int64:
					v = float64(x)
				case bool:
					if x {
						v = 1
					}
				default:
					pv, err := strconv.ParseFloat(fmt.Sprintf("%v", ret), 64)
					if err != nil {
						if !s.opt.SkipError {
							return nil, err
						}
					} else {
						v = pv
					}
				}
			}
		} else {
			v = Float()
		}
		x := reflect.New(tt)
		x.Elem().SetFloat(v)
		return x.Interface(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v := s.toNumber(fn)
		x := reflect.New(tt)
		x.Elem().SetUint(uint64(v))
		return x.Interface(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v := s.toNumber(fn)
		x := reflect.New(tt)
		x.Elem().SetInt(v)
		return x.Interface(), nil
	case reflect.Bool:
		return Boolean(), nil
	case reflect.Array, reflect.Slice:
		var n int
		if tt.Kind() == reflect.Array {
			n = tt.Len()
		} else {
			n = randInt(5, 10)
		}
		item := tt.Elem()
		list := make([]any, 0)
		for i := 0; i < n; i++ {
			v, err := s.gen(item, fn)
			if err != nil {
				if s.opt.SkipError {
					continue
				} else {
					return nil, err
				}
			}
			list = append(list, v)
		}
		return list, nil
	case reflect.Struct:
		list := map[string]any{}
		for i := 0; i < tt.NumField(); i++ {
			v := tt.Field(i)
			if v.Anonymous {
				for j := 0; j < v.Type.NumField(); j++ {
					ano := v.Type.Field(i)
					a, err := s.gen(ano.Type, ano.Tag.Get(s.opt.DatagenKey))
					if err != nil && !s.opt.SkipError {
						return nil, err
					}
					list[s.tagName(ano)] = a
				}
			} else {
				a, err := s.gen(v.Type, v.Tag.Get(s.opt.DatagenKey))
				if err != nil && !s.opt.SkipError {
					return nil, err
				}
				list[s.tagName(v)] = a
			}
		}
		return list, nil
	case reflect.Map:
		list := make(map[string]any)
		if tt.Key().Kind() == reflect.String {
			n := randInt(5, 10)
			for i := 0; i < n; i++ {
				list[String("letter")], _ = s.gen(tt.Elem(), fn)
			}
		}
		return list, nil
	default:
		panic(fmt.Sprintf("unsupport type %s", tt.Kind()))
	}
}
