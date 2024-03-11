package shim

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/xdg"
)

func removeObsolete() error {
	configDir, err := xdg.UserConfigDir()
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(configDir + "/carapace/bin")
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		return nil
	}

	for _, entry := range entries {
		specPath := fmt.Sprintf("%v/carapace/specs/%v.yaml", configDir, strings.TrimSuffix(entry.Name(), ".exe"))
		if _, err := os.Stat(specPath); os.IsNotExist(err) {
			shimPath := fmt.Sprintf("%v/carapace/bin/%v", configDir, entry.Name())
			if runtime.GOOS == "windows" {
				shimPath += ".exe"
			}

			carapace.LOG.Printf("removing shim %#v", shimPath)
			if err := os.Remove(shimPath); err != nil {
				return err
			}
		}
	}
	return nil
}
