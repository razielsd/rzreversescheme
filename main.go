package main

import (
	"flag"
    "fmt"
    "log"
	"rzreversescheme/pkg/cmdserver"
	"rzreversescheme/pkg/processor"
	"rzreversescheme/pkg/proxy"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	var pemPath string
	var port int
	var cmdPort int
	var keyPath string
	var proto string
	var proxyUrl string //127.0.0.1:8080

	flag.IntVar(&cmdPort, "cmd-port", 9500, "port to command")
	flag.IntVar(&port, "port", 9501, "port proxy")
	flag.StringVar(&pemPath, "pem", "server.pem", "path to pem file (required for https)")
	flag.StringVar(&keyPath, "key", "server.key", "path to key file(required for https)")
	flag.StringVar(&proto, "proto", "http", "Proxy protocol (http or https)")
	flag.StringVar(&proxyUrl, "proxy", "", "Remote proxy host, example: 127.0.0.1:8080")
	flag.Parse()

	if proto != "http" && proto != "https" {
		log.Fatal("Protocol must be either http or https")
	}

	server := proxy.NewProxyServer(port)
	if (proxyUrl != "") {
		server.SetRemoteProxyUrl(proxyUrl)
	}
	processor.Init()

	fmt.Println("---------------------------------------------")
	fmt.Printf("Proxy port: %d\n", port)
    fmt.Printf("Command port: %d\n", cmdPort)
    fmt.Println("---------------------------------------------")

	go server.Run(proto, pemPath, keyPath)

	releaseMode := gin.ReleaseMode
	releaseMode = gin.DebugMode // @todo get from config
	gin.SetMode(releaseMode)
	gin.DisableConsoleColor()
	r := cmdserver.LoadRouter()
	r.Run(":" + strconv.Itoa(cmdPort))

}
