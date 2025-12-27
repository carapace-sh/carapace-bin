package lazyinit

import (
	"sort"

	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd/completers"
	"github.com/carapace-sh/carapace-bridge/pkg/bridges"
	"github.com/carapace-sh/carapace-bridge/pkg/choices"
	envbridges "github.com/carapace-sh/carapace-bridge/pkg/env"
)

func Snippet(shell string) string {
	uniqueNames := map[string]bool{
		"carapace": true,
	}

	for _, b := range envbridges.Bridges() {
		if b == shell {
			continue // don't register native completions
		}
		switch b {
		case "bash":
			for _, name := range bridges.Bash() {
				uniqueNames[name] = true
			}
		case "fish":
			for _, name := range bridges.Fish() {
				uniqueNames[name] = true
			}
		case "inshellisense":
			for _, name := range bridges.Inshellisense() {
				uniqueNames[name] = true
			}
		case "zsh":
			for _, name := range bridges.Zsh() {
				uniqueNames[name] = true
			}
		}
	}

	switch shell { // don't bridge native completions
	case "bash":
		for _, name := range bridges.Bash() {
			delete(uniqueNames, name)
		}
	case "fish":
		for _, name := range bridges.Fish() {
			delete(uniqueNames, name)
		}
	case "zsh":
		for _, name := range bridges.Zsh() {
			delete(uniqueNames, name)
		}
	}

	if list, err := choices.List(false); err == nil {
		for _, choice := range list {
			// TODO filter by variant?? would need to read contents for that
			//  if s != shell { // don't bridge native completions
			uniqueNames[choice.Name] = true

		}
	}

	m, err := completers.Completers(choices.Choice{}, false)
	if err != nil {
		panic(err.Error()) // TODO handle  errror
	}
	// for _, name := range completers.Names() {
	for name := range m { // TODO already includes bridges and such
		uniqueNames[name] = true
	}

	completerNames := make([]string, 0)
	for name := range uniqueNames {
		completerNames = append(completerNames, name)
	}
	sort.Strings(completerNames)

	switch shell {
	case "bash":
		return Bash(completerNames)
	case "bash-ble":
		return BashBle(completerNames)
	case "cmd-clink":
		return CmdClink(completerNames)
	case "elvish":
		return Elvish(completerNames)
	case "fish":
		return Fish(completerNames)
	case "nushell":
		return Nushell(completerNames)
	case "oil":
		return Oil(completerNames)
	case "powershell":
		return Powershell(completerNames)
	case "tcsh":
		return Tcsh(completerNames)
	case "xonsh":
		return Xonsh(completerNames)
	case "zsh":
		return Zsh(completerNames)
	default:
		return ""
	}
}
