package main

import (
	w "github.com/burgesQ/webfmwk/v2"
	"github.com/burgesQ/webfmwk/v2/log"
)

// GetLogger return a log.ILog interface
var logger = log.GetLogger()

func main() {
	// init server w/ ctrl+c support
	s := w.InitServer(true)

	s.SetLogger(logger)

	s.GET("/test", func(c w.IContext) error {
		return c.JSONOk("ok")
	})

	// start asynchronously on :4242
	go func() {
		s.StartTLS(":4242", TLSConfig{
			Cert:     "/path/to/cert",
			Key:      "/path/to/key",
			Insecure: true,
		})
	}()

	// ctrl+c is handled internaly
	defer s.WaitAndStop()
}
