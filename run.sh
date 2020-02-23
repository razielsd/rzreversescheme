#!/usr/bin/env bash
go clean -testcache && go run main.go --proto=http

# https://community-open-weather-map.p.rapidapi.com/weather?callback=test&id=2172797&units=%22metric%22 or %22imperial%22&mode=xml%2Chtml&q=London%2Cuk