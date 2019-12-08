# Trial By Pong
A pong game with a twist written in Go.

## Building

Right now, ego2d is still in active development and therefore doesn't really work well with modules.  To play this game, just do:

```sh
go get github.com/hlfstr/tbp
cd $GOPATH/src/github.com/hlfstr/tbp/main
go build -o tbp
./tbp
```

Make sure to run the binary from the same directory as the 'assets' folder.
