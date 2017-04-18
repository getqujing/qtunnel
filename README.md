# qTunnel

### qTunnel - a simpler and (possibily) faster tunnel program

`qtunnel` is a network tunneling software working as an encryption wrapper between clients and servers (remote/local). It can work as a Stunnel/stud replacement.

`qtunnel` has been serving over 10 millions connections on [Qu Jing](http://getqujing.com) each day for the past few months.

##### Why Another Wrapper

[Stunnel](https://www.stunnel.org/index.html)/[stud](https://github.com/bumptech/stud) is great in SSL/TLS based environments, but what we want is a lighter and faster solution that only does one job: transfer encrypted data between servers and clients. We don't need to deal with certification settings and we want the transfer is as fast as possible. So we made qTunnel. Basically, it's a Stunnel/stud without certification settings and SSL handshakes, and it's written in Go.

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

### Example

Let's say, you have a `redis` server on `host-a`, you want to connect to it from `host-b`, normally, just use:

	$ redis-cli -h host-a -p 6379

will do the job. The topology is:

	redis-cli (host-b) <------> (host-a) redis-server

If the host-b is in some insecure network environment, i.e. another data center or another region, the clear-text based redis porocol is not good enough, you can use `qtunnel` as a secure wrapper

On `host-b`:

	$ qtunnel -listen=127.1:6379 -backend=host-a:6378 -clientmode=true -secret=secret -crypto=rc4

On `host-a`:

	$ qtunnel -listen=:6378 -backend=127.1:6379 -secret=secret -crypto=rc4

Then connect on `host-b` as:

	$ redis-cli -h 127.1 -p 6379

This will establish a secure tunnel between your `redis-cli` and `redis` server, the topology is:

	redis-cli (host-b) <--> qtunnel (client,host-b) <--> qtunnel (host-a) <--> redis-server

After this, you can communicate over a encrypted wrapper rather than clear text.

### Credits

Special thanks to [Paul](http://paulrosenzweig.com) for reviewing the code.

### Contributing

We encourage you to contribute to `qtunnel`! Please feel free to [submit a bug report](https://github.com/getqujing/qtunnel/issues), [fork the repo](https://github.com/getqujing/qtunnel/fork) or [create a pull request](https://github.com/getqujing/qtunnel/pulls).

### License

`qtunnel` is released under the [Apache License 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).

