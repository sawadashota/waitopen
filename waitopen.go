package waitopen

import (
	"time"
	"github.com/skratchdot/open-golang/open"
	"net/url"
	"net/http"
	"github.com/fatih/color"
)

const (
	DefaultTimeout = 30 * time.Second
	DefaultRetry   = 5
)

type Opener struct {
	URL      *url.URL
	Interval time.Duration
	Retry    int
}

// To change default timeout and retry...
type Option func(opener *Opener)

// Create Opener Object
func New(URL *url.URL) *Opener {

	return &Opener{
		URL:      URL,
		Interval: DefaultTimeout,
		Retry:    DefaultRetry,
	}
}

// Change Interval
func SetInterval(timeout int) Option {
	return func(opener *Opener) {
		opener.Interval = time.Duration(timeout) * time.Second
	}
}

// Change Interval
func SetRetry(retry int) Option {
	return func(opener *Opener) {
		opener.Retry = retry
	}
}

// Wait until access then open URL
func (o *Opener) WaitOpen(options ...Option) {
	if o.Wait(options...) {
		o.Open()
	}
}

// Wait until access
func (o *Opener) Wait(options ...Option) bool {
	for _, option := range options {
		option(o)
	}

	count := 1

	for {
		color.Green("Trying access...")

		err := o.canAccess()

		if err == nil {
			return true
		}

		color.Red("%v\n\n", err.Error())

		if count > o.Retry {
			return false
		}

		time.Sleep(o.Interval)
		count++
	}

	return false
}

// Open URL in browser
func (o *Opener) Open() {
	color.Green("OK. Open %s", o.URL.String())
	open.Run(o.URL.String())
}

// Access the URL and return error
func (o *Opener) canAccess() error {
	_, err := http.Get(o.URL.String())
	return err
}
