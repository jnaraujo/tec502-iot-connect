package network

import (
	"broker/sensor"
	"net"
	"sync"
)

func GetAllSensorsIps() []string {
	ifaces, err := net.Interfaces()

	if err != nil {
		return nil
	}

	var wg sync.WaitGroup
	results := make(chan string)

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()

		if err != nil {
			continue
		}

		for _, addr := range addrs {
			ip, _, err := net.ParseCIDR(addr.String())

			if err != nil {
				continue
			}

			if ip.IsLoopback() || ip.To4() == nil {
				continue
			}

			wg.Add(1)
			go func(ip net.IP) {
				defer wg.Done()

				println("Checking", ip.String())

				address := ip.String() + ":3333"

				if checkTcpPort(address) {
					results <- address
				}
			}(ip)
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var addrs []string

	for result := range results {
		addrs = append(addrs, result)
	}

	return addrs
}

func checkTcpPort(address string) bool {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return false
	}

	defer conn.Close()

	return sensor.ValidateHandshake(conn)
}
