package main

import (
	"net"
)

var service map[string]net.Addr //<--global variable

func main() {
	// Unprotected global variable

}

// If the following code is called from several goroutines,
//it leads to races on the service map.
//Concurrent reads and writes of the same map are not safe:

func RegisterService(name string, addr net.Addr) {
	service[name] = addr
}

func LookupService(name string) net.Addr {
	return service[name]
}

//To make the code safe, protect the accesses with a mutex:
/*
var (
	service   map[string]net.Addr
	serviceMu sync.Mutex
)

func RegisterService(name string, addr net.Addr) {
	serviceMu.Lock()
	defer serviceMu.Unlock()
	service[name] = addr
}

func LookupService(name string) net.Addr {
	serviceMu.Lock()
	defer serviceMu.Unlock()
	return service[name]
}
*/
