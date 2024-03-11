package env

import (
	_os "os"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
)

func init() {
	knownVariables["common"] = func() variables {
		return variables{
			Variables: map[string]string{
				"BROWSER":      "the browser to use",
				"EDITOR":       "the editor to use",
				"HTTP_PROXY":   "http proxy server",
				"HTTPS_PROXY":  "https proxy server",
				"HTTP_TIMEOUT": "The HTTP timeout in seconds",
				"PAGER":        "the pager to use",
				"PATH":         "A list of directories to be searched when executing commands",
			},
			VariableCompletion: map[string]carapace.Action{
				"BROWSER": bridge.ActionCarapaceBin().Split(),
				"EDITOR":  bridge.ActionCarapaceBin().Split(),
				"PAGER":   bridge.ActionCarapaceBin().Split(),
				"PATH":    carapace.ActionDirectories().List(string(_os.PathListSeparator)).NoSpace(),
			},
		}
	}
}
