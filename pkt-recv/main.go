package main

import (
	"fmt"
	"github.com/HayatoDoi/pkttools/pkt"
	"os"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(pkt.Htons(syscall.ETH_P_ALL)))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	f := os.NewFile(uintptr(fd), fmt.Sprintf("fd %d", fd))

	for {
		buf := make([]byte, 1024)
		numRead, err := f.Read(buf)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		fmt.Printf("% X\n", buf[:numRead])
	}
}
