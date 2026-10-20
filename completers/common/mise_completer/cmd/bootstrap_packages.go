package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_packagesCmd = &cobra.Command{
	Use:   "packages",
	Short: "Manage packages from `[bootstrap.packages]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_packagesCmd).Standalone()

	bootstrap_packagesCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrapCmd.AddCommand(bootstrap_packagesCmd)
}
