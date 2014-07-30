package tunnel

import (
    "testing"
    "net"
    "bytes"
)

func TestWrite(t *testing.T) {
    done := make(chan bool)
    secret := []byte("secret")
    data := []byte{0, 1, 2, 3}
    ln, err := net.Listen("tcp", "127.1:9444")
    if err != nil {
        t.Error(err)
    }
    defer ln.Close()
    go func() {
        conn, err := ln.Accept()
        if err != nil {
            t.Error(err)
        }
        defer conn.Close()
        cipher := NewCipher("rc4", secret)
        buf := make([]byte, len(data))
        conn.Read(buf)
        cipherData := make([]byte, len(data))
        cipher.encrypt(cipherData, data)
        if (bytes.Compare(cipherData, buf) != 0) {
            t.Fail()
        }
        close(done)
    }()
    conn, err := net.Dial("tcp", "127.1:9444")
    if err != nil {
        t.Error(err)
    }
    conn2 := NewConn(conn, NewCipher("rc4", secret))
    defer conn2.Close()
    conn2.Write(data)
    <-done
}

func TestRead(t *testing.T) {
    done := make(chan bool)
    secret := []byte("secret")
    data := []byte{0, 1, 2, 3}
    ln, err := net.Listen("tcp", "127.1:9444")
    if err != nil {
        t.Error(err)
    }
    defer ln.Close()
    go func() {
        conn, err := ln.Accept()
        if err != nil {
            t.Error(err)
        }
        defer conn.Close()
        cipher := NewCipher("rc4", secret)
        cipherData := make([]byte, len(data))
        cipher.encrypt(cipherData, data)
        conn.Write(cipherData)
        close(done)
    }()
    conn, err := net.Dial("tcp", "127.1:9444")
    if err != nil {
        t.Error(err)
    }
    conn2 := NewConn(conn, NewCipher("rc4", secret))
    defer conn2.Close()
    buf := make([]byte, len(data))
    conn2.Read(buf)
    if (bytes.Compare(data, buf) != 0) {
        t.Fail()
    }
    <-done
}
