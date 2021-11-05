package worker

import (
	"testing"
	"time"

	"github.com/go-load-test/config"
	"github.com/go-load-test/mocks"
	"github.com/go-load-test/network"
	"github.com/stretchr/testify/assert"
)

func Test_Clienttask(t *testing.T) {
	task := clienttask{
		task: task{
			config: config.Config{
				Headers: map[string]string{
					"x-custom-header-1": "abcd",
					"x-custom-header-2": "1234",
				},
			},
		},
		Client:    mocks.HTTPresponse(200, "", nil),
		Transport: network.NewTimingTransport(0),
	}

	err := task.Request()

	assert.Equal(
		t,
		nil,
		err,
		"Request should not return an error",
	)

	assert.Equal(
		t,
		struct {
			connectDuration time.Duration
			reqDuration     time.Duration
			duration        time.Duration
		}{
			connectDuration: 0,
			reqDuration:     0,
			duration:        0,
		},
		struct {
			connectDuration time.Duration
			reqDuration     time.Duration
			duration        time.Duration
		}{
			connectDuration: task.GetDuration("connect"),
			reqDuration:     task.GetDuration("request"),
			duration:        task.GetDuration(""),
		},
		"Request should populate iteration metrics",
	)
}
