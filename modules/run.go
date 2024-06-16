package modules

import (
	"flag"
	"log"
)

var (
	check   bool
	exploit bool
	target  string
)

func Run() {
	flag.BoolVar(&check, "c", false, "check is vuln(default)")
	flag.BoolVar(&exploit, "e", false, "reverse a shell")
	flag.StringVar(&target, "u", "", "target url")
	flag.Parse()
	if check && exploit {
		log.Fatalln("You Can Only Specific Check Mode Or Exploit Mode")
	}
	if _, _, invaild := UrlChecker(target); !invaild {
		log.Fatalln("URL Invalid.")
	}
	hik := Hik{
		Target:  target,
		Exploit: exploit,
	}
	hik.Run()
}
