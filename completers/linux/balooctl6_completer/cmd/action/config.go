package action

import (
	"errors"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"gopkg.in/ini.v1"
)

func ActionConfigValues(param string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		param = strings.ToLower(param)

		config, err := ini.Load(filepath.Join(c.Dir, "baloofilerc"))
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				return carapace.ActionValues()
			}
			return carapace.ActionMessage(err.Error())
		}

		keys := map[string]string{
			"includefolders":   "folders",
			"excludefolders":   "exclude folders",
			"excludefilters":   "exclude filters",
			"excludemimetypes": "exclude mimetypes",
		}
		name, ok := keys[param]
		if !ok {
			return carapace.ActionValues()
		}

		var values []string
		for _, key := range config.Section("General").Keys() {
			keyName := key.Name()
			if index := strings.Index(keyName, "["); index >= 0 {
				keyName = keyName[:index]
			}
			if keyName != name {
				continue
			}
			escaped := strings.ReplaceAll(key.String(), `\,`, "\x00")
			for _, value := range strings.Split(escaped, ",") {
				if value = strings.TrimSpace(strings.ReplaceAll(value, "\x00", ",")); value != "" {
					values = append(values, value)
				}
			}
		}
		return carapace.ActionValues(values...)
	}).ChdirF(traverse.XdgConfigHome)
}
