#qTunnel

### qTunnel - a simpler and (possibily) faster tunnel program

`qtunnel` is a network tunneling software working as an encryption wrapper between clients and servers (remote/local). It can work as a Stunnel/stud replacement.

`qtunnel` has been serving over 10 millions connections on [Qu Jing](http://getqujing.com) each day for the past few months.

##### Why Another Wrapper

Stunnel/stud is great in SSL/TLS based environments, but what we want is a lighter and faster solution that only does one job: transfer encrypted data between servers and clients. We don't need to deal with certification settings and we want the transfer is as fast as possible. So we made qTunnel. Basically, it's a Stunnel/stud without certification settings and SSL handshakes, and it's written in Go.

### Requirements

qtunnel is writen in [golang 1.3.1](http://golang.org/dl/), after building it can run on almost every OS.

### Build

To build `qtunnel`

`$ make`

To test `qtunnel`

`$ make test`

### Usage

	$ ./bin/qtunnel -h
	Usage of ./bin/qtunnel:
		-backend="127.0.0.1:6400": host:port of the backend
		-clientmode=false: if running at client mode
		-crypto="rc4": encryption method
		-listen=":9001": host:port qtunnel listen on
		-logto="stdout": stdout or syslog
		-secret="secret": password used to encrypt the data
 		
`qtunnel` supports two encryption methods: `rc4` and `aes256cfb`. Both servers and clients should use the same `crypto` and same `secret`.

### Credits

Special thanks to [Paul](http://paulrosenzweig.com) for reviewing the code.

### Contributing

We encourage you to contribute to `qtunnel`! Please feel free to [submit a bug report](https://github.com/getqujing/qtunnel/issues), [fork the repo](https://github.com/getqujing/qtunnel/fork) or [create a pull request](https://github.com/getqujing/qtunnel/pulls).

### License

`qtunnel` is released under the [Apache License 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).

