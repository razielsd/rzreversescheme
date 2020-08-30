package proxy

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_NewProxyServer(t *testing.T) {
    server := NewProxyServer(300)
    assert.IsType(t, SchemeProxyServer{}, server, "Expected *SchemeProxyServer for instance proxy")
}

func Test_SetRemoteProxyUrl(t *testing.T) {
    server := NewProxyServer(300)
    var proxyUrl string = "myproxy"
    server.SetRemoteProxyUrl(proxyUrl)
    assert.Equal(t, proxyUrl, server.RemoteProxyUrl)
}

