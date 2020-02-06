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

var server *http.Server
var remoteProxyUrl string

func CreateServer(port int) *http.Server {
	server = &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodConnect {
				handleTunneling(w, r)
			} else {
				handleHTTP(w, r)
			}
		}),
		// Disable HTTP/2.
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
	return server
}

func SetRemoteProxyUrl(url string) {
	remoteProxyUrl = url
}

func Run(proto string, pemPath string, keyPath string) {
	if proto == "http" {
		log.Fatal(server.ListenAndServe())
	} else {
		log.Fatal(server.ListenAndServeTLS(pemPath, keyPath))
	}

}

func handleTunneling(w http.ResponseWriter, r *http.Request) {
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
	go transfer(dest_conn, client_conn)
	go transfer(client_conn, dest_conn)
}

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}

func handleHTTP(w http.ResponseWriter, req *http.Request) {
	transport := http.Transport{}
	if (remoteProxyUrl != "") {
		url_i := url.URL{}
		urlProxy, _ := url_i.Parse(remoteProxyUrl)
		transport.Proxy = http.ProxyURL(urlProxy)// set proxy
	}
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //set ssl
	resp, err := transport.RoundTrip(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	processRequest(req, resp)

	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func processRequest(req *http.Request, resp *http.Response) {
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