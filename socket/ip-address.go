package socket

import (
	"net"
	"fmt"
	"errors"
)

func ResolveIPAddress(address string) (string,error){
	addr := net.ParseIP(address)
	if addr == nil{
		fmt.Println("Invalid address")
		return "", errors.New("invalid address "+address)
	}else {
		fmt.Println("The address is : ", addr.String())
	}
	return addr.String(), nil
}


func ResolveMask(addr string) (string, string, string, error) {
	add := net.ParseIP(addr)
	if add == nil{
		return "", "", "", errors.New("Unable to resolve address : "+addr)
	}
	mask := add.DefaultMask()
	network := add.Mask(mask)
	ones, bits := mask.Size()
	fmt.Printf("Leading ones %d leading bits %d \n", ones,bits)
	return add.String(), mask.String(), network.String(), nil
}

func ResolveHostName(addr string) (string, error){
	add, err := net.ResolveIPAddr("ip", addr)
	if err == nil{
		return add.String(),err
	}
	return "", err
}

func HostLookup(addr string) ([]string, error){
	addrs, err := net.LookupHost(addr)
	if err == nil {
		return addrs, err
	}
	return []string{}, err
}

func LookupServicePort(netType string, service string) (int, error){
	port, err := net.LookupPort(netType, service)
	if err == nil {
		return port, err
	}
	return 0, err
}

func TCPAddressResolve(nw string, addr string) (string, error){
	 tcpAddr, err := net.ResolveTCPAddr(nw, addr)
	 return tcpAddr.String(), err
}