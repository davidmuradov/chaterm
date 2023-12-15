package main

import (
    "fmt"
    "log"
    "bufio"
    "io"
    "net"
    "encoding/xml"
    "errors"
)

const (
    addr string = "127.0.0.1:5555"
)

type BaseXMLMessage struct {
    XMLName xml.Name 
    Attrs []xml.Attr
    Data []byte 
    Nodes []*BaseXMLMessage
}

func (b *BaseXMLMessage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    // Copy the name and the attributes into the BaseXML struct
    b.XMLName = start.Name
    attrs := start.Copy().Attr
    b.Attrs = attrs
    for {
        // Get every token
        tk, err := d.Token()
        log.Println(tk)
        if err != nil {
            if err == io.EOF {
                break
            } else {
                return err
            }
        }
        // Check for the type of the token
        switch tk.(type) {
            case xml.CharData:
                // Copy the data
                data := tk.(xml.CharData).Copy()
                b.Data = data

            case xml.StartElement:
               // Create a new node
               node := BaseXMLMessage{}
               // Decode the node
               strt := tk.(xml.StartElement)
               err = d.DecodeElement(&node, &strt)
               if err != nil {
                   log.Panicln(err)
               }
               // Add it to the nodes
               b.Nodes = append(b.Nodes, &node)
            case xml.EndElement:
                // Break when the end has been reached
                break
        }
    }
    return nil
}

type Animal struct {
    XMLName xml.Name `xml:"animal"`
    Name string `xml:"name"`
    Age int `xml:"age"`
}

func manageConnection(c net.Conn) {
    defer c.Close()
    fmt.Println("Client connected")
    limitReader := io.LimitReader(c, 1024*10)
    connReader := bufio.NewReaderSize(limitReader, 23)
    xmlDecoder := xml.NewDecoder(connReader)
    for {
        rawxml := BaseXMLMessage{}
        err := xmlDecoder.Decode(&rawxml)
        if err != nil {
            var syn *xml.SyntaxError
            if err == io.EOF {
                fmt.Println("EOF")
                break
            } else if errors.As(err, &syn) {
                if limitReader.(*io.LimitedReader).N <= 0 {
                    fmt.Println("Max message size reached.")
                } else {
                    fmt.Println("syntax error:", err.Error())
                }
                break
            } else {
                log.Panicln(err)
            }
        } else {
            fmt.Println("xml:", rawxml)
        }
    }
    fmt.Println("Client disconnected")
}

func main() {
    server, err := net.Listen("tcp", addr)
    if err != nil {
        log.Panicln(err)
    }
    defer server.Close()
    connClient, err := server.Accept()
    if err != nil {
        log.Panicln(err)
    }
    go manageConnection(connClient)
    for {}
}
