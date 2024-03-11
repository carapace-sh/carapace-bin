package git

import (
	"bufio"
	"bytes"
	"os"
	"strings"

	"github.com/carapace-sh/carapace"
)

func Aliases(dir, gitDir string) (map[string]string, error) {
	aliases := make(map[string]string)

	c := carapace.Context{Env: os.Environ()}
	args := []string{}
	if dir != "" {
		args = append(args, "-C", dir)
	}
	if gitDir != "" {
		args = append(args, "--git-dir", gitDir)
	}
	args = append(args, "config", "--get-regexp", "^alias\\.")

	if output, err := c.Command("git", args...).Output(); err != nil {
		return nil, err
	} else {
		scanner := bufio.NewScanner(bytes.NewReader(output))
		for scanner.Scan() {
			if splitted := strings.SplitN(scanner.Text(), " ", 2); len(splitted) == 2 && strings.HasPrefix(splitted[0], "alias.") {
				name := strings.TrimPrefix(splitted[0], "alias.")

				// in the case of duplicates, last alias wins
				aliases[name] = splitted[1]
			}
		}
	}

	return aliases, nil
}

// ActionAliases completes aliases
//
//	po (push origin)
//	ct (checkout --track)
func ActionAliases(gitDir string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if aliases, err := Aliases(c.Dir, gitDir); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := make([]string, 0)
			for name, description := range aliases {
				vals = append(vals, name, description)
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}
