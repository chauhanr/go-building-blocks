package socket

import "testing"

func TestIpAddressResolve(t *testing.T){
	tests := []struct{
		ipAddress string
		resolvedAddress string
	}{
		{"127.0.0.1", "127.0.0.1"},
		{"35.167.56.194", "35.167.56.194"},
		{"0:0:0:0:0:0:0:1","::1"},
	}

	for _, tc := range tests{
		result, err := ResolveIPAddress(tc.ipAddress)

		if err != nil{
			t.Errorf("expected ip address to be resolved to %s but got invalid address error ", result)
		}
		if result != tc.resolvedAddress {
			t.Errorf("Expected resolved address to be %s but got %s", tc.resolvedAddress, result)
		}
	}

}

func TestResolveMask(t *testing.T) {
	tests := []struct{
		ipAddress string
		resolvedAddress string
		mask string
		network string
	}{
		{"127.0.0.1", "127.0.0.1", "ff000000", "127.0.0.0"},
	}

	for _, tc := range tests{
		addr, mask, network, err := ResolveMask(tc.ipAddress)
		if err != nil{
			t.Errorf("Expected IP to be resolved properly but got error %s", err.Error())
		}
		if addr != tc.ipAddress || mask != tc.mask || network != tc.network{
			t.Errorf("Exepected (ip, mask, network) = (%s, %s, %s) but got (%s, %s, %s)", tc.ipAddress, tc.mask, tc.network, addr, mask, network)
		}
	}
}

func TestResolveHostName(t *testing.T) {
	name := "www.google.com"
	addr, err := ResolveHostName(name)
	if err != nil{
		t.Fatalf("Cannot resolve %s because of error %s", name, err.Error())
	}
	t.Logf("Address of %s resolved to %s", name, addr)
}

func TestHostLookup(t *testing.T) {
	name := "www.google.com"
	addrs, err := HostLookup(name)
	if err != nil{
		t.Fatalf("Cannot resolve %s because of error %s", name, err.Error())
	}
	for i, addr := range addrs{
		t.Logf("%d Host addr : %s", i, addr)
	}
}

func TestLookupServicePort(t *testing.T) {
	tests := []struct {
		service string
		networkType string
		port int
	}{
		{"telnet", "tcp", 23},
	}

	for _, tc := range tests{
		port, err := LookupServicePort(tc.networkType,tc.service)

		if err != nil{
			t.Errorf("Should give port %d but got error %s", tc.port, err.Error())
		}
		if port != tc.port{
			t.Errorf("Expected port %d but got %d", tc.port, port)
		}
	}
}

func TestTCPAddressResolve(t *testing.T) {
	tcpAddr, err := TCPAddressResolve("tcp", "www.google.com:80" +
		"")

	if err != nil{
		t.Fatalf("Error occured when it should work %s",err)
	}
	t.Logf("The TCP address resovled was %s", tcpAddr)
}
