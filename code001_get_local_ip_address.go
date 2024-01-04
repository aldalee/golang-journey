package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// get all network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// traverse all network interfaces
	for _, iface := range interfaces {
		// exclude some special interfaces
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		// obtain the address information of the interface
		addrs, err := iface.Addrs()
		if err != nil {
			log.Printf("Error: %v\n", err)
			continue
		}

		// traverse the address of the interface
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				fmt.Println(v.IP)
			case *net.IPAddr:
				fmt.Println(v.IP)
			}
		}
	}
}
