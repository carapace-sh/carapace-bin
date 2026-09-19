package hg

import (
	"bufio"
	"bytes"
	"os"
	"strings"

	"github.com/carapace-sh/carapace"
)

// Aliases returns the configured command aliases
func Aliases(cwd, repo string) (map[string]string, error) {
	aliases := make(map[string]string)

	c := carapace.Context{Env: os.Environ()}
	args := []string{}
	if cwd != "" {
		args = append(args, "--cwd", cwd)
	}
	if repo != "" {
		args = append(args, "--repository", repo)
	}
	args = append(args, "config", "alias")

	if output, err := c.Command("hg", args...).Output(); err != nil {
		return nil, err
	} else {
		scanner := bufio.NewScanner(bytes.NewReader(output))
		for scanner.Scan() {
			if key, value, found := strings.Cut(scanner.Text(), "="); found && strings.HasPrefix(key, "alias.") {
				aliases[strings.TrimPrefix(key, "alias.")] = value
			}
		}
	}

	return aliases, nil
}
