package hex

import (
	"bytes"
	"fmt"
)

func byte2string(b []byte) string {
	var r string
	n := len(b)
	var i, j int
	for i = 0; i < n; i += 16 {
		r += fmt.Sprintf("%06x : ", i)
		for j = 0; j < 16; j++ {
			if i+j >= n {
				r += fmt.Sprintf("   ")
			} else {
				r += fmt.Sprintf("%2.2X ", b[i+j])
			}
			if (j+1)%4 == 0 {
				r += fmt.Sprintf(" ")
			}
		}
		r += fmt.Sprintf(": ")
		for j = 0; j < 16 && i+j < n; j++ {
			bs := []byte{b[i+j]}
			str := string(bs)
			if bytes.Compare(bs, []byte{0x21}) == -1 || bytes.Compare(bs, []byte{0x7e}) == 1 {
				str = "."
			}
			r += fmt.Sprintf("%s", str)
		}
		r += fmt.Sprintf("\n")
	}
	return r
}

/*
 * Dump is print packet.
 */
func Dump(b []byte) {
	fmt.Printf(byte2string(b))
}
