package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var cache_configureCmd = &cobra.Command{
	Use:    "configure",
	Short:  "Configure Nix to use the Devbox cache as a substituter",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_configureCmd).Standalone()

	cache_configureCmd.Flags().String("user", "", "")
	cacheCmd.AddCommand(cache_configureCmd)

	carapace.Gen(cache_configureCmd).FlagCompletion(carapace.ActionMap{
		"user": os.ActionUsers(),
	})
}
