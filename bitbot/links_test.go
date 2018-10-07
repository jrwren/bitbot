package bitbot

import (
	"bytes"
	"strings"
	"testing"
)

func TestGetHTMLTitle(t *testing.T) {
	buf := &bytes.Buffer{}
	buf.WriteString("<html><head><title>")
	for i := 0; i < 1000000000; i++ {
		buf.WriteString("aaa")
	}
	buf.WriteString("</title></head></html>")
	title, err := GetHtmlTitle(buf)
	if err != true {
		t.Log("could not parse huge title")
		t.Fail()
	}
	if len(title) > 120 || !strings.HasPrefix(title, "aaaa") {
		t.Log("unexpected title")
		t.Fail()
	}
}
