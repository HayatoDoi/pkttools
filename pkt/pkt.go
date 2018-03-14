package pkt

func Htons(n uint16) uint16 {
	var (
		high uint16 = n >> 8
		ret  uint16 = n<<8 + high
	)
	return ret
}
