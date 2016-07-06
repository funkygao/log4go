package main

import (
	"time"

	log "github.com/funkygao/log4go"
)

var (
	filer *log.FileLogWriter
)

func init() {
	log.LogBufferLength = 1 << 10

	log.DeleteFilter("stdout")

	rotateEnabled, discardWhenDiskFull := true, false
	filer = log.NewFileLogWriter("_demo.log", rotateEnabled, discardWhenDiskFull, 0644)
	filer.SetFormat("[%d %T] [%L] (%S) %M")
	filer.SetRotateDaily(true)
	filer.SetRotateKeepDuration(time.Minute) // keep old logs only 1m
	log.AddFilter("file", log.DEBUG, filer)
}

func main() {
	log.Info("hello world")

	// demonstrate feature: remove dup lines
	for i := 0; i < 100; i++ {
		log.Debug("hello golang!")
	}

	log.Trace("bye!")

	filer.Rotate()

	// flush all to disk
	log.Close()
}
