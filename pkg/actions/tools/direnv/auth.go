package direnv

import (
	"bufio"
	"fmt"
	"os"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionAuths completes authorizations
//
//	/project/dir/.envrc
//	/home/user/project/.envrc
func ActionAuths() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO support XDG folders: https://github.com/direnv/direnv/blob/729fbecd96f3e827575f19497bc01df33395d679/xdg/xdg.go
		dir, err := c.Abs("~/.local/share/direnv/allow/")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		entries, err := os.ReadDir(dir)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, entry := range entries {
			if !entry.IsDir() {
				auth, err := readFirstLine(dir + entry.Name())
				if err != nil {
					return carapace.ActionMessage(err.Error())
				}

				vals = append(vals, auth)
			}
		}
		return carapace.ActionValues(vals...).StyleF(style.ForPath)
	})
}

func readFirstLine(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	return "", fmt.Errorf("file is empty: %v", path)
}
