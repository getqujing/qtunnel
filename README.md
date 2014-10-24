#qTunnel

### qTunnel - the tunnel program wrapping communication on getqujing.net

`qtunnel` is a network proxy to work as an encryption wrapper between remote client and local or remote server. It can work as a stunnel/stud replacement.

##### Why another wrapper

stud is working perfectly in SSL/TLS based environments, but sometimes all you need is a encryption tunnel without the certification settings and save some time from SSL handshake, so we build `qtunnel` to fit the special network requirement in China. It serves 10m+ connections per day on getqujing.net for months.

### Requirements

qtunnel is writen in [golang 1.3.1](http://golang.org/dl/), after build it can run on every OS now.

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
 		
`qtunnel` support two crypto method for now: `rc4` and `aes256cfb`, both side of the tunnel should use same `crypto` and same `secret`.

