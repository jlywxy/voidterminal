# void terminal
A socket terminal supports TLS.<br/>
This program is included in voidshell software kit.https://github.com/jlywxy/void

## build
```$ go clean; go build```

## usage
To connect a unix socket,
```shell
$ ./voidterminal unix:/path/to/socket
```
To connect tcp server,
```shell
$ ./voidterminal tcp:a.b.c.d:9000
```
To connect tcp over TLS,
```shell
$ ./voidterminal tcp:void-server.net:443 --tls
```


## equivalence
when connection over TLS, this program is equivalent to
```shell
$ stty raw; openssl s_client -connect tcp:void-server.net:443
```
or
```shell
$ stty raw; socat - openssl:void-server.net:443
```
or
```shell
$ stty raw; ncat --ssl void-server.net 443
```

If using raw connections (no TLS):
```shell
$ stty raw; nc -U /path/to/socket
```
or
```shell
$ stty raw; nc void-server.net 9000
```