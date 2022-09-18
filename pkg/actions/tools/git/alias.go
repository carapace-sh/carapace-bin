package git

import (
	"bufio"
	"bytes"
	"os"
	"strings"

	"github.com/rsteube/carapace"
)

func Aliases(gitArgs []string) (map[string]string, error) {
	aliases := make(map[string]string)

	c := carapace.Context{Env: os.Environ()}
	args := append(gitArgs, "config", "--get-regexp", "^alias\\.")

	if output, err := c.Command("git", args...).Output(); err != nil {
		return nil, err
	} else {
		scanner := bufio.NewScanner(bytes.NewReader(output))
		for scanner.Scan() {
			split := strings.SplitN(scanner.Text(), " ", 2)
			name := strings.TrimPrefix(split[0], "alias.")

			// in the case of duplicates, last alias wins
			aliases[name] = split[1]
		}
	}

	return aliases, nil
}
