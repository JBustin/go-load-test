package mocks

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type Task struct {
	mock.Mock
	StatusCode      int
	connectDuration time.Duration
	reqDuration     time.Duration
	duration        time.Duration
	hasFailed       bool
}

func (t Task) String() string {
	args := t.Called()
	return args.String(0)
}

func (t Task) RequestStr() string {
	args := t.Called()
	return args.String(0)
}

func (t Task) ResponseStr() string {
	args := t.Called()
	return args.String(0)
}

func (t Task) ErrorStr(err error) string {
	args := t.Called(err)
	return args.String(0)
}

func (t Task) GetStatusCode() int {
	return t.StatusCode
}

func (t Task) GetDuration(name string) time.Duration {
	switch name {
	case "connect":
		return t.connectDuration
	case "request":
		return t.reqDuration
	default:
		return t.duration
	}
}

func (t *Task) SetHasFailed(flag bool) {
	t.hasFailed = flag
}

func (t Task) HasFailed() bool {
	return t.hasFailed
}

func (t Task) Request() error {
	args := t.Called()
	return args.Error(0)
}
