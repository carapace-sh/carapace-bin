package cmd

import (
	"os"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/xdg"
	"github.com/spf13/cobra"
)

var clearCacheCmd = &cobra.Command{
	Use:   "--clear-cache",
	Short: "clear caches",
	RunE: func(cmd *cobra.Command, args []string) error {
		cacheDir, err := xdg.UserCacheDir()
		if err != nil {
			return err
		}
		return os.RemoveAll(cacheDir + "/carapace")
	},
}

func init() {
	carapace.Gen(listCmd).Standalone()
}
