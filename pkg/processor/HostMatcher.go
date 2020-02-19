package processor

import (
	"net/http"
	"strings"
	"sync"
)

const RULE_FILTER_NONE = 0
const RULE_FILTER_HOST = 1
const RULE_FILTER_PORT = 2
const RULE_FILTER_PATH = 4

type HostMatcher struct {
	cache map[string]string
	rules []HostMatchRule
	lock  sync.RWMutex
}


type HostMatchRule struct {
	Host       string `json:"host"`
	HostFilter string `json:"hostfilter"`
	PortFilter string `json:"portfilter"`
	PathFilter string `json:"pathfilter"`
	checksum   int
}


func NewHostMatcher() HostMatcher{
	var matcher = HostMatcher{}
	matcher.lock = sync.RWMutex{}
	return matcher
}


func (matcher *HostMatcher) Append(rule HostMatchRule) {
	rule.Init()
	matcher.lock.Lock()
	defer matcher.lock.Unlock()
	matcher.rules = append(matcher.rules, rule)
}


func (matcher *HostMatcher) Count() int {
	return len(matcher.rules)
}


func (matcher HostMatcher) GetHost(request http.Request) string {
	var host = matcher.getHostName(request)
	matcher.lock.RLock()
	var ruleList = matcher.rules
	matcher.lock.Unlock()
	for _, rule := range ruleList {
		var checksum = rule.MatchHost(host) +
			rule.MatchPort(request.URL.Port()) +
			rule.MatchPath(request.URL.Path)
		if checksum == rule.checksum {
			return rule.Host
		}
	}
	return host
}


func (matcher HostMatcher) getHostName(r http.Request) string {
	if r.URL.IsAbs() {
		host := r.Host
		if i := strings.Index(host, ":"); i != -1 {
			host = host[:i]
		}
		return host
	}
	return r.URL.Host + ":"
}


func (rule *HostMatchRule) Init() {
	rule.checksum = 0;
	if rule.HostFilter != "" {
		rule.checksum += RULE_FILTER_HOST
	}
	if rule.PortFilter != "" {
		rule.checksum += RULE_FILTER_PORT
	}
	if rule.PathFilter != "" {
		rule.checksum += RULE_FILTER_PATH
	}
}


func (rule HostMatchRule) HasActiveFilter() bool {
	return rule.checksum > 0
}


func (rule HostMatchRule) MatchHost(host string) int {
	if (rule.HostFilter != "") && (strings.HasPrefix(host + ":", rule.HostFilter)) {
		return RULE_FILTER_HOST
	}
	return RULE_FILTER_NONE
}


func (rule HostMatchRule) MatchPort(port string) int {
	if rule.PortFilter == "" {
		return RULE_FILTER_NONE
	}
	if (rule.PortFilter == "*") || (rule.PortFilter == port) {
		return RULE_FILTER_PORT
	}
	return RULE_FILTER_NONE
}


func (rule HostMatchRule) MatchPath(path string) int {
	if (rule.HostFilter != "") && (strings.HasPrefix(path, rule.PathFilter)) {
		return RULE_FILTER_PATH
	}
	return RULE_FILTER_NONE
}
