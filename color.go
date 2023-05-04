package datagen

import (
	"fmt"
	"image/color"
)

// Color 生成web css色值
func Color(typ string) string {
	c := randColor()
	switch typ {
	case "rgb":
		return fmt.Sprintf("rgb(%d,%d,%d)", c.R, c.G, c.B)
	case "rgba":
		return fmt.Sprintf("rgba(%d,%d,%d,%.2f)", c.R, c.G, c.B, grand.Float32())
	case "hsl":
		return fmt.Sprintf("hsl(%d,%d%%,%d%%)", grand.Intn(360), grand.Intn(100), grand.Intn(100))
	default:
		// hex
		return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
	}
}

func randColor() color.RGBA {
	return color.RGBA{
		uint8(grand.Intn(255)),
		uint8(grand.Intn(255)),
		uint8(grand.Intn(255)),
		0xff,
	}
}
