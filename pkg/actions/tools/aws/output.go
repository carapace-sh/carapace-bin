package aws

import "github.com/rsteube/carapace"

// ActionOutputFormats completes output formats
//
//	json (The output is formatted as a JSON string)
//	yaml (The output is formatted as a YAML string)
func ActionOutputFormats() carapace.Action {
	return carapace.ActionValuesDescribed(
		"json", "The output is formatted as a JSON string",
		"yaml", "The output is formatted as a YAML string",
		"yaml-stream", "The output is streamed and formatted as a YAML string",
		"text", "The output is formatted as multiple lines of tab-separated string values",
		"table", "The output is formatted as a table",
	)
}
