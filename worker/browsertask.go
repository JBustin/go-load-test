package worker

import (
	"strconv"
	"time"

	"github.com/go-load-test/config"
	"github.com/mxschmitt/playwright-go"
)

type browsertask struct {
	task
	options playwright.RunOptions
}

func newBrowserTask(config config.Config, id int, url string, hasFailed bool) *browsertask {
	return &browsertask{
		options: playwright.RunOptions{
			SkipInstallBrowsers: true,
			Browsers:            []string{"chromium"},
		},
		task: task{
			config:    config,
			id:        id,
			url:       url,
			hasFailed: hasFailed,
		},
	}
}

func (t *browsertask) Request() error {
	pw, err := playwright.Run(&t.options)
	if err != nil {
		return err
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		return err
	}
	page, err := browser.NewPage()
	if err != nil {
		return err
	}

	if err = page.SetExtraHTTPHeaders(t.config.Headers); err != nil {
		return err
	}
	resp, err := page.Goto(t.url)
	if err != nil {
		return err
	}
	handle, err := page.EvaluateHandle("window.performance.timing", struct{}{})
	if err != nil {
		return err
	}

	connectStart, err := getTimingFrom(handle, "connectStart")
	if err != nil {
		return err
	}
	connectEnd, err := getTimingFrom(handle, "connectEnd")
	if err != nil {
		return err
	}
	responseStart, err := getTimingFrom(handle, "responseStart")
	if err != nil {
		return err
	}
	responseEnd, err := getTimingFrom(handle, "responseEnd")
	if err != nil {
		return err
	}

	t.statusCode = resp.Status()
	t.connectDuration = time.Duration(connectEnd-connectStart) * time.Millisecond
	t.reqDuration = time.Duration(responseEnd-responseStart) * time.Millisecond
	t.duration = t.connectDuration + t.reqDuration

	if err = browser.Close(); err != nil {
		return err
	}
	if err = pw.Stop(); err != nil {
		return err
	}
	return nil
}

func getTimingFrom(handle playwright.JSHandle, name string) (int, error) {
	timingJsHandle, err := handle.GetProperty(name)
	if err != nil {
		return 0, err
	}
	timing, err := strconv.Atoi(timingJsHandle.String())
	if err != nil {
		return 0, err
	}
	return timing, err
}
