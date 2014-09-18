package tunnel

import (
    "testing"
    "net"
    "bytes"
    "time"
)

func TestTunnel(t *testing.T) {
    done := make(chan bool)
    data := []byte{1, 2, 3, 4}
    go backendServer(data, done, t)
    b := NewTunnel("127.1:9446", "127.1:9445", false, "rc4", "secret", 4)
    f := NewTunnel("127.1:9447", "127.1:9446", true, "rc4", "secret", 4)
    go b.Start()
    go f.Start()
    // sleep to wait all servers start
    time.Sleep(100 * time.Millisecond)
    conn, err := net.Dial("tcp", "127.1:9447")
    if err != nil {
        t.Error(err)
    }
    defer conn.Close()
    conn.Write(data)
    // wait for transmission complete
    time.Sleep(100 * time.Millisecond)
    close(done)
}

func backendServer(data []byte, done chan bool, t *testing.T) {
    ln, err := net.Listen("tcp", "127.1:9445")
    if err != nil {
        t.Error(err)
    }
    defer ln.Close()
    conn, err := ln.Accept()
    if err != nil {
        t.Error(err)
    }
    defer conn.Close()
    buf := make([]byte, len(data))
    conn.Read(buf)
    if (bytes.Compare(buf, data) != 0) {
        t.Fail()
    }
    <-done
}
