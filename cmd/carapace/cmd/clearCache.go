package cmd

import (
	"os"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/xdg"
	"github.com/spf13/cobra"
)

var clearCacheCmd = &cobra.Command{
	Use:   "--clear-cache",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		cacheDir, err := xdg.UserCacheDir()
		if err != nil {
			return err
		}
		path := cacheDir + "/carapace"
		return os.RemoveAll(path)
	},
}

func init() {
	carapace.Gen(listCmd).Standalone()
}
