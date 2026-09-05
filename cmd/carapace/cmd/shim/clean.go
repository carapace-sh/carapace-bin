package shim

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/xdg"
)

func removeObsolete() error {
	configDir, err := xdg.UserConfigDir()
	if err != nil {
		return err
	}
	return removeObsoleteIn(configDir)
}

func removeObsoleteIn(configDir string) error {
	entries, err := os.ReadDir(filepath.Join(configDir, "carapace", "bin"))
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		return nil
	}

	for _, entry := range entries {
		specPath := filepath.Join(configDir, "carapace", "specs", strings.TrimSuffix(entry.Name(), ".exe")+".yaml")
		if _, err := os.Stat(specPath); os.IsNotExist(err) {
			shimPath := filepath.Join(configDir, "carapace", "bin", entry.Name())

			carapace.LOG.Printf("removing shim %#v", shimPath)
			if err := os.Remove(shimPath); err != nil {
				return err
			}
		}
	}
	return nil
}
