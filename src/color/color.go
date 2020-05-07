package color

import colorful "github.com/lucasb-eyer/go-colorful"

func Nearest(hex uint32) int {
	target := colorful.Color{
		R: float64(hex>>0&0xFF) / float64(0xFF),
		G: float64(hex>>8&0xFF) / float64(0xFF),
		B: float64(hex>>16&0xFF) / float64(0xFF),
	}
	maxDistance := 1.0
	var result int
	for i, irc := range colors {
		if distance := target.DistanceLab(irc); distance < maxDistance {
			maxDistance = distance
			result = i
		}
	}
	return result
}

var colors = [...]colorful.Color{
	colorful.Color{R: 0xff / 0xFF, G: 0xff / 0xFF, B: 0xff / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0x00 / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0x00 / 0xFF, B: 0x7f / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0x93 / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0xff / 0xFF, G: 0x00 / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x7f / 0xFF, G: 0x00 / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x9c / 0xFF, G: 0x00 / 0xFF, B: 0x9c / 0xFF},
	colorful.Color{R: 0xfc / 0xFF, G: 0x7f / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0xff / 0xFF, G: 0xff / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0xfc / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0x93 / 0xFF, B: 0x93 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0xff / 0xFF, B: 0xff / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0x00 / 0xFF, B: 0xfc / 0xFF},
	colorful.Color{R: 0xff / 0xFF, G: 0x00 / 0xFF, B: 0xff / 0xFF},
	colorful.Color{R: 0x7f / 0xFF, G: 0x7f / 0xFF, B: 0x7f / 0xFF},
	colorful.Color{R: 0xd2 / 0xFF, G: 0xd2 / 0xFF, B: 0xd2 / 0xFF},
	colorful.Color{R: 0x47 / 0xFF, G: 0x00 / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x47 / 0xFF, G: 0x21 / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x47 / 0xFF, G: 0x47 / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x32 / 0xFF, G: 0x47 / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0x47 / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0x47 / 0xFF, B: 0x2c / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0x47 / 0xFF, B: 0x47 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0x27 / 0xFF, B: 0x47 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0x00 / 0xFF, B: 0x47 / 0xFF},
	colorful.Color{R: 0x2e / 0xFF, G: 0x00 / 0xFF, B: 0x47 / 0xFF},
	colorful.Color{R: 0x47 / 0xFF, G: 0x00 / 0xFF, B: 0x47 / 0xFF},
	colorful.Color{R: 0x47 / 0xFF, G: 0x00 / 0xFF, B: 0x2a / 0xFF},
	colorful.Color{R: 0x74 / 0xFF, G: 0x00 / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x74 / 0xFF, G: 0x3a / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x74 / 0xFF, G: 0x74 / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x51 / 0xFF, G: 0x74 / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0x74 / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0x74 / 0xFF, B: 0x49 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0x74 / 0xFF, B: 0x74 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0x40 / 0xFF, B: 0x74 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0x00 / 0xFF, B: 0x74 / 0xFF},
	colorful.Color{R: 0x4b / 0xFF, G: 0x00 / 0xFF, B: 0x74 / 0xFF},
	colorful.Color{R: 0x74 / 0xFF, G: 0x00 / 0xFF, B: 0x74 / 0xFF},
	colorful.Color{R: 0x74 / 0xFF, G: 0x00 / 0xFF, B: 0x45 / 0xFF},
	colorful.Color{R: 0xb5 / 0xFF, G: 0x00 / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0xb5 / 0xFF, G: 0x63 / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0xb5 / 0xFF, G: 0xb5 / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x7d / 0xFF, G: 0xb5 / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0xb5 / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0xb5 / 0xFF, B: 0x71 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0xb5 / 0xFF, B: 0xb5 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0x63 / 0xFF, B: 0xb5 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0x00 / 0xFF, B: 0xb5 / 0xFF},
	colorful.Color{R: 0x75 / 0xFF, G: 0x00 / 0xFF, B: 0xb5 / 0xFF},
	colorful.Color{R: 0xb5 / 0xFF, G: 0x00 / 0xFF, B: 0xb5 / 0xFF},
	colorful.Color{R: 0xb5 / 0xFF, G: 0x00 / 0xFF, B: 0x6b / 0xFF},
	colorful.Color{R: 0xff / 0xFF, G: 0x00 / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0xff / 0xFF, G: 0x8c / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0xff / 0xFF, G: 0xff / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0xb2 / 0xFF, G: 0xff / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0xff / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0xff / 0xFF, B: 0xa0 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0xff / 0xFF, B: 0xff / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0x8c / 0xFF, B: 0xff / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0x00 / 0xFF, B: 0xff / 0xFF},
	colorful.Color{R: 0xa5 / 0xFF, G: 0x00 / 0xFF, B: 0xff / 0xFF},
	colorful.Color{R: 0xff / 0xFF, G: 0x00 / 0xFF, B: 0xff / 0xFF},
	colorful.Color{R: 0xff / 0xFF, G: 0x00 / 0xFF, B: 0x98 / 0xFF},
	colorful.Color{R: 0xff / 0xFF, G: 0x59 / 0xFF, B: 0x59 / 0xFF},
	colorful.Color{R: 0xff / 0xFF, G: 0xb4 / 0xFF, B: 0x59 / 0xFF},
	colorful.Color{R: 0xff / 0xFF, G: 0xff / 0xFF, B: 0x71 / 0xFF},
	colorful.Color{R: 0xcf / 0xFF, G: 0xff / 0xFF, B: 0x60 / 0xFF},
	colorful.Color{R: 0x6f / 0xFF, G: 0xff / 0xFF, B: 0x6f / 0xFF},
	colorful.Color{R: 0x65 / 0xFF, G: 0xff / 0xFF, B: 0xc9 / 0xFF},
	colorful.Color{R: 0x6d / 0xFF, G: 0xff / 0xFF, B: 0xff / 0xFF},
	colorful.Color{R: 0x59 / 0xFF, G: 0xb4 / 0xFF, B: 0xff / 0xFF},
	colorful.Color{R: 0x59 / 0xFF, G: 0x59 / 0xFF, B: 0xff / 0xFF},
	colorful.Color{R: 0xc4 / 0xFF, G: 0x59 / 0xFF, B: 0xff / 0xFF},
	colorful.Color{R: 0xff / 0xFF, G: 0x66 / 0xFF, B: 0xff / 0xFF},
	colorful.Color{R: 0xff / 0xFF, G: 0x59 / 0xFF, B: 0xbc / 0xFF},
	colorful.Color{R: 0xff / 0xFF, G: 0x9c / 0xFF, B: 0x9c / 0xFF},
	colorful.Color{R: 0xff / 0xFF, G: 0xd3 / 0xFF, B: 0x9c / 0xFF},
	colorful.Color{R: 0xff / 0xFF, G: 0xff / 0xFF, B: 0x9c / 0xFF},
	colorful.Color{R: 0xe2 / 0xFF, G: 0xff / 0xFF, B: 0x9c / 0xFF},
	colorful.Color{R: 0x9c / 0xFF, G: 0xff / 0xFF, B: 0x9c / 0xFF},
	colorful.Color{R: 0x9c / 0xFF, G: 0xff / 0xFF, B: 0xdb / 0xFF},
	colorful.Color{R: 0x9c / 0xFF, G: 0xff / 0xFF, B: 0xff / 0xFF},
	colorful.Color{R: 0x9c / 0xFF, G: 0xd3 / 0xFF, B: 0xff / 0xFF},
	colorful.Color{R: 0x9c / 0xFF, G: 0x9c / 0xFF, B: 0xff / 0xFF},
	colorful.Color{R: 0xdc / 0xFF, G: 0x9c / 0xFF, B: 0xff / 0xFF},
	colorful.Color{R: 0xff / 0xFF, G: 0x9c / 0xFF, B: 0xff / 0xFF},
	colorful.Color{R: 0xff / 0xFF, G: 0x94 / 0xFF, B: 0xd3 / 0xFF},
	colorful.Color{R: 0x00 / 0xFF, G: 0x00 / 0xFF, B: 0x00 / 0xFF},
	colorful.Color{R: 0x13 / 0xFF, G: 0x13 / 0xFF, B: 0x13 / 0xFF},
	colorful.Color{R: 0x28 / 0xFF, G: 0x28 / 0xFF, B: 0x28 / 0xFF},
	colorful.Color{R: 0x36 / 0xFF, G: 0x36 / 0xFF, B: 0x36 / 0xFF},
	colorful.Color{R: 0x4d / 0xFF, G: 0x4d / 0xFF, B: 0x4d / 0xFF},
	colorful.Color{R: 0x65 / 0xFF, G: 0x65 / 0xFF, B: 0x65 / 0xFF},
	colorful.Color{R: 0x81 / 0xFF, G: 0x81 / 0xFF, B: 0x81 / 0xFF},
	colorful.Color{R: 0x9f / 0xFF, G: 0x9f / 0xFF, B: 0x9f / 0xFF},
	colorful.Color{R: 0xbc / 0xFF, G: 0xbc / 0xFF, B: 0xbc / 0xFF},
	colorful.Color{R: 0xe2 / 0xFF, G: 0xe2 / 0xFF, B: 0xe2 / 0xFF},
	colorful.Color{R: 0xff / 0xFF, G: 0xff / 0xFF, B: 0xff / 0xFF},
}
