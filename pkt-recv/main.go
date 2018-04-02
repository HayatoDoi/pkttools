package main

import (
	"fmt"
	"github.com/HayatoDoi/pkttools/pkt"
	"os"
	"syscall"
	"unsafe"
)

const (
	SizeOfIfReq = 40
	IFNAMSIZ    = 16
)

type ifReq struct {
	Name  [IFNAMSIZ]byte
	Flags uint16
	pad   [SizeOfIfReq - IFNAMSIZ - 2]byte
}

func main() {
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(pkt.Htons(syscall.ETH_P_ALL)))
	// _, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(pkt.Htons(syscall.ETH_P_ALL)))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var ifr ifReq
	fmt.Println(ifr)
	copy(ifr.Name[:15], "br-123293634eeb")
	fmt.Println(ifr)
	err = pkt.Ioctl(uintptr(fd), uintptr(syscall.SIOCGIFINDEX), uintptr(unsafe.Pointer(&ifr)))
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
		fmt.Printf("%X\n", buf[:numRead])
	}
}
