package xmlcoms

import (
    "bufio"
    "io"
    "net"
    "encoding/xml"
    "errors"
    
    "github.com/Vacheprime/xmlcoms/stanzas"
)

type XMLCommunicator struct {
    conn *net.TCPConn
    d *xml.Decoder
    l *io.LimitedReader
}

func NewXMLCommunicator() *XMLCommunicator {
    return &XMLCommunicator{conn: nil, d: nil}
}

func (c *XMLCommunicator) Connect(laddr, raddr, proto string) error {
    // Connect if not already connected
    if c.conn == nil {
        return errors.New("The communicator already possesses a connection!")
    }
    // Generate both local addr and remote addr
    local, err := net.ResolveTCPAddr(proto, laddr)
    if err != nil {
        return err
    }
    remote, err := net.ResolveTCPAddr(proto, raddr)
    if err != nil {
        return err
    }
    // Attempt to connect
    c.conn, err = net.DialTCP(proto, local, remote)
    if err != nil {
        return err
    }
    // Initialize the xml decoder
    c.l = io.LimitReader(c.conn, 1024 * 10).(*io.LimitedReader)
    c.d = xml.NewDecoder(bufio.NewReaderSize(c.l, 1024)) 

    return nil
} 

func (c *XMLCommunicator) ReceiveStanza() (stanzas.Stanza, error) {
    return nil, nil
}
