package logruzio

import (
	"net"
	"net/url"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	endpoint = "listener.logz.io:5050"
	proto    = "tcp"
)

// Hook represents a Logrus Logzio hook
type Hook struct {
	sync.Mutex
	Conn      net.Conn
	Context   logrus.Fields
	token     string
	Formatter logrus.Formatter
	ttl       time.Duration
}

// New creates a default Logzio hook.
// What it does is taking `token` and `appName` and attaching them to the log data.
// In addition, it sets a connection to the Logzio's Logstash endpoint.
// If the connection fails, it returns an error.
//
// To set more advanced configurations, initialize the hook in the following way:
//
// hook := &Hook{HookOpts{
//		Conn: myConn,
//		Context: logrus.Fields{...},
//		Formatter: myFormatter{}
// }
func New(token string, appName string, ttl time.Duration, ctx logrus.Fields) (*Hook, error) {
	h := &Hook{Context: logrus.Fields{}, token: token, ttl: ttl}

	h.Context["app"] = appName
	h.Context["meta"] = ctx
	h.Formatter = &SimpleFormatter{}

	return h, nil
}

// Fire writes `entry` to Logzio
func (h *Hook) Fire(entry *logrus.Entry) error {
	if h.Conn == nil {
		h.Lock()
		conn, err := net.DialTimeout(proto, endpoint, h.ttl)
		if err != nil {
			defer h.Unlock()
			return err
		}
		h.Conn = conn
		defer h.Unlock()
	}

	// Add in context fields.
	for k, v := range h.Context {
		// We don't override fields that are already set
		if _, ok := entry.Data[k]; !ok {
			entry.Data[k] = v
		}
	}

	dataBytes, err := h.Formatter.Format(entry)
	if err != nil {
		return err
	}

	_, err = h.Conn.Write(dataBytes)

	switch err := err.(type) {
	case net.Error:
		if err.Timeout() {
			h.Conn = nil
			return h.Fire(entry)
		}
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			h.Conn = nil
			return h.Fire(entry)
		}
	}

	if err != nil {
		return err
	}

	return nil
}

// Levels returns logging levels
func (h *Hook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}
