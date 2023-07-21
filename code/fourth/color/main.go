package main

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

var (
	testcolor = 1725515
)

func main() {
	fmt.Print("int2rbg测试结果:")
	fmt.Println(int2rbg(int64(testcolor)))
	colorHex := strconv.FormatInt(int64(testcolor), 16)
	fmt.Println("十六进制测试结果:", colorHex)
	fmt.Print("parseHexColor测试结果:")
	fmt.Println(parseHexColor(colorHex))
}

func int2rbg(t int64) (int64, int64, int64) {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], uint64(t))
	b, g, r := int64(buf[0]), int64(buf[1]), int64(buf[2])
	return r, g, b
}

func parseHexColor(x string) (r, g, b, a int) {
	x = strings.TrimPrefix(x, "#")
	a = 255
	if len(x) == 3 {
		format := "%1x%1x%1x"
		fmt.Sscanf(x, format, &r, &g, &b)
		r |= r << 4
		g |= g << 4
		b |= b << 4
	}
	if len(x) == 6 {
		format := "%02x%02x%02x"
		fmt.Sscanf(x, format, &r, &g, &b)
	}
	if len(x) == 8 {
		format := "%02x%02x%02x%02x"
		fmt.Sscanf(x, format, &r, &g, &b, &a)
	}
	return
}
