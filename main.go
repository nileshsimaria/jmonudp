package main

import (
	"encoding/json"
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
	jsonData = flag.Bool("json", false, "Convert data into JSON")
)

func handleUDPConnection(conn *net.UDPConn) {

	buffer := make([]byte, 4096)

	n, _, err := conn.ReadFromUDP(buffer)

	ts := &tt.TelemetryStream{}
	if err := proto.Unmarshal(buffer[:n], ts); err != nil {
		log.Fatalln("Failed to parse: ", err)
	}

	if *jsonData {
		b, err := json.MarshalIndent(ts, "", "  ")
		if err != nil {
			log.Fatalln("JSON indent error: ", err)
		}
		fmt.Printf("%s\n", b)
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
					if *jsonData {
						b, err := json.MarshalIndent(qm, "", "  ")
						if err != nil {
							log.Fatalln("JSON indent error: ", err)
						}
						fmt.Printf("%s\n", b)
					}
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
