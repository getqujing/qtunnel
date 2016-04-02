package tunnel

import (
    "testing"
    "net"
    "bytes"
    "math/rand"
)

func makeTestBuf() ([]byte, []byte) {
    buf1 := make([]byte, rand.Intn(2048000) + 20480000)
    for i := 0; i <= 100; i++ {
        buf1[rand.Intn(len(buf1))] = 1
    }
    buf2 := make([]byte, len(buf1))
    copy(buf1, buf2)
    return buf1, buf2
}

func TestWrite(t *testing.T) {
    done := make(chan bool)
    secret := []byte("secret")
    data, data2 := makeTestBuf()
    pool := NewRecycler(4)
    ln, err := net.Listen("tcp", "127.0.0.1:9444")
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
        total := 0
        n, err := conn.Read(buf)
        for {
            total += n
            if total < len(data) {
                n, err = conn.Read(buf[total:len(buf)])
                if err != nil {
                    break
                }
            } else {
                break
            }
        }
        cipherData := make([]byte, len(data))
        cipher.encrypt(cipherData, data)
        if (bytes.Compare(cipherData, buf) != 0) {
            t.Fail()
        }
        close(done)
    }()
    conn, err := net.Dial("tcp", "127.0.0.1:9444")
    if err != nil {
        t.Error(err)
    }
    conn2 := NewConn(conn, NewCipher("rc4", secret), pool)
    defer conn2.Close()
    total := 0
    n, err := conn2.Write(data2)
    for {
        total += n
        if total < len(data2) {
            n, err = conn2.Write(data[total:len(data2)])
            if err != nil {
                break
            }
        } else {
            break
        }
    }
    <-done
}

func TestRead(t *testing.T) {
    done := make(chan bool)
    secret := []byte("secret")
    data, data2 := makeTestBuf()
    pool := NewRecycler(4)
    ln, err := net.Listen("tcp", "127.0.0.1:9444")
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
        total := 0
        n, err := conn.Write(cipherData)
        for {
            total += n
            if total < len(cipherData) {
                n, err = conn.Write(cipherData[total:len(cipherData)])
                if err != nil {
                    break
                }
            } else {
                break
            }
        }
        close(done)
    }()
    conn, err := net.Dial("tcp", "127.0.0.1:9444")
    if err != nil {
        t.Error(err)
    }
    conn2 := NewConn(conn, NewCipher("rc4", secret), pool)
    defer conn2.Close()
    buf := make([]byte, len(data2))
    total := 0
    n, err := conn2.Read(buf)
    for {
        total += n
        if total < len(buf) {
            n, err = conn2.Read(buf[total:len(buf)])
            if err != nil {
                break
            }
        } else {
            break
        }
    }
    if (bytes.Compare(data2, buf) != 0) {
        t.Fail()
    }
    <-done
}
