package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/dnovikoff/mimage/pkg/image"
)

var (
	netFlag    = flag.String("net", "tcp", "network")
	addrFlag   = flag.String("addr", ":8080", "address")
	spriteFlag = flag.String("sprite", "pkg/image/test_data/sprite.png", "sprites location")
	maxLenFlag = flag.Int("maxlen", 50, "max symbols per string")
)

func main() {
	flag.Parse()
	img, err := image.LoadPNG(*spriteFlag)
	check(err)
	sprites := image.DefaultSprites(img)
	lis, err := net.Listen(*netFlag, *addrFlag)
	check(err)
	check(http.Serve(lis, &image.Handler{
		Sprites: sprites,
		MaxLen:  *maxLenFlag,
	}))
}

func check(err error) {
	if err == nil {
		return
	}
	log.Fatal(err)
}
