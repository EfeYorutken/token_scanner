package main

import (
	"strconv"
	"time"
	"net"
	"fmt"
)

type Target struct {
	name string
	ports []int
	protocols []string
} 

func scan_address_on_port(target string, port int, protocol string) string{
	connection, err := net.DialTimeout(protocol, target + ":" + strconv.Itoa(port), 60 * time.Second)

	if err == nil{
		connection.Close()
		return fmt.Sprintf("SUCCESSFUL SCAN ON %s:%d with %s", target, port, protocol)
	}
		return fmt.Sprintf("FAILED SCAN ON %s:%d with %s", target, port, protocol)
}

func  NewTarget(n string, p []int, prot []string) Target{
	name := n
	var ports []int
	var protocols []string
	if p == nil{
		for i := 1; i < 100; i++{
			ports = append(ports, i)
		}
	}else{
		ports = p
	}
	if prot == nil{
		protocols = append(protocols, "tcp")
	}
	return Target{name, ports, protocols}
}

func (target Target)ScanAddressOnPorts() []string{
	target_name := target.name
	ports := target.ports
	prots := target.protocols

	var res []string
	for i := 0; i < len(ports); i++{
		for _,prot := range prots{
		res = append(res, scan_address_on_port(target_name, ports[i], prot))
		}
	}
	return res
}

//might be redundent
func ScanAddressOnRange(target_name string, begin_port int, end_port int) []string{
	var res []string
	for port := begin_port; port <= end_port ; port++{
		res = append(res, scan_address_on_port(target_name, port, "tcp"))
	}
	return res
}
