package datagen

import (
	"fmt"
)

func IPv4() string {
	x := func() int { return grand.Intn(256) }
	return fmt.Sprintf("%d.%d.%d.%d", x(), x(), x(), x())
}

func IPv6() string {
	x := func() int { return grand.Intn(65536) }
	return fmt.Sprintf("%x:%x:%x:%x:%x:%x:%x:%x", x(), x(), x(), x(), x(), x(), x(), x())
}

func URL() string {
	return Regexp(`http(s)?://[a-z0-9]{3,10}\.(com(\.cn)?|org|net|io)(/[a-z0-9]{5,10}){0,3}`)
}

func Domain() string {
	return Regexp(`[a-z0-9]{3,10}\.(com(\.cn)|org|net|io)`)
}

func UUID() string {
	return Regexp(`[a-z0-9]{8}(-[a-z0-9]{4}){3}-[a-z0-9]{12}`)
}

func Email() string {
	return Regexp(`[a-z0-9]{5,18}@[a-z0-9]{2,5}\.(com|org|net)`)
}

var httpstatuscodes = []int{
	200, 301, 302, 400, 401, 403, 404, 405, 500, 502, 503,
}

func HTTPCode() int {
	return pick(httpstatuscodes)
}

var httpmethods = []string{
	"POST", "GET", "PUT", "PATCH", "DELETE", "OPTION",
}

func HTTPMethod() string {
	return pick(httpmethods)
}
