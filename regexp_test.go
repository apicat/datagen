package datagen

import (
	"testing"
)

func BenchmarkRegexp(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			Regexp(`[a-z0-9]{8}(-[a-z0-9]){3}-[a-z0-9]{12}`)
		}
	})
}
