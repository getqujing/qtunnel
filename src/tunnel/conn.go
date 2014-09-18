package tunnel

import (
    "net"
    "time"
)

type Conn struct {
    conn net.Conn
    cipher *Cipher
}

func NewConn(conn net.Conn, cipher *Cipher) *Conn {
    return &Conn{conn, cipher}
}

func (c *Conn) Read(b []byte) (int, error) {
    c.conn.SetReadDeadline(time.Now().Add(30 * time.Minute))
    if c.cipher == nil {
        return c.conn.Read(b)
    }
    cipherData := make([]byte, len(b))
    n, err := c.conn.Read(cipherData)
    if n > 0 {
        c.cipher.decrypt(b[0:n], cipherData[0:n])
    }
    return n, err
}

func (c *Conn) Write(b []byte) (int, error) {
    if c.cipher == nil {
        return c.conn.Write(b)
    }
    cipherData := make([]byte, len(b))
    c.cipher.encrypt(cipherData, b)
    return c.conn.Write(cipherData)
}

func (c *Conn) Close() {
    c.conn.Close()
}

func (c *Conn) CloseRead() {
    if conn, ok := c.conn.(*net.TCPConn); ok {
        conn.CloseRead()
    }
}

func (c *Conn) CloseWrite() {
    if conn, ok := c.conn.(*net.TCPConn); ok {
        conn.CloseWrite()
    }
}
