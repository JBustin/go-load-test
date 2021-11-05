package worker

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/go-load-test/config"
	"github.com/stretchr/testify/assert"
)

func Test_NewTask(t *testing.T) {
	conf := config.DefaultConfig
	task := newTask(conf, 0, "https://mysite.com", false)
	assert.Equal(t, "*worker.clienttask", reflect.TypeOf(task).String(), "should return a clienttask by default")

	conf.IsBrowser = true
	task = newTask(conf, 0, "https://mysite.com", false)
	assert.Equal(t, "*worker.browsertask", reflect.TypeOf(task).String(), "should return a clienttask by default")
}

func Test_TaskMethods(t *testing.T) {
	conf := config.DefaultConfig
	task := newTask(conf, 0, "https://mysite.com", false)

	assert.Equal(t, "[0] --> Request https://mysite.com", task.RequestStr(), "should format request message")
	assert.Equal(t, "[0] <-- 0s\thttps://mysite.com\n", task.ResponseStr(), "should format response message")
	assert.Equal(t, "[0] ! https://mysite.com  boum", task.ErrorStr(errors.New("boum")), "should format error message")
	assert.Equal(t, 0, task.GetStatusCode(), "should return statusCode")
	assert.Equal(t, 0*time.Millisecond, task.GetDuration(""), "should return total duration")
	assert.Equal(t, 0*time.Millisecond, task.GetDuration("request"), "should return request duration")
	assert.Equal(t, 0*time.Millisecond, task.GetDuration("connect"), "should return connect duration")

	assert.Equal(t, false, task.HasFailed(), "should return failed state")
	task.SetHasFailed(true)
	assert.Equal(t, true, task.HasFailed(), "should update failed state")
}
