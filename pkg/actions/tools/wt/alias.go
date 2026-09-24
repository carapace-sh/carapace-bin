package wt

import (
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
	toml "github.com/pelletier/go-toml"
)

// ActionAliasNames completes worktrunk alias names
//
//	wsc
//	wtpr
func ActionAliasNames() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		names := make(map[string]struct{})

		for _, path := range systemConfigPaths(c) {
			addAliasNames(path, names)
		}

		if path, ok := c.LookupEnv("WORKTRUNK_CONFIG_PATH"); ok && path != "" {
			addAliasNames(path, names)
		} else if base, err := traverse.XdgConfigHome(c); err == nil {
			addAliasNames(filepath.Join(base, "worktrunk", "config.toml"), names)
		}

		if root, err := traverse.GitWorkTree(c); err == nil {
			if path := projectConfigPath(root, c); path != "" {
				addAliasNames(path, names)
			}
		}

		vals := make([]string, 0, len(names))
		for name := range names {
			vals = append(vals, name)
		}
		sort.Strings(vals)
		return carapace.ActionValues(vals...).Tag("aliases")
	})
}

func addAliasNames(path string, names map[string]struct{}) {
	for _, name := range aliasNames(path) {
		names[name] = struct{}{}
	}
}

func systemConfigPaths(c carapace.Context) []string {
	if path, ok := c.LookupEnv("WORKTRUNK_SYSTEM_CONFIG_PATH"); ok && path != "" {
		return []string{path}
	}

	var dirs []string
	if dirsEnv := c.Getenv("XDG_CONFIG_DIRS"); dirsEnv != "" {
		for _, dir := range strings.Split(dirsEnv, ":") {
			if dir != "" {
				dirs = append(dirs, dir)
			}
		}
	}
	if len(dirs) == 0 {
		switch runtime.GOOS {
		case "darwin":
			dirs = []string{"/Library/Application Support", "/etc/xdg"}
		case "windows":
			if programData := c.Getenv("PROGRAMDATA"); programData != "" {
				dirs = []string{programData}
			}
		default:
			dirs = []string{"/etc/xdg"}
		}
	}

	paths := make([]string, 0, len(dirs))
	for _, dir := range dirs {
		paths = append(paths, filepath.Join(dir, "worktrunk", "config.toml"))
	}
	return paths
}

func projectConfigPath(root string, c carapace.Context) string {
	if path, ok := c.LookupEnv("WORKTRUNK_PROJECT_CONFIG_PATH"); ok && path != "" {
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
