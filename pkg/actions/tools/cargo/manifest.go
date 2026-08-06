package cargo

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/util"
)

type metadataJson struct {
	Packages []struct {
		Name         string
		Version      string
		Dependencies []struct {
			Name string
			Req  string
		}
		Features map[string][]string
		Targets  []struct {
			Name string
			Kind []string
		}
	}
	WorkspaceMembers []string
}

func readMetadata(path string) (m metadataJson, err error) {
	var output []byte
	if output, err = (carapace.Context{}).Command("cargo", "metadata", "--no-deps", "--format-version", "1", "--manifest-path", path).Output(); err == nil {
		err = json.Unmarshal(output, &m)
	}
	return
}

func readMetadataAction(path string, f func(m metadataJson, args []string) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if path == "" {
			var err error
			if path, err = util.FindReverse(c.Dir, "Cargo.toml"); err != nil {
				return carapace.ActionMessage(err.Error())
			}
		}

		if m, err := readMetadata(path); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return f(m, c.Args)
		}
	})
}
