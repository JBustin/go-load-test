package worker

import (
	"fmt"
	"time"

	"github.com/go-load-test/config"
)

type Tasker interface {
	String() string
	RequestStr() string
	ResponseStr() string
	ErrorStr(error) string
	Request() error
	StatusCode() int
	Duration(string) time.Duration
	SetHasFailed(bool)
	HasFailed() bool
}

type task struct {
	config          config.Config
	id              int
	url             string
	statusCode      int
	connectDuration time.Duration
	reqDuration     time.Duration
	duration        time.Duration
	hasFailed       bool
}

func newTask(config config.Config, id int, url string, hasFailed bool) Tasker {
	if config.IsBrowser {
		return newBrowserTask(config, id, url, hasFailed)
	} else {
		return newClientTask(config, id, url, hasFailed)
	}
}

func (t task) String() string {
	return fmt.Sprintf(`
	Url: %v [%v]
	Status code: %v
	Timings: 
		- Connect: %v 
		- Request: %v 
		- Total: %v 
	`, t.url, t.id, t.statusCode, t.connectDuration, t.reqDuration, t.duration)
}

func (t task) RequestStr() string {
	return fmt.Sprintf("[%v] --> Request %v", t.id, t.url)
}

func (t task) ResponseStr() string {
	return fmt.Sprintf("[%v] <-- Duration=%v\t%v\n", t.id, t.duration, t.url)
}

func (t task) ErrorStr(err error) string {
	return fmt.Sprintf("[%v] ! %v  %v", t.id, t.url, err)
}

func (t task) StatusCode() int {
	return t.statusCode
}

func (t task) Duration(name string) time.Duration {
	switch name {
	case "connect":
		return t.connectDuration
	case "request":
		return t.reqDuration
	default:
		return t.duration
	}
}

func (t *task) SetHasFailed(flag bool) {
	t.hasFailed = flag
}

func (t task) HasFailed() bool {
	return t.hasFailed
}
