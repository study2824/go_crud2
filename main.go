package main

import (
	"github.com/koko2824/goSample/db"
	"github.com/koko2824/goSample/server"
)

func main() {
	db.Init()
	server.Init()
}