package stanza

import (
    "testing"
)

type XMLData struct {
    Data []byte 
    ExpectedBase BaseXML
}

func TestUnmarshalXML(t *testing.T) {
    var data []byte = []byte(`<message><to>Testing</to></message>`)
    var name xml.Name = xml.Name{"", "message"}
    var attrs []xml.Attr = []xml.Attr
    var tks []xml.Token = []xml.Token
    var newData XMLData = XMLData{Data: data}
}
