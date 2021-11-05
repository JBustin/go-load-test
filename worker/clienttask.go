package worker

import (
	"net/http"

	"github.com/go-load-test/config"
	"github.com/go-load-test/network"
)

type clienttask struct {
	task
	Client    network.HTTPClient
	Transport *network.TimingTransport
}

func newClientTask(config config.Config, id int, url string, hasFailed bool) *clienttask {
	transport := network.NewTimingTransport(config.TimeoutMs)

	return &clienttask{
		Client:    &http.Client{Transport: transport},
		Transport: transport,
		task: task{
			config:    config,
			id:        id,
			url:       url,
			hasFailed: hasFailed,
		},
	}
}

func (t *clienttask) Request() error {
	req, err := http.NewRequest("GET", t.url, nil)
	if err != nil {
		return err
	}
	defer func() { req.Close = true }()
	for key, value := range t.config.Headers {
		req.Header.Add(key, value)
	}
	resp, err := t.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	t.statusCode = resp.StatusCode
	t.connectDuration = t.Transport.ConnDuration()
	t.reqDuration = t.Transport.ReqDuration()
	t.duration = t.Transport.Duration()

	return nil
}
