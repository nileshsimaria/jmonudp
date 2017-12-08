package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	qmon "github.com/nileshsimaria/jmonudp/protos/qmon"
	tt "github.com/nileshsimaria/jmonudp/protos/telemetry_top"
	flag "github.com/spf13/pflag"
	"log"
	"net"
)

var (
	hostname = "0.0.0.0"
	port     = flag.Int64("port", 0, "UDP port number to listen on")
)

func handleUDPConnection(conn *net.UDPConn) {

	buffer := make([]byte, 4096)

	n, addr, err := conn.ReadFromUDP(buffer)
	fmt.Println("UDP client : ", addr)

	ts := &tt.TelemetryStream{}
	if err := proto.Unmarshal(buffer[:n], ts); err != nil {
		log.Fatalln("Failed to parse: ", err)
	}
	if proto.HasExtension(ts.Enterprise, tt.E_JuniperNetworks) {
		jns_i, err := proto.GetExtension(ts.Enterprise, tt.E_JuniperNetworks)
		if err != nil {
			log.Fatalln("Cant get extension: ", err)
		}
		switch jns := jns_i.(type) {
		case *tt.JuniperNetworksSensors:
			if proto.HasExtension(jns, qmon.E_JnprQmonExt) {
				qm_i, err := proto.GetExtension(jns, qmon.E_JnprQmonExt)
				if err != nil {
					log.Fatalln("Cant get extension: ", err)
				}
				switch qm := qm_i.(type) {
				case *qmon.QueueMonitor:
					fmt.Printf("%v\n", qm)
				}
			}
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	address := fmt.Sprintf("%s:%d", hostname, *port)
	udpAddr, err := net.ResolveUDPAddr("udp4", address)

	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	for {
		handleUDPConnection(ln)
	}
}
