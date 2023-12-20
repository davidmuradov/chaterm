package main
// test

import (
    "time"
    //"fmt"
    "log"
    "net"
    "encoding/xml"
)

const (
    addr string = "127.0.0.1:5555"
)

type Animal struct {
    XMLName xml.Name `xml:"animal"`
    Name string `xml:"name"`
    Age int `xml:"age"`
    MaxAge int `xml:"maxage"`
}

func main() {
    remoteAddr, err := net.ResolveTCPAddr("tcp", addr)
    if err != nil {
        log.Panicln(err)
    }
    conn, err := net.DialTCP("tcp", nil, remoteAddr)
    if err != nil {
        log.Panicln(err)
    }
    defer conn.Close()
    //xmlEncoder := xml.NewEncoder(conn)
    var animal1 Animal = Animal{Name: "Dog<name>", Age: 6, MaxAge: 20}
    mars, _ := xml.Marshal(animal1)
    log.Println(string(mars))
    _, err = conn.Write(mars[:len(mars) - 9])
    if err != nil {
        log.Panicln(err)
    }
    time.Sleep(time.Second * 1)
    //_, err = conn.Write(mars[len(mars) - 9:len(mars)])
    //err = xmlEncoder.Encode(animal1)
    //if err != nil {
    //    log.Panicln(err)
    //}
    for {
        _, err := conn.Write([]byte("abcedfef"))
        if err != nil {
            log.Panicln(err)
        }
    }
}
