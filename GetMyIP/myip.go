package main

/*
Usage of ./getmyip --get_ip=(external|internal)
URL: http://myexternalip.com/#golang
URL: http://changsijay.com/2013/07/28/golang-get-ip-address/
*/

import (
	"flag"
	"io"
	"net"
	"net/http"
	"os"
)

var get_ip = flag.String("get_ip", "", "external|internal")

func main() {

	flag.Parse()

	if *get_ip == "external" {
		get_external()
	}

	if *get_ip == "internal" {
		get_internal()
	}

}

func get_external() {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	os.Exit(0)
}

func get_internal() {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}

	os.Exit(0)

}

