package completers

import (
	"maps"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/carapace-sh/carapace-bridge/pkg/bridges"
	bridge_env "github.com/carapace-sh/carapace-bridge/pkg/env"
	"github.com/carapace-sh/carapace/pkg/xdg"
)

func Names() ([]string, error) {
	configDir, err := xdg.UserConfigDir()
	if err != nil {
		return nil, err
	}

	// internal completers
	names := slices.Collect(maps.Keys(completers))

	// specs
	entries, err := os.ReadDir(filepath.Join(configDir, "carapace", "specs"))
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".yaml") {
			names = append(names, strings.TrimSuffix(entry.Name(), ".yaml"))
		}
	}

	// bridges
	for _, b := range bridge_env.Bridges() {
		switch b {
		case "bash":
			names = append(names, bridges.Bash()...)
		case "fish":
			names = append(names, bridges.Fish()...)
		case "inshellisense":
			names = append(names, bridges.Inshellisense()...)
		case "zsh":
			names = append(names, bridges.Zsh()...)
		}
	}

	// choices
	entries, err = os.ReadDir(filepath.Join(configDir, "carapace", "choices"))
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			names = append(names, entry.Name())
		}
	}

	sort.Strings(names)
	return slices.Compact(names), nil
}
