package main

import (
    "os"
    "os/signal"
    "syscall"
    "github.com/postcode/postcode"
    ylog "github.com/postcode/lib/ylog"
)

// Version - версия
var Version string

func main() {
    postcode.Ident = os.Args[0]
    PidFile := postcode.Init(Version)
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    go func() {
	<-sigs
	ylog.Destr()
	os.Remove(PidFile)
	os.Exit(0)
    }()

    ylog.YLog(1, postcode.Ident, "Start programm")

}
