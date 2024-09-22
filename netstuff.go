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
	target_link := target + ":" + strconv.Itoa(port)
	connection, err := net.DialTimeout(protocol, target_link, 60 * time.Second)

	if err == nil{
		connection.Close()
		return fmt.Sprintf("SUCCESSFUL SCAN ON %s:%d with %s", target, port, protocol)
	}
		return fmt.Sprintf("FAILED SCAN ON %s:%d with %s", target, port, protocol)
}

func (t Target) get_target_w_ports(ps []int) Target{
	var ports []int
	if ps == nil{
		for i := 1; i < 1000; i++{
			ports = append(ports, i)
		}
	}else{
		ports = ps
	}
	return Target{t.name, ports, t.protocols}
}

func (t Target) get_target_w_protocols(prots []string) Target{
	protocols := ternary(prots == nil, []string{"tcp"}, prots)
	return Target{t.name, t.ports, protocols}
}

func NewTarget(n string, p []int, prot []string) Target{
	return Target{n,nil,nil}.get_target_w_ports(p).get_target_w_protocols(prot)
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
