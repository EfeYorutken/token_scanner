package main

import (
	"fmt"
	"net"
	"slices"
	"strconv"
	"sync"
	"time"
)

type Target struct {
	name string
	ports []string
	protocols []string
} 

func scan_address_on_port(target string, port string, protocol string) string{
	target_link := target + ":" + port
	connection, err := net.DialTimeout(protocol, target_link, 60 * time.Second)

	if err == nil{
		connection.Close()
		return fmt.Sprintf("SUCCESSFUL SCAN ON %s:%s with %s", target, port, protocol)
	}
		return fmt.Sprintf("FAILED SCAN ON %s:%s with %s", target, port, protocol)
}

func (t Target) get_target_w_ports(ps []string) Target{
	var ports []string
	if ps == nil{
		for i := 1; i < 1000; i++{
			ports = append(ports, strconv.Itoa(i))
		}
	}else{
		ports = append(t.ports, ps...)
	}
	return Target{t.name, ports, t.protocols}
}

func (t Target) get_target_w_protocols(prots []string) Target{
	protocols := ternary(prots == nil, []string{"tcp"}, append(t.ports, prots...))
	return Target{t.name, t.ports, protocols}
}

func NewTarget(n string, p []string, prot []string) Target{
	return Target{n,nil,nil}.get_target_w_ports(p).get_target_w_protocols(prot)
}

//you need to use either a mutex or a waitgroup
func (target Target)ScanAddressOnPorts() []string{
	target_name := target.name
	ports := target.ports
	prots := target.protocols
	res := slices.Repeat([]string{""}, len(ports) * len(prots))
	
	protocol_count := len(prots)

	wg := sync.WaitGroup{}
	wg.Add(len(ports) * len(prots))

	for port_index, port := range ports{
		for protocol_index, protocol := range prots{
			go func(){
				fn_res := scan_address_on_port(target_name, port, protocol)
				res[port_index * protocol_count + protocol_index] = fn_res
				wg.Done()
			}()
		}
	}

	wg.Wait()

	return res
}
