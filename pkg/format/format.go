package format

import (
	"encoding/json"
	"encoding/xml"
)

func IsJson(s string) bool {
	return json.Valid([]byte(s))
}

func IsXml(text string) bool {
	_, err := xml.Marshal(text)
	return err == nil
}