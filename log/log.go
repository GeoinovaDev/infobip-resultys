package log

import (
	"os"
	"sync"

	"git.resultys.com.br/lib/lower/str"
	"git.resultys.com.br/lib/lower/time/datetime"
)

// Log ...
type Log struct {
	mutex *sync.Mutex
}

var current *Log

// GetInstance ...
func GetInstance() *Log {
	if current == nil {
		current = &Log{
			mutex: &sync.Mutex{},
		}
	}

	return current
}

// Add ...
func (l *Log) Add(message string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	f, _ := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	f.Write([]byte(str.Format("{0} - {1}\n", datetime.Now().String(), message)))
}
