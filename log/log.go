package log

import (
	"os"
	"strconv"
	"sync"
	"time"

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

	d := time.Now()

	hoje := str.Format("{0}.{1}.{2}", strconv.Itoa(d.Day()), d.Month().String(), strconv.Itoa(d.Year()))

	f, _ := os.OpenFile(str.Format("/home/pabx/{0}.log", hoje), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	f.Write([]byte(str.Format("{0} - {1}\n\n", datetime.Now().String(), message)))
}
