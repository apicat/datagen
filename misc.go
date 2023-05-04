package datagen

func Boolean(v ...bool) bool {
	switch len(v) {
	case 0:
		return randInt(0, 1) == 1
	case 1:
		return v[0]
	default:
		return pick(v)
	}
}

// Float
// v [min,max,fixed]
// n = len(v)
// n == 0 :random
// n == 1 :max == min
func Float(v ...float64) float64 {
	switch len(v) {
	case 0:
		return grand.Float64()
	case 1:
		return v[0]
	case 2:
		return randFloat64(v[0], v[1])
	default:
		return toFixed(randFloat64(v[0], v[1]), int(v[2]))
	}
}

func Integer(v ...int) int {
	switch len(v) {
	case 0:
		return grand.Int()
	case 1:
		return v[0]
	default:
		return randInt(v[0], v[len(v)-1])
	}
}

// OneOf 从数组中随机取一个值
func OneOf(v ...any) any {
	return pick(v)
}

// NumberPattern 数字字符串模式
// 使用#替代数字
// example:
// `(###)##-###` -> `(219)10-231`
func NumberPattern(p string) string {
	if p == "" {
		return p
	}
	bs := []byte(p)
	for i := 0; i < len(bs); i++ {
		if bs[i] == '#' {
			bs[i] = byte(randDigit())
		}
	}
	if bs[0] == '0' {
		bs[0] = byte(randInt(1, 9))
	}
	return string(bs)
}