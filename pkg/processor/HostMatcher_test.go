package processor

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)


func Test_HostMatcher_Append(t *testing.T) {
	var rule = HostMatchRule{
		Host:       "example.com",
		HostFilter: "mytest.local",
	}
	var matcher = NewHostMatcher()
	assert.Equal(t, 0, len(matcher.rules), "HostMatcher.rules must be empty after create")
	matcher.Append(rule)
	assert.Equal(t, 1, len(matcher.rules), "HostMatcher.rules must be have single rule after add")
}


func Test_HostMatcher_Count(t *testing.T) {
	var rule = HostMatchRule{
		Host:       "example.com",
		HostFilter: "mytest.local",
	}
	var matcher = NewHostMatcher()
	assert.Equal(t, 0, matcher.Count(), "HostMatcher.Count() must be empty after create")
	matcher.Append(rule)
	assert.Equal(t, 1, matcher.Count(), "HostMatcher.Count() must be have single rule after add")
}


func Test_HostMatcher_GetHostName_NoPort(t *testing.T) {
	var matcher = NewHostMatcher()
	var request, _ = http.NewRequest(http.MethodGet, "http://localtest/add/item", nil)
	assert.Equal(t, "localtest", matcher.getHostName(*request),)
}


func Test_HostMatcher_GetHostName_WithPort(t *testing.T) {
	var matcher = NewHostMatcher()
	var request, _ = http.NewRequest(http.MethodGet, "http://localtest:7080/add/item", nil)
	assert.Equal(t, "localtest", matcher.getHostName(*request),)
}


func Test_MatchRule_MatchHost_MatchFull(t *testing.T) {
	rule := HostMatchRule{
		Host:       "example.com",
		HostFilter: "mytest.local",
	}
	assert.Equal(t, RULE_FILTER_HOST, rule.MatchHost("mytest.local"))
}


func Test_MatchRule_MatchHost_MatchPartial(t *testing.T) {
	rule := HostMatchRule{
		Host:       "example.com",
		HostFilter: "mytest.l",
	}
	assert.Equal(t, RULE_FILTER_HOST, rule.MatchHost("mytest.local"))
}


func Test_MatchRule_MatchHost_Match_None(t *testing.T) {
	rule := HostMatchRule{
		Host:       "example.com",
		HostFilter: "mytest.local",
	}
	assert.Equal(t, RULE_FILTER_NONE, rule.MatchHost("mytest.33"))
}


func Test_MatchRule_MatchPort_Match_Star(t *testing.T) {
	rule := HostMatchRule{
		Host:       "example.com",
		PortFilter: "*",
	}
	assert.Equal(t, RULE_FILTER_PORT, rule.MatchPort("8091"))
}


func Test_MatchRule_MatchPort_Match_Success(t *testing.T) {
	rule := HostMatchRule{
		Host:       "example.com",
		PortFilter: "8080",
	}
	assert.Equal(t, RULE_FILTER_PORT, rule.MatchPort("8080"))
}


func Test_MatchRule_MatchPort_Match_None(t *testing.T) {
	rule := HostMatchRule{
		Host:       "example.com",
		PortFilter: "8080",
	}
	assert.Equal(t, RULE_FILTER_NONE, rule.MatchPort("80"))
}

