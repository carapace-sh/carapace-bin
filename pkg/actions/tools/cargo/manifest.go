package cargo

import (
	"encoding/json"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
)

type manifestJson struct {
	Dependencies []struct {
		Name string
		Req  string
	}
	Features map[string][]string

	Targets []struct {
		Name string
		Kind []string
	}
}

func readManifest(path string) (m manifestJson, err error) {
	var output []byte
	if output, err = (carapace.Context{}).Command("cargo", "read-manifest", "--offline", "--manifest-path", path).Output(); err == nil { // TODO read-manifest is deprecated
		err = json.Unmarshal(output, &m)
	}
	return
}

func readManifestAction(path string, f func(m manifestJson, args []string) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if path == "" {
			var err error
			if path, err = util.FindReverse(c.Dir, "Cargo.toml"); err != nil {
				return carapace.ActionMessage(err.Error())
			}
		}

		if m, err := readManifest(path); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return f(m, c.Args)
		}
	})
}
