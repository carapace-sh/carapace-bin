package env

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
)

func init() {
	knownVariables["common_unix"] = variables{
		Names: map[string]string{
			"USER":    "The current logged in user",
			"HOME":    "The home directory of the current user",
			"EDITOR":  "The default file editor to be used",
			"SHELL":   "The path of the current user’s shell, such as bash or zsh",
			"LOGNAME": "The name of the current user",
			"LANG":    "The current locales settings",
			"TERM":    "The current terminal emulation",
			"MAIL":    "Location of where the current user’s mail is stored",
		},
		Completion: map[string]carapace.Action{
			"HOME":    carapace.ActionDirectories(),
			"LANG":    os.ActionLanguages(),
			"LOGNAME": os.ActionUsers(),
			"USER":    os.ActionUsers(),
		},
	}

}
