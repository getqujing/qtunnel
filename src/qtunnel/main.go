package main

import (
    "os"
    "log"
    "flag"
    "tunnel"
)

func main() {
    log.SetOutput(os.Stdout)
    var faddr, baddr, cryptoMethod, secret string
    var clientMode bool
    flag.StringVar(&faddr, "listen", ":9001", "host:port qtunnel listen on")
    flag.StringVar(&baddr, "backend", "127.0.0.1:6400", "host:port of the backend")
    flag.StringVar(&cryptoMethod, "crypto", "rc4", "encryption method")
    flag.StringVar(&secret, "secret", "secret", "password used to encrypt the data")
    flag.BoolVar(&clientMode, "clientmode", false, "if running at client mode")
    flag.Parse()

    t := tunnel.NewTunnel(faddr, baddr, clientMode, cryptoMethod, secret)
    t.Start()
}
