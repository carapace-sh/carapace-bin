package youtubedl

import (
	"regexp"
	"strings"
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/cache"
)

// ActionFormats completes formats
//   278 (webm       256x144    144p   74k , webm_dash container, vp9@  74k, 25fps, vid...)
//   313 (webm       3840x2160  2160p 16135k , webm_dash container, vp9@16135k, 25fps, ...)
func ActionFormats(url string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("youtube-dl", "--list-formats", url)(func(output []byte) carapace.Action {
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
	})
}

// ActionSubLangs completes subtitle languages
//   af (vtt, ttml, srv3, srv2, srv1)
//   am (vtt, ttml, srv3, srv2, srv1)
func ActionSubLangs(url string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("youtube-dl", "--list-subs", url)(func(output []byte) carapace.Action {
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
	})
}
