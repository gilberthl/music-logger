package main

import (
	"music-logger/pkg/listener"
	"music-logger/pkg/song"
	"os"
	"os/signal"
)

const (
	deviceID = 2
)

func main() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	logCh := make(chan interface{}, 1)
	l := listener.LogListener{
		Chan: logCh,
	}
	go l.Listen("../../fake-logger/log/app.log")
	song.Start(deviceID, logCh)
	<-signalCh
}
