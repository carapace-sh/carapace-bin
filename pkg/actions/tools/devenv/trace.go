package devenv

import (
	"slices"
	"strings"

	"github.com/carapace-sh/carapace"
)

var traceFormats = []string{
	"full", "verbose human readable output",
	"json", "json output (default)",
	"otlp-grpc", "OTLP over gRPC",
	"otlp-http-json", "OTLP over HTTP (json)",
	"otlp-http-protobuf", "OTLP over HTTP (protobuf)",
	"pretty", "human readable output",
}

func traceFormatNames() []string {
	names := make([]string, 0, len(traceFormats)/2)
	for index := 0; index < len(traceFormats); index += 2 {
		names = append(names, traceFormats[index])
	}
	return names
}

// ActionTraceTargets completes tracing targets in `[format:]destination` format
//
//	pretty: (human readable output)
//	stderr (write to stderr)
func ActionTraceTargets() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if format, _, ok := strings.Cut(c.Value, ":"); ok && slices.Contains(traceFormatNames(), format) {
			if strings.HasPrefix(format, "otlp-") {
				return actionTraceEndpoints().Prefix(format + ":")
			}
			return actionTraceDestinations().Prefix(format + ":")
		}

		return carapace.Batch(
			carapace.ActionValuesDescribed(traceFormats...).Suffix(":").NoSpace(':'),
			actionTraceDestinations(),
		).ToA()
	}).Tag("trace targets")
}

func actionTraceDestinations() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if strings.HasPrefix(c.Value, "file:") {
			return carapace.ActionFiles().Prefix("file:")
		}

		return carapace.ActionValuesDescribed(
			"file:", "write to a file",
			"stderr", "write to stderr",
			"stdout", "write to stdout",
		).NoSpace(':')
	})
}

// actionTraceEndpoints completes the collector endpoint of an OTLP format
func actionTraceEndpoints() carapace.Action {
	return carapace.ActionValues("http://", "https://").NoSpace('/')
}
