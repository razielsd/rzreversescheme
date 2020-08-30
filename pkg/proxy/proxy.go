package proxy

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"rzreversescheme/pkg/core"
	"rzreversescheme/pkg/processor"
	"time"
)


type SchemeProxyServer struct {
    Server *http.Server
    RemoteProxyUrl string
}

var server *http.Server
var remoteProxyUrl string

func NewProxyServer(port int) SchemeProxyServer {
    proxy := SchemeProxyServer{}
	server = &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodConnect {
				proxy.HandleTunneling(w, r)
			} else {
				proxy.HandleHTTP(w, r)
			}
		}),
		// Disable HTTP/2.
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
	proxy.Server = server

	return proxy
}

func (proxy *SchemeProxyServer) SetRemoteProxyUrl(url string) {
	proxy.RemoteProxyUrl = url
}

func (proxy *SchemeProxyServer) Run(proto string, pemPath string, keyPath string) {
	if proto == "https" {
        log.Fatal(proxy.Server.ListenAndServeTLS(pemPath, keyPath))
	} else {
        log.Fatal(proxy.Server.ListenAndServe())
	}
}

func (proxy *SchemeProxyServer) HandleTunneling(w http.ResponseWriter, r *http.Request) {
	dest_conn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	client_conn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
	go proxy.Transfer(dest_conn, client_conn)
	go proxy.Transfer(client_conn, dest_conn)
}

func (proxy SchemeProxyServer) Transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}

func (proxy *SchemeProxyServer) HandleHTTP(w http.ResponseWriter, req *http.Request) {
	transport := http.Transport{}
	if (proxy.RemoteProxyUrl != "") {
		url_i := url.URL{}
		urlProxy, _ := url_i.Parse(proxy.RemoteProxyUrl)
		transport.Proxy = http.ProxyURL(urlProxy)// set proxy
	}
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //set ssl
	resp, err := transport.RoundTrip(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	proxy.processRequest(req, resp)

	proxy.copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func (proxy SchemeProxyServer) copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func (proxy *SchemeProxyServer) processRequest(req *http.Request, resp *http.Response) {
	clientReq := core.ClientRequest{}
	var bodyBytes []byte
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	clientReq.Response.Body = string(bodyBytes)
	clientReq.Response.Header = resp.Header
	clientReq.Response.StatusCode = resp.StatusCode
	clientReq.Request = *req
	processor.GetChannel() <- clientReq
}
