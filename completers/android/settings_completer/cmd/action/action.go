package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionNamespaces completes setting namespaces
//
//	system
//	secure
//	global
func ActionNamespaces() carapace.Action {
	return carapace.ActionValuesDescribed(
		"system", "system settings",
		"secure", "secure settings",
		"global", "global settings",
	).Tag("namespaces")
}

// ActionKeys completes setting keys of the given namespace
//
//	enabled_notification_listeners (...)
//	location_providers_allowed (...)
func ActionKeys(namespace string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if namespace == "" {
			return carapace.ActionValues()
		}
		return carapace.ActionExecCommand("settings", "list", namespace)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)
			for _, line := range lines {
				if splitted := strings.SplitN(line, "=", 2); len(splitted) == 2 {
					vals = append(vals, splitted[0], splitted[1])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		}).Tag("setting keys")
	})
}
