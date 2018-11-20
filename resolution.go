package resparse

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
ParseResolution converts a resolution string into a pair of pixel sizes.

E.G
"640x480" -> 640, 480, nil
"1600,1200" -> 1600, 1200, nil
"640-480" -> 640, 480, nil
"HD" -> 1920, 1080, nil
"1080p" -> 1920, 1080, nil
"4K" -> 3840, 2160, nil
"UHD" -> 3840, 2160, nil
*/
func ParseResolution(in string) (int, int, error) {

	start := -1
	end := -1
	sepStart := -1
	sepEnd := -1

	// Only do the upper casing if it is needed.
	needUpper := false

	// make a single pass through the string to work out start, end and seperator points.
	// This makes things slightly faster than multiple calls to indexof, lastindex and trim
	for i, c := range in {
		if !unicode.IsSpace(c) {
			if start == -1 {
				start = i
			}
			end = i
		}

		if c == 'X' || c == 'x' || c == ',' || c == ' ' || c == '|' || c == '*' {
			if sepStart == -1 {
				if start != -1 {
					sepStart = i
				}
			} else {
				if sepEnd == -1 || sepEnd == i-1 {
					sepEnd = i
				}
			}
		}
		if unicode.IsLower(c) {
			needUpper = true
		}
	}

	if start == -1 || end == -1 {
		return -1, -1, fmt.Errorf("could not parse \"%v\" as a resolution", in)
	}

	upper := in[start : end+1]
	if needUpper {
		upper = strings.ToUpper(upper)
	}

	if v, ok := known[upper]; ok {
		return v.x, v.y, nil
	}

	// if it is not in our lookup table then try and split the string
	if sepStart == -1 || sepStart == start || sepEnd == end {
		return -1, -1, fmt.Errorf("could not parse \"%v\" as a resolution", in)
	}

	if sepStart != -1 && sepEnd == -1 {
		sepEnd = sepStart
	}

	x, err := strconv.Atoi(in[start:sepStart])
	if err != nil {
		return -1, -1, fmt.Errorf("could not parse \"%v\" as a resolution", in)
	}

	y, err := strconv.Atoi(in[sepEnd+1 : end+1])
	if err != nil {
		return -1, -1, fmt.Errorf("could not parse \"%v\" as a resolution", in)
	}

	return x, y, nil

}

type res struct {
	x int
	y int
}

var known = map[string]res{
	"1080I":   {1920, 1080},
	"1080P":   {1920, 1080},
	"3K":      {3000, 2000},
	"480I":    {440, 480},
	"4K UHD":  {3840, 2160},
	"4K":      {3840, 2160},
	"5K":      {5120, 2880},
	"720P":    {1280, 720},
	"BLU-RAY": {1920, 1080},
	"EGA":     {640, 350},
	"FHD":     {1920, 1080},
	"FHD+":    {1920, 1280},
	"HD":      {1366, 768},
	"HD+":     {1600, 900},
	"HDTV":    {1920, 1080},
	"IMAX":    {5616, 4096},
	"NTSC":    {440, 480},
	"PAL":     {520, 576},
	"QHD":     {2560, 1440},
	"QWXGA":   {2048, 1152},
	"QXGA":    {2048, 1536},
	"SECAM":   {520, 576},
	"SVGA":    {800, 600},
	"SXGA":    {1280, 1024},
	"UHDTV":   {3840, 2160},
	"UXGA":    {1600, 1200},
	"WQHD":    {2560, 1440},
	"WSVGA":   {1024, 600},
	"WSXGA+":  {1680, 1050},
	"WUXGA":   {1920, 1200},
	"WXGA":    {1280, 768},
	"WXGA+":   {1440, 900},
	"XGA":     {1024, 768},
	"XGA+":    {1152, 864},
}
