package datagen

import (
	"strings"
	"sync"
	"time"
)

func Date(format ...string) string {
	d := time.Date(
		randInt(1949, 2023),
		time.Month(randInt(1, 12)),
		randInt(1, 31),
		0, 0, 0, 0,
		time.Local,
	)
	return d.Format(getGoDateLayout("2006-01-02", format...))
}

func Time(format ...string) string {
	return time.Unix(int64(grand.Intn(3600*24)), 0).Format(getGoDateLayout("15:04:05", format...))
}

func DateTime(format ...string) string {
	d := time.Now().Add(-time.Duration(grand.Int63n(3600*24*365*10)) * time.Second)
	return d.Format(getGoDateLayout(time.RFC3339, format...))
}

func Timestamp() int64 {
	return time.Now().Unix() - int64(grand.Int63n(9999999))
}

func Now(format ...string) string {
	return time.Now().Format(getGoDateLayout(time.RFC3339, format...))
}

func getGoDateLayout(defaultLayout string, layout ...string) string {
	if len(layout) == 0 {
		return defaultLayout
	}
	l := pick(layout)
	if l == "" {
		return defaultLayout
	}
	v, ok := goDateLayoutMappingCache.Load(l)
	if ok {
		return v.(string)
	}
	var buf strings.Builder
	ls := []rune(l)
	n := len(ls)
	for i := 0; i < n; i++ {
		if !goDateLayoutTokens[ls[i]] {
			buf.WriteRune(ls[i])
			continue
		}
		j, token := i+1, ls[i]
		for ; j < n; j++ {
			if ls[j] != token {
				break
			}
		}
		buf.WriteString(goDateLayoutMapping(ls[i:j]))
		i = j - 1
	}
	s := buf.String()
	goDateLayoutMappingCache.Store(l, s)
	return s
}

func goDateLayoutMapping(token []rune) string {
	s := token
	if len(s) > 4 {
		s = s[:4]
	}
	if v, ok := goDateLayoutKeywords[string(s)]; ok {
		return v
	}
	return string(token)
}

var (
	goDateLayoutMappingCache = &sync.Map{}
	goDateLayoutTokens       = func() map[rune]bool {
		m := make(map[rune]bool)
		for _, v := range "yYMDdEaHhmsSzZX" {
			m[v] = true
		}
		return m
	}()

	// java format
	goDateLayoutKeywords = map[string]string{
		"YYYY": "2006",
		"yyyy": "2006",
		"YY":   "06",
		"yy":   "06",
		"MMMM": "January",
		"MMM":  "Jan",
		"MM":   "01",
		"M":    "1",
		"DDD":  "002",
		"dd":   "02",
		"d":    "2",
		"EEEE": "Monday",
		"EEE":  "Mon",
		"HH":   "15",
		"hh":   "03",
		"h":    "3",
		"mm":   "04",
		"m":    "4",
		"ss":   "05",
		"s":    "5",
		"SSS":  "000",
		"a":    "PM",
		"z":    "MST",
		"Z":    "-0700",
		"X":    "Z07",
		"XX":   "Z0700",
		"XXX":  "Z07:00",
	}
)
