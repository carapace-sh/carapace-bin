package supervisor

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"gopkg.in/ini.v1"
)

// ActionGroups completes groups
//
//	cat (/bin/cat)
//	sleep (/bin/sleep 10m)
func ActionGroups(path string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if path == "" {
			var err error
			if path, err = configPath(c); err != nil {
				return carapace.ActionMessage(err.Error())
			}
		}

		cfg, err := ini.Load(path)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, section := range cfg.Sections() {
			carapace.LOG.Println(section.Name())
			if after, ok := strings.CutPrefix(section.Name(), "program:"); ok {
				vals = append(vals, after, section.KeysHash()["command"])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
