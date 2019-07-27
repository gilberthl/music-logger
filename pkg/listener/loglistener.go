package listener

import (
	"github.com/hpcloud/tail"
)

type LogListener struct {
	Chan chan<- interface{}
}

func (l *LogListener) Listen(filepath string) {
	t, err := tail.TailFile(filepath, tail.Config{Follow: true})
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		l.Chan <- line.Text
	}
}
