package log4go

type SyslogNgWriter chan *LogRecord

func NewSyslogNgWriter() SyslogNgWriter {
	records := make(SyslogNgWriter, LogBufferLength)
	go records.run()
	return records
}

func (w SyslogNgWriter) LogWrite(rec *LogRecord) {
	w <- rec
}

func (w SyslogNgWriter) Close() {
	close(w)
}

func (w SyslogNgWriter) run() {
	for rec := range w {
		_ = rec

	}
}
