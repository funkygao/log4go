package main

import (
	"time"

	log "github.com/funkygao/log4go"
)

func init() {

}

func main() {
	log.DeleteFilter("stdout")

	rotateEnabled, discardWhenDiskFull := true, true
	filer := log.NewFileLogWriter("demo.log", rotateEnabled, discardWhenDiskFull, 0644)
	filer.SetFormat("[%d %T] [%L] (%S) %M")
	filer.SetRotateDaily(true)
	filer.SetRotateKeepDuration(time.Minute) // keep old logs only 1m
	log.AddFilter("file", log.DEBUG, filer)

	filer.Rotate()

	log.Info("hello world")

	// demonstrate feature: remove dup lines
	for i := 0; i < 100; i++ {
		log.Debug("hello golang!")
	}

	log.Trace("bye!")

	// flush all to disk
	log.Close()
}
