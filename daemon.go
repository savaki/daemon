package daemon

import (
	"net/http"

	"contrib.go.opencensus.io/exporter/aws"
	"github.com/pkg/errors"
	"go.opencensus.io/trace"
)

// Run the handler as a web server on the specified address
func Run(addr string, h http.Handler) error {
	exporter, err := aws.NewExporter()
	if err != nil {
		return errors.Wrapf(err, "unable to start x-ray tracer")
	}
	defer exporter.Close()

	trace.RegisterExporter(exporter)

	return http.ListenAndServe(addr, h)
}
