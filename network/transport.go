package network

import (
	"net"
	"net/http"
	"time"
)

type TimingTransport struct {
	rtp       http.RoundTripper
	dialer    *net.Dialer
	connStart time.Time
	connEnd   time.Time
	reqStart  time.Time
	reqEnd    time.Time
}

func NewTimingTransport(timeout int) *TimingTransport {
	tr := &TimingTransport{
		dialer: &net.Dialer{
			Timeout:   time.Duration(timeout) * time.Millisecond,
			KeepAlive: 300 * time.Second,
		},
	}
	tr.rtp = &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		Dial:                tr.dial,
		TLSHandshakeTimeout: time.Duration(timeout/2) * time.Millisecond,
	}
	return tr
}

func (tr *TimingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	tr.reqStart = time.Now()
	resp, err := tr.rtp.RoundTrip(r)
	tr.reqEnd = time.Now()
	return resp, err
}

func (tr *TimingTransport) dial(network, addr string) (net.Conn, error) {
	tr.connStart = time.Now()
	cn, err := tr.dialer.Dial(network, addr)
	tr.connEnd = time.Now()
	return cn, err
}

func (tr *TimingTransport) ReqDuration() time.Duration {
	return tr.Duration() - tr.ConnDuration()
}

func (tr *TimingTransport) ConnDuration() time.Duration {
	return tr.connEnd.Sub(tr.connStart)
}

func (tr *TimingTransport) Duration() time.Duration {
	return tr.reqEnd.Sub(tr.reqStart)
}
