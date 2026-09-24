package wt

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/carapace-sh/carapace"
	toml "github.com/pelletier/go-toml"
)

// ActionAliasNames completes worktrunk alias names
//
//	wsc
//	wtpr
func ActionAliasNames() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommandE("git", "rev-parse", "--show-toplevel")(func(output []byte, err error) carapace.Action {
			paths := make(map[string]struct{})

			if path := userConfigPath(); path != "" {
				paths[path] = struct{}{}
			}
			if err == nil {
				if path := projectConfigPath(strings.TrimSpace(string(output))); path != "" {
					paths[path] = struct{}{}
				}
			}

			names := make(map[string]struct{})
			for path := range paths {
				for _, name := range aliasNames(path) {
					names[name] = struct{}{}
				}
			}

			vals := make([]string, 0, len(names))
			for name := range names {
				vals = append(vals, name)
			}
			sort.Strings(vals)
			return carapace.ActionValues(vals...).Tag("aliases")
		})
	})
}

func userConfigPath() string {
	if path, exists := os.LookupEnv("WORKTRUNK_CONFIG_PATH"); exists && path != "" {
		return path
	}
	if base, err := os.UserConfigDir(); err == nil {
		return filepath.Join(base, "worktrunk", "config.toml")
	}
	return ""
}

func projectConfigPath(root string) string {
	if path, exists := os.LookupEnv("WORKTRUNK_PROJECT_CONFIG_PATH"); exists && path != "" {
		if filepath.IsAbs(path) {
			return path
		}
		return filepath.Join(root, path)
	}
	return filepath.Join(root, ".config", "wt.toml")
}

func aliasNames(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	tree, err := toml.LoadBytes(content)
	if err != nil {
		return nil
	}
	aliases := tree.Get("aliases")
	table, ok := aliases.(*toml.Tree)
	if !ok {
		return nil
	}
	return table.Keys()
}
