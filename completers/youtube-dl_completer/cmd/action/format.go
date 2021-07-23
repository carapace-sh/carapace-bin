package action

import (
	"regexp"
	"strings"
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/cache"
)

func ActionFormats() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Args) == 0 {
			return carapace.ActionMessage("missing url")
		} else {
			return carapace.ActionExecCommand("youtube-dl", "--list-formats", c.Args[0])(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				pattern := regexp.MustCompile(`^(?P<format>[^ ]+) +(?P<description>.*)$`)

				found := false
				vals := make([]string, 0)
				for _, line := range lines {
					if !found && strings.HasPrefix(line, "format code") {
						found = true
						continue
					}
					if found {
						if pattern.MatchString(line) {
							m := pattern.FindStringSubmatch(line)
							vals = append(vals, m[1], m[2])
						}
					}
				}
				return carapace.ActionValuesDescribed(vals...)
			}).Cache(1*time.Hour, cache.String(c.Args[0]))
		}
	})
}

func ActionSubLangs() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Args) == 0 {
			return carapace.ActionMessage("missing url")
		} else {
			return carapace.ActionExecCommand("youtube-dl", "--list-subs", c.Args[0])(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				pattern := regexp.MustCompile(`^(?P<languate>[^ ]+) +(?P<formats>.*)$`)

				found := false
				vals := make([]string, 0)
				for _, line := range lines[:len(lines)-1] {
					if !found && line == "Language formats" {
						found = true
						continue
					}
					if found {
						if pattern.MatchString(line) {
							m := pattern.FindStringSubmatch(line)
							vals = append(vals, m[1], m[2])
						}
					}
				}
				return carapace.ActionValuesDescribed(vals...)
			}).Cache(1*time.Hour, cache.String(c.Args[0]))
		}
	})
}
