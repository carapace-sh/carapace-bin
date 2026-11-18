package ssh

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-shlex/v2"
)

// ActionHosts completes ssh hosts
//
//	github.com
//	example.com
func ActionHosts(config string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		configs := map[string]string{config: "~/.ssh"}
		if config == "" {
			configs = map[string]string{"~/.ssh/config": "~/.ssh", "/etc/ssh/ssh_config": "/etc/ssh"}
		}

		vals := make([]string, 0)
		for file, base := range configs {
			path, err := c.Abs(file)
			if err != nil {
				continue
			}
			c.Dir = base
			for _, s := range configHosts(c, path, make(map[string]bool)) {
				if s != "" && !strings.ContainsAny(s, `+*?()|[]{}^$!`+" \t") {
					vals = append(vals, s)
				}
			}
		}
		return carapace.ActionValues(vals...)
	}).Tag("ssh hosts")
}

// Collect names without evaluating Match conditions (particularly Match exec).
func configHosts(c carapace.Context, path string, seen map[string]bool) []string {
	if seen[path] {
		return nil
	}
	seen[path] = true
	content, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	var hosts []string
	for line := range strings.SplitSeq(string(content), "\n") {
		line = strings.TrimSpace(line)
		i := strings.IndexAny(line, " \t=")
		if i < 0 {
			continue
		}
		key := strings.ToLower(line[:i])
		if key != "host" && key != "match" && key != "include" {
			continue
		}
		tokens, err := shlex.Split(strings.TrimPrefix(strings.TrimSpace(line[i:]), "="), shlex.Bash)
		if err != nil {
			continue
		}
		args := tokens.Words().Strings()
		switch key {
		case "host":
			hosts = append(hosts, args...)
		case "match":
			for len(args) > 0 {
				criterion := strings.ToLower(args[0])
				args = args[1:]
				switch strings.TrimPrefix(criterion, "!") {
				case "all", "canonical", "final":
					continue
				}
				if len(args) == 0 {
					break
				}
				if criterion == "host" || criterion == "originalhost" {
					hosts = append(hosts, strings.Split(args[0], ",")...)
				}
				args = args[1:]
			}
		case "include":
			for _, pattern := range args {
				pattern, err := c.Abs(pattern)
				if err != nil {
					continue
				}
				files, _ := filepath.Glob(pattern)
				for _, file := range files {
					hosts = append(hosts, configHosts(c, file, seen)...)
				}
			}
		}
	}
	return hosts
}
