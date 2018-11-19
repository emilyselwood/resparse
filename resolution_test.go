package resparse

import (
	"math/rand"
	"testing"
)

type testCase struct {
	input string
	x     int
	y     int
	err   string
}

var cases = []testCase{
	{"", -1, -1, "could not parse \"\" as a resolution"},
	{" SVGA", 800, 600, ""},
	{"WSVGA", 1024, 600, ""},
	{"XGA\t", 1024, 768, ""},
	{"XGA+", 1152, 864, ""},
	{"WXGA   ", 1280, 768, ""},
	{"SXGA", 1280, 1024, ""},
	{"HD ", 1366, 768, ""},
	{"WXGA+", 1440, 900, ""},
	{"HD+", 1600, 900, ""},
	{"WSXGA+", 1680, 1050, ""},
	{"FHD", 1920, 1080, ""},
	{"WUXGA	", 1920, 1200, ""},
	{"WQHD", 2560, 1440, ""},
	{"4K UHD", 3840, 2160, ""},
	{"800x600", 800, 600, ""},
	{"1600|1200", 1600, 1200, ""},
	{"1600           1200", 1600, 1200, ""},
	{"4000,4000", 4000, 4000, ""},
	{"4000 , 4000", 4000, 4000, ""},
	{"4000 ,4000", 4000, 4000, ""},
	{" 4000 ,4000", 4000, 4000, ""},
	{"4000 ,4000 ", 4000, 4000, ""},
	{"4000, 4000", 4000, 4000, ""},
	{"dgsfgd,4000", -1, -1, "could not parse \"dgsfgd,4000\" as a resolution"},
	{"4000,djdysi", -1, -1, "could not parse \"4000,djdysi\" as a resolution"},
}

func TestBasicParse(t *testing.T) {

	for _, test := range cases {
		t.Run(test.input, func(t *testing.T) {
			x, y, err := ParseResolution(test.input)
			if (err != nil && test.err == "") || (err != nil && err.Error() != test.err) {
				t.Errorf("wrong error from parsing \"%v\" got \"%v\" expected \"%v\"", test.input, err.Error(), test.err)
			} else if x != test.x || y != test.y {
				t.Errorf("got wrong result from parsing \"%v\" got (%v,%v) expected (%v,%v)", test.input, x, y, test.x, test.y)
			}
		})
	}
}

func BenchmarkParseResolution(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ParseResolution(cases[rand.Intn(len(cases))].input)
	}
}
