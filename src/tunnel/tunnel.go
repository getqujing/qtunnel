package tunnel

import (
    "io"
    "net"
    "log"
)

type Tunnel struct {
    faddr, baddr *net.TCPAddr
    clientMode bool
    cryptoMethod, secret string
}

func NewTunnel(faddr, baddr string, clientMode bool, cryptoMethod, secret string) *Tunnel {
    a1, err := net.ResolveTCPAddr("tcp", faddr)
    if err != nil {
        log.Fatalln("resolve frontend error:", err)
    }
    a2, err := net.ResolveTCPAddr("tcp", baddr)
    if err != nil {
        log.Fatalln("resolve backend error:", err)
    }
    return &Tunnel{a1, a2, clientMode, cryptoMethod, secret}
}

func (t *Tunnel) pipe(dst, src *Conn) {
    defer dst.Close()
    _, err := io.Copy(dst, src)
    if err != nil {
        log.Print(err)
    }
}

func (t *Tunnel) transport(conn net.Conn) {
    conn2, err := net.DialTCP("tcp", nil, t.baddr)
    if err != nil {
        log.Print(err)
        return
    }
    cipher := NewCipher(t.cryptoMethod, t.secret)
    var bconn, fconn *Conn
    if t.clientMode {
        fconn = NewConn(conn, nil)
        bconn = NewConn(conn2, cipher)
    } else {
        fconn = NewConn(conn, cipher)
        bconn = NewConn(conn2, nil)
    }
    go t.pipe(bconn, fconn)
    go t.pipe(fconn, bconn)
}

func (t *Tunnel) Start() {
    ln, err := net.ListenTCP("tcp", t.faddr)
    if err != nil {
        log.Fatal(err)
    }
    defer ln.Close()

    for {
        conn, err := ln.AcceptTCP()
        if err != nil {
            log.Println("accept:", err)
            continue
        }
        go t.transport(conn)
    }
}
