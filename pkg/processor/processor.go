package processor

import (
	"fmt"
	"runtime"
	"rzreversescheme/pkg/core"
	"rzreversescheme/pkg/format"
	"rzreversescheme/pkg/swagger"
	"strings"
	"sync"
)

type IApiProcessor interface {
	GetScheme() string
	Process(clientReq core.ClientRequest)
}


var clientChan chan core.ClientRequest
var procList map[string]IApiProcessor
var procMutex = &sync.Mutex{}

func Init()  {
	clientChan = make(chan core.ClientRequest, 100)
	procList = make(map[string]IApiProcessor)
	go worker(clientChan)
}


func GetChannel() chan core.ClientRequest {
	return clientChan
}


func GetProcessorByHost(host string) IApiProcessor  {
	procMutex.Lock()
	defer procMutex.Unlock()
	proc, ok := procList[host];
	if !ok {
		return nil
	}
	return proc
}


func worker(reqChan chan core.ClientRequest) {
	for {
		processClientRequest(<- reqChan)
		runtime.Gosched()
	}
}

func processClientRequest(clientReq core.ClientRequest) {
	//fmt.Println(clientReq.Request.RequestURI)
	//fmt.Printf("Host: %s\n", clientReq.Request.Host)
	//fmt.Println(clientReq.Response.Header.Get("Content-Type"))
	proc := getProcessor(clientReq)
	if proc == nil {
		return ;
	}
	proc.Process(clientReq)
	fmt.Println(clientReq.Response.Body)
}


func getProcessor(clientReq core.ClientRequest) IApiProcessor  {
	var host = extractHost(clientReq.Request.Host)
	fmt.Printf("Get for host: %s\n", host)
	var contentType = strings.ToLower(clientReq.Response.Header.Get("content-type"))
	var skipTypeList = [3]string{"image", "video", "audio"}
	for _, typeName := range skipTypeList {
		if strings.Contains(contentType, typeName) {
			return nil;
		}
	}
	if format.IsJson(clientReq.Response.Body) == true {
		procMutex.Lock()
		defer procMutex.Unlock()

		proc, ok := procList[host];
		if !ok {
			proc = swagger.NewSwaggerProcessor()
			procList[host] = proc
			fmt.Printf("Add host: %s\n", host)
		}
		return proc
	}
	return nil
}

/**
  Drop port number for test,
 */
func extractHost(host string) string  {
	var index = strings.Index(host, ":")
	if index > 0 {
		host = host[0:index]
	}
	return host
}