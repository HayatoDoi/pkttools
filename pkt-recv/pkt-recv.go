package main

import (
	"../hex"
	"fmt"
	"github.com/HayatoDoi/pkt"
	flags "github.com/jessevdk/go-flags"
	"os"
)

type Options struct {
	InterfaceName string `short:"I" required:"true" description:"Set interface name."`
}

var opts Options

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "See the output of pkt-recv -h for a summary of options.\n")
		os.Exit(1)
	}

	conn, err := pkt.NewPFConn(opts.InterfaceName)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	buf := make([]byte, 5120)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		hex.Dump(buf[:n])
	}
}
