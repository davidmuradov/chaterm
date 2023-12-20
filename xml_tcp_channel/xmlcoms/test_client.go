package main

import (
    "log"
    "github.com/Vacheprime/xmlcoms/stanza"
    "github.com/Vacheprime/xmlcoms/testxml"
)

const (
    addr string = "127.0.0.1:5555"
)

func main() {
    xmlComm := xmlcoms.NewXMLCommunicator()
    // Connect to server
    err := xmlComm.Connect("", addr, "tcp")
    if err != nil {
        log.Panicln(err)
    }
    // Receive a message
    stz, err := xmlComm.ReceiveStanza()
    if err != nil {
        log.Panicln(err)
    }
    log.Println(stz.(stanza.Message).Body)
}
