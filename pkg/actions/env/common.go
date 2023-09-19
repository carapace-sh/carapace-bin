package env

import (
	_os "os"

	"github.com/rsteube/carapace"
)

func init() {
	knownVariables["common"] = variables{
		Names: map[string]string{
			"PATH": "A list of directories to be searched when executing commands",
		},
		Values: map[string]carapace.Action{
			"PATH": carapace.ActionDirectories().List(string(_os.PathListSeparator)).NoSpace(),
		},
	}

}
