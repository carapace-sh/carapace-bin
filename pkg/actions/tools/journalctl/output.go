package journalctl

import "github.com/carapace-sh/carapace"

// ActionOutputs completes output modes
//
//	cat (generates a very terse output)
//	export (serializes the journal into a binary)
func ActionOutputs() carapace.Action {
	return carapace.ActionValuesDescribed(
		"short", "generates an output that is mostly identical to the formatting of classic syslog files",
		"short-full", "shows timestamps in the format the --since= and --until= options accept",
		"short-iso", "shows ISO 8601 wallclock timestamps.",
		"short-iso-precise", "as for short-iso but includes full microsecond precision.",
		"short-precise", "shows classic syslog timestamps with full microsecond precision",
		"short-monotonic", "shows monotonic timestamps instead of wallclock timestamps",
		"short-unix", "shows seconds passed since January 1st 1970 UTC",
		"verbose", "shows the full-structured entry items with all fields",
		"export", "serializes the journal into a binary",
		"json", "formats entries as JSON objects",
		"json-pretty", "formats entries as JSON data structures",
		"json-sse", "formats entries as JSON data structures",
		"json-seq", "formats entries as JSON data structures",
		"cat", "generates a very terse output",
		"with-unit", "similar to short-full, but prefixes the unit and user unit names",
	)
}
