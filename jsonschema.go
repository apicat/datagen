package datagen

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type JSONSchemaOption struct {
	Datagen   string
	SkipError bool
	// json/xml defaut json
	// OutputFormat string
}

func JSONSchemaGen(data []byte, opt *JSONSchemaOption) ([]byte, error) {
	if opt == nil {
		opt = &JSONSchemaOption{}
	}
	if opt.Datagen == "" {
		opt.Datagen = "x-datagen"
	}
	b := &jsonschemaBuilder{
		src: data,
		opt: *opt,
	}
	var obj JSchema
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, err
	}
	result, err := b.gen(obj)
	if err != nil {
		return nil, err
	}
	return json.Marshal(result)
}

type jsonschemaBuilder struct {
	src []byte
	opt JSONSchemaOption
}

func (j *jsonschemaBuilder) gen(obj JSchema) (any, error) {
	if obj["enum"] != nil {
		var enums []any
		if err := json.Unmarshal(obj["enum"], &enums); err != nil {
			if j.opt.SkipError {
				return nil, err
			}
		} else {
			return pick(enums), nil
		}
	}
	// type 有可能是数组
	typ, ok := obj["type"]
	if !ok {
		return nil, fmt.Errorf("not type")
	}
	var jsType string
	if typ[0] == '"' {
		json.Unmarshal(typ, &jsType) // nolint:errcheck
	} else {
		var types []string
		json.Unmarshal(typ, &types) // nolint:errcheck
		jsType = pick(types)
	}
	switch jsType {
	case "string":
		return j.toString(obj)
	case "object":
		return j.toObject(obj)
	case "array":
		return j.toArray(obj)
	case "integer", "number":
		return j.toNumber(obj, jsType)
	case "boolean":
		return Boolean(), nil
	case "null":
		return nil, nil
	}
	if j.opt.SkipError {
		return nil, nil
	}
	return nil, fmt.Errorf("not support type %s", jsType)
}

type JSchema map[string]json.RawMessage

func (j *jsonschemaBuilder) toString(obj JSchema) (any, error) {
	if raw, ok := obj[j.opt.Datagen]; ok {
		// 强制转化格式 避免结果不是string
		return fmt.Sprintf("%v", CallFunction(j.rawString(raw))), nil
	}
	if raw, ok := obj["pattern"]; ok {
		return Regexp(j.rawString(raw)), nil
	}
	if raw, ok := obj["format"]; ok {
		return j.stringFormat(j.rawString(raw)), nil
	}
	min, max := j.rangeint(obj, 5, 20, "minLength", "maxLength")
	return String("ansic", min, max), nil
}

func (j *jsonschemaBuilder) stringFormat(format string) string {
	switch format {
	case "date-time":
		return DateTime()
	case "time":
		return Time()
	case "date":
		return Date()
	case "email":
		return Email()
	case "idn-email":
		return Email()
	case "hostname":
		return Typography().Word(4, 10)
	case "idn-hostname":
		return Typography("en", "zh").Word(4, 10)
	case "ipv4":
		return IPv4()
	case "ipv6":
		return IPv6()
	case "uuid":
		return UUID()
	case "uri":
		return URL()
	case "uri-reference":
		return URL() + "#" + String("ansic", 3, 5)
	}
	return String("ansic", 6, 10)
}

func (j *jsonschemaBuilder) rawString(b json.RawMessage) string {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return ""
	}
	return s
}

func (j *jsonschemaBuilder) toNumber(obj JSchema, typ string) (any, error) {
	if raw, ok := obj[j.opt.Datagen]; ok {
		v := CallFunction(j.rawString(raw))
		switch x := v.(type) {
		case int, int64:
			return x, nil
		case float64:
			if typ == "number" {
				return x, nil
			}
		case string:
			if typ == "number" {
				if ret, err := strconv.ParseFloat(x, 64); err == nil {
					return ret, nil
				}
			} else {
				if ret, err := strconv.ParseInt(x, 10, 64); err == nil {
					return ret, nil
				}
			}
		}
	}
	min, max := j.rangeint(obj, 10, 100000, "minimum", "maximum")
	if typ == "integer" {
		return Integer(min, max), nil
	} else {
		return Float(float64(min), float64(max), 4), nil
	}
}

func (j *jsonschemaBuilder) toObject(obj JSchema) (any, error) {
	x := make(map[string]any)
	props, ok := obj["properties"]
	if ok {
		var ps map[string]JSchema
		if err := json.Unmarshal(props, &ps); err != nil {
			if j.opt.SkipError {
				return x, nil
			}
			return nil, err
		}
		for k, v := range ps {
			p, err := j.gen(v)
			if err != nil {
				if j.opt.SkipError {
					continue
				}
				return nil, err
			}
			x[k] = p
		}
	}
	return x, nil
}

func (j *jsonschemaBuilder) toArray(obj JSchema) (any, error) {
	x := make([]any, 0)
	var items JSchema
	if raw, ok := obj["items"]; ok {
		if !ok {
			return x, nil
		}
		if err := json.Unmarshal(raw, &items); err != nil {
			if j.opt.SkipError {
				return x, nil
			}
			return nil, err
		}
	}
	min, max := j.rangeint(obj, 3, 10, "minItems", "maxItems")
	n := randInt(min, max)
	for i := 0; i < n; i++ {
		v, err := j.gen(items)
		if err != nil {
			if j.opt.SkipError {
				continue
			}
			return nil, err
		}
		x = append(x, v)
	}
	return x, nil
}

func (j *jsonschemaBuilder) rangeint(src JSchema, defaultmin, defaultmax int, minkey, maxkey string) (int, int) {
	min, max := defaultmin, defaultmax
	if raw, ok := src[minkey]; ok {
		json.Unmarshal(raw, &min) // nolint:errcheck
	}
	if raw, ok := src[maxkey]; ok {
		json.Unmarshal(raw, &max) // nolint:errcheck
	}
	return min, max
}