package log

import (
	"sync"

	"store/infra/log/gosimplefilelog"
)

type ILogFactory interface {
	Init(arg string) gosimplefilelog.Logger
}

type LogFactory struct {
}

var mutex sync.Mutex

func New() *LogFactory {
	return &LogFactory{}
}

func (lf LogFactory) Init(arg string) *gosimplefilelog.Logger {
	mutex.Lock()

	log := gosimplefilelog.NewLogger(
		gosimplefilelog.NewFileLogger(
			arg,
			""),
		gosimplefilelog.LEVEL_LOG)

	mutex.Unlock()

	return log
}
