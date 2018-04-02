package pkt

import (
	"syscall"
)

func Htons(n uint16) uint16 {
	var (
		high uint16 = n >> 8
		ret  uint16 = n<<8 + high
	)
	return ret
}

func Ioctl(fd, op, arg uintptr) error {
	_, _, ep := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(op), arg)
	if ep != 0 {
		return syscall.Errno(ep)
	}
	return nil
}
