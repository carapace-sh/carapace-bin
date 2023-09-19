package env

import (
	_os "os"

	"github.com/rsteube/carapace"
)

func init() {
	knownVariables["common"] = variables{
		Names: map[string]string{
			"HTTP_PROXY":   "http proxy server",
			"HTTP_TIMEOUT": "The HTTP timeout in seconds",
			"HTTPS_PROXY":  "https proxy server",
			"PATH":         "A list of directories to be searched when executing commands",
		},
		Values: map[string]carapace.Action{
			"PATH": carapace.ActionDirectories().List(string(_os.PathListSeparator)).NoSpace(),
		},
	}

}
