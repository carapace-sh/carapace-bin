package kubeadm

import "github.com/carapace-sh/carapace"

// ActionOutputFormats completes output formats
//
//	go-template
//	json
func ActionOutputFormats() carapace.Action {
	return carapace.ActionValues(
		"go-template",
		"go-template-file",
		"json",
		"jsonpath",
		"jsonpath-as-json",
		"jsonpath-file",
		"template",
		"templatefile",
		"text",
		"yaml",
	).Tag("output formats")
}
