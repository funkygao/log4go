package log4go

import (
	"fmt"
	"github.com/funkygao/golib/hack"
	"net"
	"time"
)

type SyslogNgWriter struct {
	ch       chan *LogRecord
	sockPath string
	tag      string
	conn     net.Conn
}

func NewSyslogNgWriter(sockPath string, tag string) (w *SyslogNgWriter, err error) {
	w = new(SyslogNgWriter)
	w.sockPath = sockPath
	w.tag = tag
	w.ch = make(chan *LogRecord, LogBufferLength)
	if w.conn, err = net.Dial("unix", sockPath); err != nil {
		w.conn = nil
		return
	}

	go w.run()
	return
}

func (w SyslogNgWriter) LogWrite(rec *LogRecord) {
	if w.conn == nil {
		return
	}

	w.ch <- rec
}

func (w SyslogNgWriter) Close() {
	close(w.ch)
	if w.conn != nil {
		w.conn.Close()
	}
}

func (w SyslogNgWriter) run() {
	var line string
	for rec := range w.ch {
		line = fmt.Sprintf("%s,%d,%s\n", w.tag, rec.Created.Unix(), rec.Message)
		w.conn.SetWriteDeadline(time.Now().Add(time.Second * 2)) // loggr shouldn't block app
		w.conn.Write(hack.Byte(line))
	}
}
