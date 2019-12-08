# Trial By Pong
A pong game with a twist written in Go.

## Building

Right now, ego2d is still in active development and therefore doesn't really work well with modules.  

Make sure you have SDL installed on your machine (btw, this only works on Linux afaik)

On __Ubuntu 14.04 and above__, type:\
`apt install libsdl2{,-image,-mixer,-ttf,-gfx}-dev`

On __Fedora 25 and above__, type:\
`yum install SDL2{,_image,_mixer,_ttf,_gfx}-devel`

On __Arch Linux__, type:\
`pacman -S sdl2{,_image,_mixer,_ttf,_gfx}`

On __Gentoo__, type:\
`emerge -av libsdl2 sdl2-{image,mixer,ttf,gfx}`

Then install the Go SDL Bindings (this'll take a minute)
`go get -uv github.com/veandco/go-sdl2/{sdl,img,mix,ttf}`

You'll also need:
`go get -u github.com/hlfstr/{ego2d,logit,flagger}`

Then get and build this game (build from the 'main' folder)

```sh
go get github.com/hlfstr/tbp
cd $GOPATH/src/github.com/hlfstr/tbp/main
go build -o tbp
./tbp
```

Make sure to run the binary from the same directory as the 'assets' folder.

Running just `./tbp` will tell you no commands passed, and the `./tbp help` is currently broken.  The only working method right now is `./tbp single`
